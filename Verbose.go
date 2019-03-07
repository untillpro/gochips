/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

Simple logging for CLI utilities using functions:

- Doing
- Info
- Verbose

*/

package gochips

import (
	"fmt"
	"log"
)

// Doing printlns "$obj..." to stdout
var Doing func(arg interface{})

// Info printlns to stdout
var Info func(args ...interface{})

// Verbose If IsVerbose is true  prints VerbosePrefix + subj + ": " + ...args
// args are printed in %#v format (json-like)
// A newline is always appended if the last character is not already a newline https://golang.org/pkg/log/#Logger.Output
var Verbose func(subj string, args ...interface{})

// IsVerbose enables Verbose
var IsVerbose = false

// VerbosePrefix prefixes Verbose output
var VerbosePrefix = "--- "

func init() {
	Doing = implDoing
	Info = implInfo
	Verbose = implVerbose
}

func implDoing(arg interface{}) {
	fmt.Println(fmt.Sprintf("%v...", arg))
}

func implInfo(args ...interface{}) {
	fmt.Println(args...)
}

func implVerbose(subj string, args ...interface{}) {
	if IsVerbose {
		res := VerbosePrefix + subj + ": "
		for _, arg := range args {
			res += fmt.Sprintf("%#v, ", arg)
		}
		log.Print(res)
	}
}
