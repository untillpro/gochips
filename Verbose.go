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

All functions use Output to actually output the final string

*/

package gochips

import (
	"fmt"
)

// Doing printlns "$obj..." to stdout
var Doing func(arg interface{})

// Info printlns to stdout
var Info func(args ...interface{})

// Verbose If IsVerbose is true  printlns VerbosePrefix + subj + ": " + ...args
// args are printed in %#v format (json-like) if jsonLike is true, %v othervide
var Verbose func(subj string, args ...interface{})

// IsVerbose enables Verbose
var IsVerbose = false

// Output is used by all functions
var Output func(funcName, s string)

// VerbosePrefix prefixes Verbose output
var VerbosePrefix = "--- "

func init() {
	Doing = implDoing
	Info = implInfo
	Verbose = implVerbose
	Output = implOutput
}

func implDoing(arg interface{}) {
	Output("Doing", fmt.Sprintln(fmt.Sprintf("%v...", arg)))
}

func implInfo(args ...interface{}) {
	Output("Info", fmt.Sprintln(args...))
}

func implVerbose(subj string, args ...interface{}) {

	if IsVerbose {
		res := VerbosePrefix + subj + ": "
		for idx, arg := range args {
			if idx > 0 {
				res += ", "
			}
			res += fmt.Sprintf("%+v", arg)
		}
		Output("Verbose", fmt.Sprintln(res))
	}
}

func implOutput(funcName, s string) {
	fmt.Print(s)
}
