// Copyright 2020 @aliazam (Alee). All Rights Reserved.
// Use of this source code is governed under the MIT license
// that can be found in the LICENSE file.

// parser: compiles Bundles from source files to .exe-s.

package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Parse Read-s, Tokenize-s, and Compile-s the source code in
// the file at the given path.
//
// The executable resides in the current directory, with it's
// name the same as that of the file at the provided path*
//
// * the first word of the file's name
func Parse(path string) {
	name := strings.ReplaceAll(filepath.Base(path),
		     filepath.Ext(path), "")
	name = strings.Fields(name)[0]
	Compile(Tokenize(Read(path)), name)
}

// Read reads the contents of the file at the given path and
// formats it to be further used for tokenization.
//
// Errors out if an invalid path, or insufficient permissions
// are provided.
func Read(name string) []string {
	f, err := os.Open(name)
	report(err, "Unable to open " + name)

	b, err := ioutil.ReadAll(f)
	report(err, "Unable to read " + name)

	ss := strings.Split(string(b), "\n")
	return ss
}

// An Item is an Action + Value pair.
// The prescribed action is executed (e.g. app is launched)
// using the given value.
type Item struct {
	// Can either be "app", "url", "file" or "shell".
	Action	string
	// Contains the value for above action.
	Value	string
}

// A Bundle is a single launch-able/executable unit of code
// in Bundles.
// It consists of a slice of Item, which are the actions
// launched when the Bundle, using it's Label, is executed.
type Bundle struct {
	// Name of the bundle.
	Label	string
	// A slice of Item-s.
	Items	[]Item
}

// Regex patterns of the syntax, used during tokenization.
const (
	labelPattern string = "^[a-zA-Z0-9_]+:$"
	itemsPattern string = "^((app)|(file)|(shell)|(url)) ?, ?[^\n`]+$"
)

// Tokenize breaks down the source code into units of lexically
// analysed Bundles.
//
// It is a pseudo-lexer, if you will, that performs the functions
// of a lexer, breaking down the source code into Bundle-s and
// Label-s to be compiled.
//
// Errors out in case of syntactic errors, and if invalid paths
// (according to context and system) are provided.
func Tokenize(source []string) []Bundle {
	var bundles []Bundle
	var bundle  Bundle

	for i, s := range source {
		s = strings.TrimSpace(s)

		switch {
		case matches(labelPattern, s):
		if bundle.Label != "" {
			if len(bundle.Items) == 0 {
				SyntaxError("Empty bundle.", bundle, i)
			}
			bundles = append(bundles, bundle)
			bundle  = Bundle{}
		}
		label := strings.TrimSpace(s)
		label = strings.Trim(s, ":")

		bundle.Label = label

		case matches(itemsPattern, s):
		if bundle.Label == "" {
			SyntaxError("Unlabelled bundle.", bundle, i)
		}
		item   := Item{}
		values := strings.Split(s, ",")
		for i, value := range values {
			values[i] = strings.TrimSpace(value)
		}

		action := strings.ToLower(values[0])
		value  := strip(strings.Join(values[1:], ","))

		if action == "app" {
			if !isExec(value) {
				ValueError("No .exe at path.", bundle, i)
			}
		} else if action == "file" {
			if !isFile(value) {
				ValueError("No file at path.", bundle, i)
			}
		} else if action == "url" {
			if !isURL(value) {
				ValueError("Invalid URL.", bundle, i)
			}
		} else if action != "shell" {
			SyntaxError("Unrecognized action.", bundle, i)
		}

		item.Action = action
		item.Value  = value

		bundle.Items = append(bundle.Items, item)

		case s == "" || strings.HasPrefix(s, "#"):
		continue

		case strings.Contains(s, "`"):
		IllegalChar("` is an illegal character.", bundle, i)

		default:
		SyntaxError("Invalid token/s.", bundle, i)
		}
		if len(source) - 1 == i {
			if bundle.Label != "" && len(bundle.Items) == 0 {
				SyntaxError("Empty bundle.", bundle, i)
			}
			bundles = append(bundles, bundle)
		}
	}
	return bundles
}

// Pre-written code for transpilation.
const (
	transCode string = `
// Pre-written transpilation code to GoLang.
// Written by Bundles (which in turn was written by Alee).

package main

import (
	"fmt"
	"os"
	"../launcher"
)

func main() {
	os.Args = append(os.Args, "")
	switch os.Args[1] {%s
	default:
	fmt.Println("Bundle does not exist.")
	}
}
`
	caseCode  string = "\n\tcase \"%s\":\n"
	itemCode  string = "\tlauncher.Launch(\"%s\", `%s`)"
)

