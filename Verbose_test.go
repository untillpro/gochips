/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

Simple logging for CLI utilities
- Doing
- Infonl
- Info
- Verbose

*/

package gochips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VerboseWriters(t *testing.T) {

	IsVerbose = false

	err := new(PipedExec).
		Command("echo", "!!! Should not see it").
		Run(VerboseWriters())
	assert.Nil(t, err)

	err = new(PipedExec).
		Command("sh", "-c", "echo '***' Redirect, should see 1>&2").
		Run(VerboseWriters())
	assert.Nil(t, err)

	IsVerbose = true
	err = new(PipedExec).
		Command("echo", "*** Verbose Should see it").
		Run(VerboseWriters())
	assert.Nil(t, err)
}

func Test_Error(t *testing.T) {
	Error("This is an error", "err arg1", "err arg2")
}

func Test_Doing(t *testing.T) {
	Doing("Doing")
}

func Test_Info(t *testing.T) {
	Info("Info1")
	Info("Info1.1", "Info1.2")
}

type myPoint struct {
	x int
	y int
}

func Test_Verbose(t *testing.T) {

	IsVerbose = false

	Verbose("Should not see it")
	IsVerbose = true
	Verbose("mysubj", "Verbose 1")
	Verbose("two values", "Verbose 1.1", "Verbose 1.2")
	Verbose("slice", []string{"slice 1", "slice 2"})

	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	Verbose("map", m)
	Verbose("myPoint", myPoint{1, 2})
	Verbose("&myPoint", &myPoint{1, 2})
}
