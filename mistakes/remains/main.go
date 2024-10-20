package main

import "fmt"

func main() {
	fmt.Println("BadOdd(-1):", BadOdd(-1))
	fmt.Println("GoodOdd1(-1):", GoodOdd1(-1))
	fmt.Println("GoodOdd2(-1):", GoodOdd2(-1))
}

func BadOdd(n int) bool {
	// The remainder operator can give negative answers if the dividend is negative.
	// If n is an odd negative number, n % 2 equals -1
	return n%2 == 1
}

func GoodOdd1(n int) bool {
	// Any remainder different of Zero will be an odd number
	return n%2 != 0
}
func GoodOdd2(n int) bool {
	// Odd number using bitwise & (AND) operator equals to 1
	/*
		-1 in binary is (1 sign)001
		& 1 in binary is (0 sign)001

			1001
		& 	0001
		---------------------------
			0001 = 1 in decimal

			0011 (3 in decimal)
		&	0001
			---------------------------
			0001 = 1 in decimal
	*/
	return n&1 == 1
}
