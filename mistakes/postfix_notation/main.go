package main

import "fmt"

func main() {
	i := 0
	// don't compile
	//fmt.Println(++i)
	//fmt.Println(i++)

	// Go increment and decrement operations can't be used as expressions, only as statements.
	// Also, only the 'postfix notation' is allowed.
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}
