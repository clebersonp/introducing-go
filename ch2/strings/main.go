package main

import "fmt"

func main() {
	// double cotes accept escape character
	fmt.Println("Hello,\nworld") // \n is a new line
	fmt.Println(len("Hello, world"))
	// The operation will return an uint8 (byte)
	// will print 101 because the character is represented by a byte (byte is an integer)
	// indexes in Go starts at 0
	fmt.Println("Hello, world"[1])
	fmt.Println("Hello, " + "World") // concatenation

	// ` backticks are used to print multiline strings
	// ` backticks will ignore scape character
	fmt.Println(`Hello,
\nWorld`)
}
