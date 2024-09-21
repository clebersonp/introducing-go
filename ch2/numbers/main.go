package main

import "fmt"

func main() {
	// three parts: the numeric 1 literal, the operator +, and another numeric literal 1
	// int numbers
	fmt.Println("1 + 1 =", 1+1)

	fmt.Println("1 + 1 =", 1.0+1.0)

	/*
		Go supports Standard arithmetic operators
		+	Addition
		-	Subtraction
		*	Multiplication
		/	Division
		%	Modulo/Remainder
	*/

	fmt.Println("32.132 * 42.452 =", 32.132*42.452)
}
