package main

import "fmt"

func main() {
	const (
		century = 100
		decade  = 010
		year    = 001
	)

	// Why 8 instead of 10 for de const 'decade'?
	fmt.Println("century:", century, "decade:", decade, "year:", year)
	fmt.Println("She was", century+2*decade+2*year, "years old.")

	// 010 is a number in base 8, therefore it means 8, not 10.
	// Integer literals in Go are specified in octal, decimal or hexadecimal. The number 16 can be written as 020,
	// 16 or 0x10.
	// There is no decimal zero integer literal in Go.
}
