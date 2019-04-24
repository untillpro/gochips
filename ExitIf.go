/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

Runtime assertions

*/

package gochips

import (
	"fmt"
	"os"
	"testing"
)

// Fatal exists with given parameters
func Fatal(args ...interface{}) {
	ExitIfFalse(false, args...)
}

// FatalIfError calls t.Fatal(args, err) if err is not nil
func FatalIfError(t *testing.T, err error, args ...interface{}) {
	if err == nil {
		return
	}
	args = append(args, err)
	t.Fatal(args...)
}

// ExitIfError s.e.
func ExitIfError(err error, args ...interface{}) {
	if nil != err {
		fmt.Fprintln(os.Stderr, args...)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// ExitIfFalse s.e.
func ExitIfFalse(cond bool, args ...interface{}) {
	if !cond {
		fmt.Fprintln(os.Stderr, args...)
		os.Exit(1)
	}
}

// PanicIfError panics if err is not null
func PanicIfError(err error) {
	if nil != err {
		panic(err)
	}
}
