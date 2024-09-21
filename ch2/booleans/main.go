package main

import "fmt"

func main() {

	// Three logical operators
	// && (AND)
	// || (OR)
	// ! (NOT)

	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("false && true =", false && true)
	fmt.Println("false && false =", false && false)
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println("false || true =", false || true)
	fmt.Println("false || false =", false || false)
	fmt.Println("!true =", !true)
	fmt.Println("!false =", !false)

	fmt.Println("(true && false) || (false && true) || !(false && false) =", (true && false) || (false && true) || !(false && false))
}
