# bundles
Bundles is a transcompiler, with [it's own syntax](#syntax), that allows you to launch clusters of routine applications. It transpiles the source file into Go code, creating a temporary file (fittingly "./temp/temp.go") that is then built using the native Go compiler (using the "build" command) thus creating an executable, the arguments to which launch the clusters of applications, called bundles.

### Wait, what?..
Bundles allows you to gather the websites, applications, commands, and files that you frequently use, and launch them with a single command (using an executable).

### Details
Bundles is a transcompiler that transpiles (translates) source code (written in the [Bundles syntax](#syntax)) into idiomatic Go code and builds the file, thus creating an executable. A single executable can (and should!) house multiple 'bundles' or clusters of the applications that you want to launch. 

A bundle consists of a "label", the name of the bundle used to refer to it, and "items", the specified applications that you want to launch or execute.

## Getting Started
### Requirements
Bundles requires the following:
* An installation of [GoLang](https://golang.org/dl/) >= 1.15
* `"github.com/skratchdot/open-golang/open"`

### Usage
Download the "github.com/skratchdot/open-golang/open" package using `go get "github.com/skratchdot/open-golang/open"`, then download the source code (or clone the repository using `git clone https://github.com/aliazam/bundles.git`). Ensure you also have GoLang installed on your system.

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
  # I am not very social!
  url, https://www.facebook.com
  url, https://www.instagram.com
  url, https://www.reddit.com
  
Books:
  # Books I am reading right now.
  file, "D:\Books\The Art of Assembly Language.pdf"
  file, "D:\Books\Black Beauty.pdf"
  url, https://www.goodreads.com
  
Programming:
  # Things I need to get started.
  app, "C:\Program Files\VSC\Code.exe"
  app, "C:\Program Files\GoLand\goland.exe"
  url, https://www.github.com
  shell, ls
```

## Execution
🔪🩸💀

No, not that kind of execution! This section covers how you go about compiling and running a Bundles script.

Go to the root directory and run, `go run main.go filename.txt` replacing "filename.txt" with the path to the file where you have written your Bundles script.

Assuming you followed [the syntax section](#syntax) properly, you should see an executable file in the current directory. It's name will be the same as that of the script file supplied, in my case "filename.exe".

Go ahead and run, `./filename example_bundle` and voilà, all the specified actions under the `example_bundle` bundle execute!

### Errors
**Upon compilation:**
SyntaxErrors, ValueErrors, and IllegalChars will tell you if your source code does not follow [the specified syntax](#syntax), has invalid values (paths and/or URLs), or contains a  "\`" (backtick) character, respectively.

**Upon execution:**
If you see a message like, "Bundle does not exist", it means that either there is no bundle labelled with the argument you just supplied, or you did not supply any arguments at all! As mentioned, the syntax for executing a bundle (.exe) is, `./name-of-executable-file name-of-declared-label`

## License
MIT

## Credits
An especial thanks to [@Abdul-Muiz-Iqbal](https://github.com/Abdul-Muiz-Iqbal) whose request for funky project ideas led to this one.

Feel free to open up issues for suggestions, improvements or bugs.

**written by [@aliazam (Alee)](https://www.github.com/aliazam)**
