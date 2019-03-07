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
)

// Doing printlns "$obj..." to stdout
var Doing func(arg interface{})

// Info printlns to stdout
var Info func(args ...interface{})

// Verbose If IsVerbose is true  printlns VerbosePrefix + subj + ": " + ...args
// args are printed in %#v format (json-like) if jsonLike is true, %v othervide
var Verbose func(subj string, jsonLike bool, args ...interface{})

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

func implVerbose(subj string, jsonLike bool, args ...interface{}) {
	var fmtStr string
	if jsonLike {
		fmtStr = "%#v, "
	} else {
		fmtStr = "%v, "
	}

	if IsVerbose {
		res := VerbosePrefix + subj + ": "
		for _, arg := range args {
			res += fmt.Sprintf("%#v, ", arg)
		}
		fmt.Println(res)
	}
}
