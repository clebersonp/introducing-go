package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7}
	fmt.Println("Printing slice's index...")
	badRangeLoop(primes)
	fmt.Println()
	fmt.Println("Printing slice's elements...")
	goodRangeLoop(primes)
}

// badRangeLoop prints only the slice's index
func badRangeLoop(numbers []int) {
	// For arrays and slices, the range loop generates two values:
	// the first one will be the index and the second the element.
	// When range loop uses only 1 one value, it will be the index.
	for n := range numbers {
		fmt.Println(n)
	}
}

func goodRangeLoop(numbers []int) {
	for _, n := range numbers {
		fmt.Println(n)
	}
}
