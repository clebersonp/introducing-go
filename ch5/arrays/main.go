package main

import "fmt"

func main() {
	// arrays is a numbered sequence of elements of a single type with a fixed length
	// arrays are indexed starting from 0.
	var x [5]int
	x[4] = 100
	fmt.Println(len(x), x)
	// get the value of the array at index 4
	fmt.Println(x[4])

	var xs [5]float64
	xs[0] = 98
	xs[1] = 93
	xs[2] = 77
	xs[3] = 82
	xs[4] = 83
	result := average(xs)
	fmt.Println("Average:", result)

	// shorter syntax for creating arrays:
	ys := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	fmt.Println("Shorter syntax array:", ys)

	// Obs. Prefer slices to arrays. We rarely see arrays used directly in Go code.
	// Slice is a type built on top of an array
}

func average(xs [5]float64) float64 {
	var total float64 = 0

	// special for form
	// The Go won't allow you to create variable that you never use.
	// we need use _ to ignore the array index
	// range for array return two values: index and value
	for _, value := range xs {
		total += value
	}
	// float64(<int value>) to convert int to float64
	// in general, to convert between types, we use the type name like a function
	return total / float64(len(xs))
}
