/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

Piped execution a-la shell

	// echo README.md | grep README.md | sed s/READ/forgive/

	err := new(PipedExec).
		Command("echo", "README.md").WorkingDir("/").
		Command("grep", "README.md").
		Command("echo", "good").
		Run(os.Stdout, os.Stdout)
	assert.Nil(t, err)

Ref. also PipedExec_test.go


Links

- https://stackoverflow.com/questions/25190971/golang-copy-exec-output-to-log

*/

package gochips

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"sync"
)

// https://github.com/b4b4r07/go-pipe/blob/master/README.md

// PipedExec allows to execute commands in pipe
type PipedExec struct {
	cmds []*pipedCmd
}

// Stderr redirection
const (
	StderrRedirectNone = iota
	StderrRedirectStdout
	StderrRedirectNull
)

type pipedCmd struct {
	stderrRedirection int
	cmd               *exec.Cmd
}

func (Self *PipedExec) command(name string, stderrRedirection int, args ...string) *PipedExec {
	cmd := exec.Command(name, args...)
	lastIdx := len(Self.cmds) - 1
	if lastIdx > -1 {
		var err error
		cmd.Stdin, err = Self.cmds[lastIdx].cmd.StdoutPipe()
		PanicIfError(err)
	} else {
		cmd.Stdin = os.Stdin
	}
	Self.cmds = append(Self.cmds, &pipedCmd{stderrRedirection, cmd})
	return Self
}

// GetCmd returns cmd with given index
func (Self *PipedExec) GetCmd(idx int) *exec.Cmd {
	return Self.cmds[idx].cmd
}

// Command adds a command to a pipe
func (Self *PipedExec) Command(name string, args ...string) *PipedExec {
	return Self.command(name, StderrRedirectNone, args...)
}

// WorkingDir sets working directory for the last command
func (Self *PipedExec) WorkingDir(wd string) *PipedExec {
	pipedCmd := Self.cmds[len(Self.cmds)-1]
	pipedCmd.cmd.Dir = wd
	return Self
}

// Wait until all cmds finish
func (Self *PipedExec) Wait() error {
	for _, cmd := range Self.cmds {
		err := cmd.cmd.Wait()
		if nil != err {
			return err
		}
	}
	return nil
}

// Start all cmds
func (Self *PipedExec) Start(out io.Writer, err io.Writer) error {
	for _, cmd := range Self.cmds {
		if cmd.stderrRedirection == StderrRedirectNone && nil != err {
			cmd.cmd.Stderr = err
		}
	}
	lastIdx := len(Self.cmds) - 1
	if lastIdx < 0 {
		return errors.New("Empty command list")
	}
	if nil != out {
		Self.cmds[lastIdx].cmd.Stdout = out
	}

	for _, cmd := range Self.cmds {
		Verbose("PipedExec", cmd.cmd.Path, cmd.cmd.Args)
		err := cmd.cmd.Start()
		if nil != err {
			return err
		}
	}
	return nil
}

// Run starts the pipe
func (Self *PipedExec) Run(out io.Writer, err io.Writer) error {
	e := Self.Start(out, err)
	if nil != e {
		return e
	}
	return Self.Wait()
}

// RunToStrings runs the pipe and saves outputs to strings
func (Self *PipedExec) RunToStrings() (stdout string, stderr string, err error) {
	// _, stdoutw := io.Pipe()
	// _, stderrw := io.Pipe()

	var wg sync.WaitGroup

	wg.Add(2)

	lastCmd := Self.cmds[len(Self.cmds)-1]
	stdoutPipe, err := lastCmd.cmd.StdoutPipe()
	if nil != err {
		return "", "", err
	}
	stderrPipe, err := lastCmd.cmd.StderrPipe()
	if nil != err {
		return "", "", err
	}

	err = Self.Start(nil, nil)
	if nil != err {
		return "", "", err
	}

	go func() {
		defer wg.Done()
		buf := new(bytes.Buffer)
		buf.ReadFrom(stdoutPipe)
		stdout = buf.String()
	}()

	go func() {
		defer wg.Done()
		buf := new(bytes.Buffer)
		buf.ReadFrom(stderrPipe)
		stderr = buf.String()
	}()

	wg.Wait()

	return stdout, stderr, nil

}
