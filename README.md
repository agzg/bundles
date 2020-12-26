# bundles
A bundle is, as defined by the Oxford Languages Dictionary, "a collection of things or quantity of material tied or wrapped up together." And as you will see, reading further, the program that resides in this repository is quite the same.

Bundles launches multiple apps, URLs, files and shell commands (specified by you) in groups/clusters called, well, bundles. This program is aimed at making the process of launching the same common applications that you use everyday painless by allowing you to launch all of them with a **single command**, instead of opening each individually.

## Details
Bundles is, in some sense, a pseudo-compiler. It has a very simple language, with a [simple pre-defined syntax](#syntax), that can be used to create executables that launch bundles of apps, files, etc. (called "actions" or "items"). The name of the bundle (called a "label") should be passed as an argument to the executable.

The source code from the provided file are broken down into syntactic tokens (labels and actions). A pair of a label and action/item slice makes a "bundle". This logic is transpiled into a temporary .go file, fittingly `./temp/temp.go`, that is then compiled into an executable.

## Getting Started
### Requirements
Bundles requires the following:
* An installation of [GoLang](https://golang.org/dl/) >= 1.15
* `"github.com/skratchdot/open-golang/open"`

### Usage
Download the source code, either directly or using `git clone https://github.com/aliazam/bundles.git` after ensuring you have an installation of GoLang and have installed `"github.com/skratchdot/open-golang/open"` using the `go get` command.

Read [the syntax section](#syntax) to learn how to write a Bundles script, and [the execution section](#execution) to learn how to execute said script.

## Syntax
This section covers the syntax that you must follow in your Bundles scripts.

The "label" is the name of the bundle, and is passed as an argument to the executable created to launch the "actions" within. It can only contain alphanumeric characters and underscores, ending with a single ":" (colon) sign, without the quotes.

Examples: `Writing:` `Just_An_Awesome_Label_123:` `Books:`

An "action" is a single instance of an application within a bundle. It encompasses what apps, files, commands and websites are launched, executed or opened. They follow a general syntax: name of action (either "app", "url", "shell", or "file" without the quotes), a "," (comma) sign, and a "value". As of `v1.018e9724` there are **four** possible actions:
* app, launches an executable application; syntax: `app, path/to/app.exe`
* file, opens a file using the default application; syntax: `file, path/to/file`
* url, opens a website using the default browser; syntax: `url, https://url-of-website.com`
* shell, executes a command using the system terminal; syntax: `shell, command`

Examples:
`app, "C:\Program Files\WhatsApp\WhatsApp.exe"` `file, D:\Bundles\main.go` `url, https://www.github.com` `shell, ls`

Comments, in the Bundles syntax, can **not** be inline. The syntax for comments is the same as that of the Python programming language, a "#" symbol, without the quotes, followed by whatever text you want to write.

Finally, it is recommended that you indent actions within a bundle, leaving an empty line between each declared bundle (yes, there can be multiple within a file). Below is an example of a bundles files ready to be compiled:

```
programming:
  # the things I need to get started to program!
  app, "C:\Program Files\VSC\Code.exe"
  file, D:\Recent\Arctic\main.c
  shell, go run D:/Recent/Dashboard/dashboard.go
  url, www.github.com
  
Books:
  # the books that I am reading right now
  file, "D:\Books\The Art of Assembly Language.pdf"
  file, "D:\Books\Black Beauty.pdf"
  url, https://www.goodreads.com
```

## Execution
🔪🩸💀

No, not that kind of execution! This section covers how you go about compiling and running a Bundles script.

Go to the root directory and run, `go run main.go filename.txt` replacing "filename.txt" with the absolute or relative path to the file where you have written your Bundles script. You can even compile `./main.go` and use the executable as a sort of a compiler for all of your Bundles scripts!

Assuming you followed [the syntax section](#syntax) properly, you should see an executable file in the current directory. It's name will be the same as that of the script file supplied, in my case "filename.exe".

Go ahead and run, `./filename example_bundle` and voilà, all the specified actions under the `example_bundle` bundle execute!

### Errors
**Upon compilation:**
SyntaxErrors, ValueErrors, and IllegalChars will tell you if your source code does not follow [the specified syntax](#syntax), has invalid values (paths and/or URLs), or contains a  "\`" (backtick) characters, respectively.

**Upon execution:**
If you see a message like, "Bundle does not exist", it means that either there is no bundle labelled with the argument you just supplied, or you did not supply any arguments at all! As mentioned, the syntax for executing a bundle (.exe) is, `./name-of-executable-file name-of-valid-label`

Refer to the [syntax](#syntax) section for more.

## License
Use of this source code is governed under the MIT license that can be found in [the LICENSE file](https://github.com/aliazam/bundles/blob/main/LICENSE) and below:
```MIT License

Copyright (c) 2020 Ali Azam

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
## Addendum
### Upcoming Features
* A special `init` label, that executes bundles on start-up or execution.
* Ability to launch apps without needing paths to their executables.
* Ability to execute commands in separate terminals, if specified.
* Ability to open files and websites in applications other than their defaults.
* Freeze code. Some features (especially on MacOS and Linux) haven't been tested/implemented yet.
### Credits
An especial thanks to [@Abdul-Muiz-Iqbal](https://github.com/Abdul-Muiz-Iqbal) who asked for ideas for making small funky projects, such as this one.

Feel free to open up issues for suggestions, improvements, or bug-smashing!

**written by [@aliazam (Alee)](https://www.github.com/aliazam)**
