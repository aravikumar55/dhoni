Lab: Introduction to Go
Take me to the lab!

Help for the VSCode editor.

Where was the Go programming language designed ?
VMware
Facebook
Google
Microsoft
ans: Reveal
Google

Go was created to support the following
The ease of programming of an interpreted, dynamically typed language.
For creating web pages easily.
Provide the efficiency and safety of a statically typed, compiled language.
Aimed to be modern, with support for networked and multicore computing.
ans: Reveal
A, C, D

Although Go is not dynamically typed or interpreted, the language was created with simplicity in mind to make it feel like using a language like Python which is these things.

It's not for creating web pages. That's HTML and JavaScript.

Compiled languages with static types provide safety - many bugs are caught by the compiler before you even get to run the program. Static typing provides efficiency because dynamic typing requires much more work behind the scenes (runs slower as a result).

Go has support built into the language itself for multicore computing such as "channels". You will learn about these in the Advanced Golang course. It also has great package support for networking - it's simple to construct API servers.

Go is a...
compiled language
interpreted language
ans: Reveal
compiled language

What is a package in go?
a file that ends with .mod extension.
a file that ends with .go extension
a collection of files that live in the same directory and have the same package statement at the beginning.
Set of core packages to enhance and extend the language.
ans: Reveal
C

In a Go program, where are packages declared ?
At the start of the program
After the import staetment
At the end of the program
Anywhere in the code.
ans: Reveal
At the start of the program

The first line of every .go file that is not a blank line or a comment must be a package statement to declare the package that the code belongs to.

Choose the correct format of importing a package in Go:
If we wanted to import the fmt package, which is the correct syntax?

import fmt
import (fmt)
import "fmt"
"import fmt"
ans: Reveal
import "fmt"

import is a keyword so must not be enclosed within quotes. The name of the package must be quoted.

What is the entry point in a Go program?
The function that's declared last
The main function
Function imported by fmt package
The function that's declared first
ans: Reveal
The main function

The go runtime looks for a function called main and calls it. That's how the program starts. This is taken from the venerable C language!

Which package consists of main() function ?
independent of package name
package main
we can create our own main function, no package needs to be imported
package greetings
ans: Reveal
package main

By convention, you create the main package first, and put the main() function in it:

package main

// Imports go here - if you need any

func main() {
    // your program starts here
}
Our program began with package main, what would the files in the fmt package begin with?
package os
package fmt
pacakge object
package main
ans: Reveal
package fmt

fmt is a package, which is meant to be imported by other programs. Here's a snippet of one of the files from this package:

package fmt

import (
    "internal/fmtsort"
    "io"
    "os"
    "reflect"
    "strconv"
    "sync"
    "unicode/utf8"
)

// Some lines removed for brevity

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
Choose the correct syntax to write a comment in Go
# this is a comment
// this is a comment
/*
this is
a multiline
comment
*/
"""
this is
a multiline
comment
"""
ans: Reveal
B, C

The other two commenting styles are those of Python.

Now let's run some go commands using the terminal. Open an new terminal and identify the version of go installed.
In the VSCode terminal window, run the following:

go version
Select the appropriate answer.

Which of the following commands is not valid?
If unsure, run go help and check out the commands and their uses.

go version
go run
go compile
go generate
ans: Reveal
go compile

This is not valid, because the command to compile without running is go build.

A simple go program has been created for you at the location /root/code/simple-project. However, there is an error in the code.
Identify and fix the problem.

In the explorer pane, click on simple-project to reveal main.go
Click on main.go to load it into the editor.
Review the answer to question 6 to know the issue, then edit the code so it is correct.

ans: Reveal
Packages imported with the import statement must have the package name in double-quotes.

Fix the import statement thus:

import "fmt"
