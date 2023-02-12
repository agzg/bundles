// Copyright 2021 Ali Azam. All Rights Reserved.
// Use of this source code is governed under the MIT license
// that can be found in the LICENSE file.

// main: the entry point of Bundles.

package main

import (
	"fmt"
	"os"

	"./parser"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("go run main.go <filename>")
		return
	}
	parser.Parse(os.Args[1])
}
