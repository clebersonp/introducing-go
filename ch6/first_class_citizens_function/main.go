package main

import (
	"fmt"
)

// In a given programming language design, a first-class citizen is an entity which supports all the operations
// generally available to other entities. These operations typically include being passed as an argument,
// returned from a function, and assigned to a variable

// biOp is a type of function
// We can create functions as type
type biOp func(int, int) int

func main() {
	// we can declare variable of function type
	var op biOp
	add := func(a, b int) int {
		return a + b
	}
	op = add
	result := op(100, 200) // 100 + 200
	fmt.Println(result)

	// Subtract operation
	subtract := func(a, b int) int {
		return a - b
	}
	// we can pass function as argument
	result = perform(200, 100, subtract) // 200 - 100
	fmt.Println(result)

	// we can pass function as anonymous argument
	dividedOperation := perform(200, 10, func(a, b int) int { return a / b }) // 200 / 10
	fmt.Println(dividedOperation)

	// we can return function from another function
	multiply := preparedMultiply(10, 20)
	fmt.Printf("Type: %T, Result: %d\n", multiply, multiply()) // Type: func() int, Result: 200, multiply() to call
}

func perform(a, b int, operation biOp) int {
	// we can pass function as argument
	return operation(a, b)
}

func preparedMultiply(a, b int) func() int {
	// we can return a function from another function
	return func() int { return a * b }
}
