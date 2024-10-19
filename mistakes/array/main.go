package main

import "fmt"

func FooArray(a [2]int) {
	a[0] = 8
}

func FooSlice(a []int) {
	a[0] = 8
}

func main() {
	a := [2]int{1, 2}
	FooArray(a)    // Try to change a[0].
	fmt.Println(a) // Output: [1 2]
	// The array in Go is passed by value, so the change in FooArray is not reflected in the caller.

	// If you want FooArray to update the elements of a, use a slice instead.
	// A slice does not store any data, it just describes a section of an underlying array.
	// When you change an element of a slice, you modify the corresponding element of its underlying array,
	// and other slices that share the same underlying array will see the change.
	b := []int{1, 2}
	FooSlice(b)
	fmt.Println(b) // Output: [8 2]
}
