/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package gochips

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPipedExec_Basics(t *testing.T) {

	IsVerbose = true

	// echo
	{
		err := new(PipedExec).
			Command("echo", "hello").
			Run(os.Stdout, os.Stdout)
		assert.Nil(t, err)
	}

	// echo hello2 | grep hello2
	{
		err := new(PipedExec).
			Command("echo", "hello2").
			Command("grep", "hello2").
			Run(os.Stdout, os.Stdout)
		assert.Nil(t, err)
	}

	// echo hi | grep hello
	{
		err := new(PipedExec).
			Command("echo", "hi").
			Command("grep", "hello").
			Run(os.Stdout, os.Stdout)
		assert.NotNil(t, err)
	}

	// echo hi | grep hi | echo good
	{
		err := new(PipedExec).
			Command("echo", "hi").
			Command("grep", "hi").
			Command("echo", "good").
			Run(os.Stdout, os.Stdout)
		assert.Nil(t, err)
	}

	// ls at "/""
	{
		err := new(PipedExec).
			Command("ls").WorkingDir("/").
			Run(os.Stdout, os.Stdout)
		assert.Nil(t, err)
	}

}

// Working directory
func TestPipedExec_Wd(t *testing.T) {

	/* Create structure
	tmpDir
	  tmpDir1
	  	1.txt
	  tmpDir2
	  	2.txt
	*/

	tmpDir, err := ioutil.TempDir("", "Wd")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	tmpDir1, err := ioutil.TempDir(tmpDir, "Wd")
	assert.Nil(t, err)

	tmpDir2, err := ioutil.TempDir(tmpDir, "Wd")
	assert.Nil(t, err)

	ioutil.WriteFile(filepath.Join(tmpDir1, "1.txt"), []byte("11.txt"), 0644)
	ioutil.WriteFile(filepath.Join(tmpDir2, "2.txt"), []byte("21.txt"), 0644)

	// Run ls commands

	err = new(PipedExec).
		Command("ls", "1.txt").WorkingDir(tmpDir1).
		Run(os.Stdout, os.Stdout)
	assert.Nil(t, err)

	err = new(PipedExec).
		Command("ls", "2.txt").WorkingDir(tmpDir2).
		Run(os.Stdout, os.Stdout)
	assert.Nil(t, err)

	err = new(PipedExec).
		Command("ls", "1.txt").WorkingDir(tmpDir2).
		Run(os.Stdout, os.Stdout)
	assert.NotNil(t, err)

	err = new(PipedExec).
		Command("ls", "2.txt").WorkingDir(tmpDir1).
		Run(os.Stdout, os.Stdout)
	assert.NotNil(t, err)

}

func TestPipedExec_PipeFall(t *testing.T) {

	// echo hi | grep hi | echo good => OK
	{
		err := new(PipedExec).
			Command("echo", "hi").
			Command("grep", "hi").
			Command("echo", "good").
			Run(os.Stdout, os.Stdout)
		log.Println("***", err)
		assert.Nil(t, err)
	}

	// echo hi | grep hello | echo good => FAIL
	{
		err := new(PipedExec).
			Command("echo", "hi").
			Command("grep", "hello").
			Command("echo", "good").
			Run(os.Stdout, os.Stdout)
		log.Println("***", err)
		assert.NotNil(t, err)
	}
}

func TestPipedExec_WrongCommand(t *testing.T) {
	err := new(PipedExec).
		Command("qqqqqqjkljlj", "hello").
		Run(os.Stdout, os.Stdout)
	assert.NotNil(t, err)
	log.Println(err)
}

func TestPipedExec_EmptyCommandList(t *testing.T) {
	err := new(PipedExec).
		Run(os.Stdout, os.Stdout)
	assert.NotNil(t, err)
	log.Println(err)
}

func TestPipedExec_KillProcessUsingFirst(t *testing.T) {
	pe := new(PipedExec)
	pe.Command("sleep", "10")
	cmd := pe.GetCmd(0)

	go func() {
		defer fmt.Println("Bye")
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Killing process...")
			cmd.Process.Kill()
		}
	}()

	fmt.Println("Running...")
	err := pe.Run(os.Stdout, os.Stderr)
	fmt.Println("err=", err)
}
