# bundles
Bundles allows you to gather the websites, applications, commands, and files that you frequently use, and launch them with a single command (using an executable).

### Details
Bundles is a transcompiler that transpiles (translates) source code (written in the [Bundles syntax](#syntax)) into idiomatic Go code and builds the file, thus creating an executable. A single executable can house multiple 'bundles' or clusters of the applications that you want to launch. 

A bundle consists of a "label", the name of the bundle used to refer to it, and "items", the specified applications that you want to launch or execute.

## Getting Started
### Requirements
Bundles requires the following:
* An installation of [Golang](https://golang.org/dl/) >= 1.15
* `go get github.com/skratchdot/open-golang/open`

### Usage
Download the "github.com/skratchdot/open-golang/open" package using `go get "github.com/skratchdot/open-golang/open"`, then download the source code (or clone the repository using `git clone https://github.com/agzg/bundles.git`).

Read [the syntax section](#syntax) to learn how to code in Bundles, and [the execution section](#execution) to learn how to compile the said script.

## Syntax
A "bundle", as mentioned before, consists of a "label" and "items". To create a bundle you must *declare* a "label" and ensure that you *add* at least one "item" to it.

The "label" is the name of the bundle, which is passed to the executable file created to launch the applications within the bundle. To declare it simply write the name you wish to give to it (alphanumeric characters and "\_" (underscores) only) followed by a ":" (colon).

Examples: `awesomeLabel123:` `Books:` `social_media:`

"Items" are the applications inside a bundle. To add items to a bundle, first declare a "label" (following the steps described above). Then (on the next line, after indenting) write either "app", "url", "file" or "shell" to launch an application, website, file, or (shell) command respectively; followed by a "," (comma). Then (right after the comma) add the value associated to the application you want to launch. For files and applications, it is their location; for websites, it is the URL of the website; for commands, it is the command itself.

**Note:** Websites are launched in the default browser.

Examples (without the optional indent): `app, "C:\Program Files\WhatsApp\WhatsApp.exe"` `file, "D:\Bundles\main.go"` `url, https://www.github.com` `shell, ls`

Finally, you can add comments to the code. They can **not** be inline, and follow the same syntax as that of the Python programming language (i.e. a "#" (hash) followed by any sort of text). Hope this one doesn't need any examples.

Below is an example of what a typical Bundles script should look like:

```
social_media:
  url, https://www.facebook.com
  url, https://www.instagram.com
  url, https://www.reddit.com
  
books:
  # Books I am reading right now.
  file, "D:\Books\The Art of Assembly Language.pdf"
  file, "D:\Books\Black Beauty.pdf"
  url, https://www.goodreads.com
  
code:
  app, "C:\Program Files\VSC\Code.exe"
  app, "C:\Program Files\GoLand\goland.exe"
  url, https://www.github.com
  shell, ls
```

## Execution
Go to the root directory and run, `go run main.go filename.txt` replacing "filename.txt" with the path to the file where you have written your Bundles script.

Assuming you followed [the syntax section](#syntax) properly, you should see an executable file in the current directory. It's name will be the same as that of the script file supplied, in my case "filename.exe".

Go ahead and run, `./filename example_bundle` and voilà, all the specified actions under the `example_bundle` bundle execute!

### Errors
SyntaxErrors, ValueErrors, and IllegalChars will tell you if your source code does not follow [the specified syntax](#syntax), has invalid values (paths and/or URLs), or contains a  \` (backtick character), respectively.

## License
[MIT](https://github.com/agzg/bundles/blob/main/LICENSE)

## Credits
An especial thanks to [mbyx](https://github.com/mbyx) whose request for funky project ideas led to this one.
