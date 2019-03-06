/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

Simple logging for CLI utilities

*/

package gochips

import (
	"fmt"
	"log"
)

// Doing prints "$obj..." to stdout
var Doing func(obj interface{})

// Infonl printlns to stdout
var Infonl func(obj ...interface{})

// Info prints obj to stdout
var Info func(obj ...interface{})

// Debug  log.Println()
// A newline is appended if the last character of s is not already a newline https://golang.org/pkg/log/#Logger.Output
var Debug func(obj ...interface{})

func init() {
	Doing = impl_Doing
	Infonl = impl_Infonl
	Info = impl_Info
	Debug = impl_Debug
}

func impl_Doing(obj interface{}) {
	fmt.Println(obj, "...")
}

func impl_Infonl(obj ...interface{}) {
	fmt.Println(obj...)
}

func impl_Info(obj ...interface{}) {
	fmt.Print(obj...)
}

func impl_Debug(obj ...interface{}) {
	log.Println(obj...)
}
