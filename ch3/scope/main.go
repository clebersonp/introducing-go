package main

import "fmt"

//scope is the range of places where you ar allowed to use a variable, type or function

// variable outside the main function. This means that other functions can access this variable
// outside variables can't be used with the := syntax. This is a global variable
var outsideVariable = "I'm outside the main function"

func main() {
	// insideVariable is a local variable.
	// It's visible only inside the main function between open and closed curly braces {}
	insideVariable := "I'm inside the main function"
	fmt.Println("From function main() =", outsideVariable)
	fmt.Println(insideVariable)
	f()
}

func f() {
	// function f can access the outsideVariable
	fmt.Println("From function f() =", outsideVariable)
}

func g() {
	// function g can't access the insideVariable
	// insideVariable is a local variable and can't be used outside the function
	//fmt.Println(insideVariable)
}
