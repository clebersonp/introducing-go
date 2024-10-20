package main

import (
	"fmt"
	"math"
)

func main() {
	// a² + b² = c²
	// (3,4,5)
	fmt.Println("3²+4² == 5²", 3^2+4^2 == 5^2)

	// (6,8,10)
	fmt.Println("6²+8² == 10²", 6^2+8^2 == 10^2) // wrong, it's false. Why?

	// The circumflex ^ denotes bitwise XOR in Go. The computation written in base 2 looks like this:
	/*
		0011 ^ 0010 == 0001   (3^2 == 1)
		0100 ^ 0010 == 0110   (4^2 == 6)
		0101 ^ 0010 == 0111   (5^2 == 7)
	*/

	// To raise an integer to the power 2, use multiplication:
	fmt.Println("6²+8² == 10²", 6*6+8*8 == 10*10)

	// Or using function math.Pow
	fmt.Println("6²+8² == 10²", math.Pow(6, 2)+math.Pow(8, 2) == math.Pow(10, 2))
}
