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

import "testing"

func Test_implDoing(t *testing.T) {
	Doing("Doing")
}

func Test_implInfo(t *testing.T) {
	Info("Info1")
	Info("Info1.1", "Info1.2")
}

type myPoint struct {
	x int
	y int
}

func Test_implVerbose(t *testing.T) {
	Verbose("Should not see it")
	IsVerbose = true
	Verbose("mysubj", "Verbose 1")

	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	Verbose("map", m)
	Verbose("myPoint", myPoint{1, 2})
	Verbose("&myPoint", &myPoint{1, 2})
}