// Compile turns a slice of Bundle-s into an executable file in
// the current directory.
//
// It does so by transpiling the Bundle-s into idiomatic Go code
// that is then built, "go build ...", in the current directory.
// The new executable takes the name of the file provided at the
// start (final).
//
// Errors out if it is unable to create, build or access
// the ./temp/temp.go file.
func Compile(bundles []Bundle, exe string) {
	if len(bundles) == 0 { return }

	var addition string
	for _, bundle := range bundles {
		addition += fmt.Sprintf(caseCode, bundle.Label)

		for _, item := range bundle.Items {
			addition += fmt.Sprintf(itemCode, item.Action, item.Value)
			addition += "\n"
		}
	}
	code  := strings.TrimSpace(transCode)
	bcode := []byte(fmt.Sprintf(code, addition))

	if _, err := os.Stat("./temp"); os.IsNotExist(err) {
		err = os.Mkdir("./temp", 0777)
		report(err, "Unable to create temporary directory; DIY.")
	}

	_, err := os.Create("./temp/temp.go")
	report(err, "Unable to create ./temp/temp.go.")

	err = ioutil.WriteFile("./temp/temp.go", bcode, 0777)
	report(err, "Unable to write to ./temp/temp.go.")

	err = exec.Command("go", "build", "-o", exe, "./temp/temp.go").Run()
	report(err, "Unable to compile ./temp/temp.go.")
}

// report gracefully handles and describes errors.
func report(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

// matches reports whether a string matches the regular expression
// pattern provided.
func matches(pattern, s string) bool {
	matched, _ := regexp.Match(pattern, []byte(s))
	return matched
}

// strip removes " and 's from the start and end of the provided
// string.
func strip(s string) string {
	if s[0] == '\'' && s[len(s) - 1] == '\'' {
		s = strings.Trim(s, "'")
	} else if s[0] == '"' && s[len(s) - 1] == '"' {
		s = strings.Trim(s, "\"")
	}
	return s
}

// isFile reports whether the provided path with a valid and
// accessible file.
func isFile(name string) bool {
	info, err := os.Stat(name)
	if err != nil { return false }

	if info.Mode().IsRegular() {
		return true
	}
	return false
}

// isExec reports whether the provided path ends with a valid
// executable file (application).
//
// Note: the author is unsure whether this affects apps on
// MacOS, as he does not own a Mac!
func isExec(name string) bool {
	return filepath.Ext(name) == ".exe"
}

// isURL uses regular expression (it's long and ugly!) to report
// whether the provided URL is valid or not.
func isURL(name string) bool {
	urlPattern := "[(http(s)?):\\/\\/(www\\.)?a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)"
	return matches(urlPattern, name)
}

// SyntaxErrors gives the location and Bundle, most of
// the time, of syntactic errors in the provided source code.
func SyntaxError(msg string, bundle Bundle, i int) {
	if bundle.Label != "" {
		format := "SyntaxError: in %s, on line %d\n\t%s\n"
		fmt.Printf(format, bundle.Label, i + 1, msg)
	} else {
		format := "SyntaxError: on line %d\n\t%s\n"
		fmt.Printf(format, i + 1, msg)
	}
	os.Exit(1)
}

// ValueErrors gives the location and Bundle of invalid values 
// in the provided source code.
func ValueError(msg string, bundle Bundle, i int) {
	if bundle.Label != "" {
		format := "ValueError: in %s, on line %d\n\t%s\n"
		fmt.Printf(format, bundle.Label, i + 1, msg)
	} else {
		format := "ValueError: on line %d\n\t%s\n"
		fmt.Printf(format, i + 1, msg)
	}
	os.Exit(1)
}

// IllegalChar gives the location and Bundle of any invalid 
// characters in the provided source code.
func IllegalChar(msg string, bundle Bundle, i int) {
	if bundle.Label != "" {
		format := "IllegalChar: in %s, on line %d\n\t%s\n"
		fmt.Printf(format, bundle.Label, i + 1, msg)
	} else {
		format := "IllegalChar: on line %d\n\t%s\n"
		fmt.Printf(format, i + 1, msg)
	}
	os.Exit(1)
}
