package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(s)

	// Go strings are immutable and behave like read-only byte slices
	//s[0] = 'H' // cannot assign to s[0]

	// To update the data, use a rune slice instead.
	buf := []rune("hello") // or []rune{'h', 'e', 'l', 'l', 'o'} or []rune(s)
	fmt.Println(buf)
	buf[0] = 'H'
	newS := string(buf)
	fmt.Println(newS) // "Hello"
}
