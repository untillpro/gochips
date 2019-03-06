/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package gochips

import "fmt"

// Doing prints "$obj..." to stdout
var Doing func(obj interface{})

func doingImpl(obj interface{}) {
	fmt.Println(obj, "...")
}

func init() {
	Doing = doingImpl
}
