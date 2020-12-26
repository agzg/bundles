// Copyright 2020 Ali Azam (Alee). All Rights Reserved.
// Use of this source code is governed under the MIT license
// that can be found in the LICENSE file.

// launcher: launches applications, files, URLs, and commands.

package launcher

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"os/exec"
	"path/filepath"
	"strings"
)

// Launch executes applications, files, opens URLs and run shell
// commands as subsidiary for the parser package.
//
// It heavily relies on the "github.com/skratchdot/open-golang/open"
// package (thank you @skratchdot for this package!), so ensure it's
// installed on the $GOPATH.
func Launch(action, value string) {
	var err error

	switch action {
	case "app":
	path, _ := filepath.Abs(value)
	err = open.Start(path)

	case "url":
	err = open.Start(value)

	case "file":
	path, _ := filepath.Abs(value)
	err = open.Start(path)

	case "shell":
	fmt.Println("$", value)

	args := strings.Fields(value)
	if len(args) > 1 {
		b, _ := exec.Command(args[0], args[1:]...).Output()
		fmt.Println(string(b))
	} else {
		b, _ := exec.Command(args[0]).Output()
		fmt.Println(string(b))
	}
	}

	if err != nil {
		fmt.Println("Unable to launch", value)
	}
}
