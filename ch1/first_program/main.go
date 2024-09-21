// Package main is the entry point for the program.
// Go program must start with the package name.
// Packages are Go's way of organizing and reusing code.
//
//	There are two types of Go programs:
//		- Executables: Package name must be main and must have function main().
//		- Libraries: Can be any valid name. For convention, we use the same name of directory where is the code.
package main

// import is used to include other packages into the current code
import "fmt"

// Run from root directory
// go run ./ch1/first_program/main.go

// main function is special because it's a function that gets called when you execute the program.
// All functions start with the keyword func followed by the name of the function, a list of zero or more parameters
// surrounded by parentheses, an optional return type, and a body which is surrounded by curly braces and zero or more
// steps called statements between curly braces.
func main() {
	// this is a comment
	// To see documentation in the terminal:
	// go doc fmt Println
	fmt.Println("Hello, World")
	fmt.Println("Hello, my name is Golang")
}
