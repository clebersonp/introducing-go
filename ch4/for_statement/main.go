package main

import "fmt"

func main() {

	// Go only has one type of loop: the for statement
	// For loop can be used in a variety of different ways.

	// call the simple for
	simpleFor()

	fmt.Println()

	// call the for with all statements
	forWithAllStatements()

}

func simpleFor() {
	// declare a variable to control the loop
	i := 1

	// verifying condition. If is true it will be executed, otherwise don't
	for i <= 10 {
		// if the for conditional is true, this block will be executed
		fmt.Println(i)
		// increment the variable i to next loop. otherwise we will have an infinite loop
		i += 1
	}
}

func forWithAllStatements() {
	// variable initialization; condition; increment the variable
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
