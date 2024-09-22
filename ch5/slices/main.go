package main

import "fmt"

func main() {
	// A slice is a segment of an array. Like arrays, slices are indexable and have a length.
	// Unlike arrays, this length is allowed to change. The only difference between this and an array is the missing
	// length between the brackets [].

	// To create a slice using built-in make function
	x := make([]float64, 5)
	fmt.Println("Length:", len(x), "Capacity:", cap(x), "Content:", x)

	// Slices are always associated with come array, and although they can never be longer than the array,
	// they can be smaller.
	y := make([]float64, 5, 10)
	fmt.Println("Length:", len(y), "Capacity:", cap(y), "Content:", y)

	// Another way to create slices is to use [low : high] expression
	arr := [5]float64{1, 2, 3, 4, 5}
	// start index : end index -1 (but not including the index itself)
	// we are also allowed to omit low, high, or even both low and high.
	// arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5], and
	// arr[:] is the same as arr[0:len(arr)]
	z := arr[0:5]
	fmt.Println("Length:", len(z), "Capacity:", cap(z), "Content:", z)
}
