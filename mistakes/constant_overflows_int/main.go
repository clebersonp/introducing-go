package main

import "fmt"

func main() {
	// this code doesn't compile
	const n = 9876543210 * 9876543210 // overflows int
	// the untyped constant n must be converted to a type before it can be assigned to the interface{} (any) parameter
	// in the call to fmt.Println
	//fmt.Println(n)

	// When the type can't be inferred from the context, an untyped constant is converted to a bool, int, float64,
	// complex128, string or rune depends on the format of the constant.
	// In this case the constant is an integer, but n is larger than the maximum value of an int.
	// However, n can be represented as a float64
	fmt.Println(float64(n))

}
