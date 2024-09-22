package main

import "fmt"

// Pointers reference a location in memory where a value is stored rather than the value itself.

func main() {
	x := 5
	zeroWithoutPointer(x)
	fmt.Println("Calling zeroWithoutPointer function. Value:", x)

	// &x = accessing and pass the copy of address memory (pointer) instead of the value itself.
	// we use the & operator to find the address of a variable.
	// &x returns a *int (pointer to an int) because x is an int.
	zeroWithPointer(&x)
	fmt.Println("Calling zeroWithPointer function. Value:", x)

	fmt.Println("x value itself:", x, ", x memory address (pointer):", &x)
}

// When we call a function that takes an argument, that argument is copied to the function

// zeroWithoutPointer takes a copy from the caller and pass the copy to x argument.
// This case will receive a copy of value itself.
func zeroWithoutPointer(x int) {
	// this modification only impacts here. Because of variables scope.
	x = 0
}

// zeroWithPointer takes a copy from the caller and pass the copy to x argument.
// But, this time will receive a pointer and not a value itself
// By using a pointer (*int), the function is able to modify the original variable.
// Pointer is represented using an asterisk (*) followed by the type of the stored value.
// In this function, x is a pointer to an int
func zeroWithPointer(x *int) {
	// *x = accessing the value itself through the memory address (pointer)
	// Dereferencing a pointer gives us access to the value the pointer points to.
	// When we write *x = 0, we are saying 'store the int 0 in the memory location x refers to'.
	*x = 0
	fmt.Println("zeroWithPointer function. x value itself:", *x, ", x memory address:", &x, ", x pointer address:", x)
}
