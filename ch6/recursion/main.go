package main

import "fmt"

func main() {
	x := uint(5)
	fmt.Println("Factorial of", x, "is equal to", factorial(x))
}

// factorial is a function that call itself to evaluate the factorial number x
// Closure and recursion form the basis of a paradigm known as 'function programming'.
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
