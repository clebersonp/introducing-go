package main

import "fmt"

const (
	pi float64 = 3.14159
	e  float64 = 2.71828
)

// use the keyword var (or const) followed by parentheses with each variable on its own line.
var (
	a = 5
	b = 10
	c = 15
)

func main() {
	// const
	fmt.Println(pi, e)

	//  var
	fmt.Println(a, b, c)
}
