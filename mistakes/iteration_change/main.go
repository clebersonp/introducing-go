package main

import "fmt"

func main() {
	// array
	var a [2]int
	for _, x := range a {
		// why the last x value has not changed even the statement a[1] = 8 already did?
		// The range expression 'a' is evaluated once before beginning the loop and a copy of the array is used
		// to generate the iteration values.
		// To avoid copying the array, iterate over a slice instead.
		fmt.Println("x = ", x)
		a[1] = 8
	}
	fmt.Println("a = ", a)

	for _, x := range a[:] {
		fmt.Println("x = ", x)
		a[1] = 16
	}
	fmt.Println("a = ", a)

}
