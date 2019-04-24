/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

Simple logging for CLI utilities using functions:

- Error
- Doing
- Info
- Verbose

All functions use Output to actually output the final string

- VerboseWriters() are intended to be used with PipedExec.Run()

*/

package gochips

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Doing printlns "$obj..." to stdout
var Doing func(arg interface{})

// Info printlns to stdout
var Info func(args ...interface{})

// Error printlns to stdout
var Error func(args ...interface{})

// Verbose If IsVerbose is true  printlns VerbosePrefix + subj + ": " + ...args
// args are printed in %#+v format (json-like)
var Verbose func(subj string, args ...interface{})

// IsVerbose enables Verbose
var IsVerbose = false

// Output is used by all functions
var Output func(funcName, s string)

// VerbosePrefix prefixes Verbose output
var VerbosePrefix = "--- "

// ErrorPrefix prefixes Verbose output
var ErrorPrefix = "*** "

// VerboseWriters returns (os.Stdout, os.Stderr) if IsVerbose, (ioutil.Discard, os.Stderr) otherwise
var VerboseWriters func() (out io.Writer, err io.Writer)

func init() {
	Doing = implDoing
	Info = implInfo
	Verbose = implVerbose
	Error = implError
	Output = implOutput
	VerboseWriters = implVerboseWriters
}

func implDoing(arg interface{}) {
	Output("Doing", fmt.Sprintln(fmt.Sprintf("%v...", arg)))
}

func implInfo(args ...interface{}) {
	Output("Info", fmt.Sprintln(args...))
}

func implError(args ...interface{}) {
	Output("Error", ErrorPrefix+fmt.Sprintln(args...))
}

func implVerbose(subj string, args ...interface{}) {

	if IsVerbose {
		res := VerbosePrefix + subj + ": "
		for idx, arg := range args {
			if idx > 0 {
				res += ", "
			}
			res += fmt.Sprintf("%#+v", arg)
		}
		Output("Verbose", fmt.Sprintln(res))
	}
}

func implOutput(funcName, s string) {
	fmt.Print(s)
}

func implVerboseWriters() (out io.Writer, err io.Writer) {
	if IsVerbose {
		return os.Stdout, os.Stderr
	}
	return ioutil.Discard, os.Stderr
}
