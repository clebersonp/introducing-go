package main

import "fmt"

func main() {

	// copy takes two arguments: dst and src.
	// All entries in src are copied into dst overwriting whatever is there.
	// If the lengths of the two slices are not the same, the smaller of the two will be used.

	// Copying slices
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("Slice1. Length:", len(slice1), "Capacity:", cap(slice1), "Content:", slice1)
	fmt.Println("Slice2. Length:", len(slice2), "Capacity:", cap(slice2), "Content:", slice2)

}
