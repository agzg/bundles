// Copyright 2020 @aliazam (Alee). All Rights Reserved.
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
		fmt.Println(`Usage:

$ go run main.go filename (optional: clean)
  filename 	the name of the file to be compiled.
  clean		(optional) deletes temporary files created during compilation.`)
		return
	}
	parser.Parse(os.Args[1])
}