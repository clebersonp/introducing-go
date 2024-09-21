package main

import "fmt"

// other functions can use it
const outsidePi float64 = 3.14159

func main() {

	// constants are variables whose values cannot be changed
	// we use the keyword 'const' to declare constant instead of 'var'
	// constants can be declared at the package level
	const insidePi float64 = 3.14159
	fmt.Println(insidePi)
	f()
	g()
}

func f() {
	// insidePi is out of scope
	// fmt.Println(insidePi)
}

func g() {
	// outsidePi is in scope
	fmt.Println(outsidePi)
}
