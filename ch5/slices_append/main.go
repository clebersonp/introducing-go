package main

import "fmt"

func main() {

	// to add elements onto the end of a slice.
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Slice1,", "Length:", len(slice1), "Capacity:", cap(slice1), "Value:", slice1)
	fmt.Println("Slice2,", "Length:", len(slice2), "Capacity:", cap(slice2), "Value:", slice2)
}
