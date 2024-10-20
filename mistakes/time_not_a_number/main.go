package main

import (
	"fmt"
	"time"
)

func main() {
	n := 100
	fmt.Printf("Type n: %[1]T, n: %[1]d\n", n)

	// Why is this not compile?
	//time.Sleep(n * time.Second)
	// You can only multiply a time.Duration with another time.Duration
	// or an untyped integer constant
	var d time.Duration = 2000
	fmt.Printf("Type d: %[1]T, d: %[1]d\n", d)
	time.Sleep(d * time.Millisecond)

	fmt.Println("OR")
	// OR
	time.Sleep(2000 * time.Millisecond)

	fmt.Println("OR")
	// OR
	const e = 2000
	fmt.Printf("Type of constant e: %[1]T, e: %[1]d\n", e)
	time.Sleep(e * time.Millisecond)

	fmt.Println("OR")
	// OR
	f := time.Duration(2000)
	fmt.Printf("Type f: %[1]T, f: %[1]d\n", f)
	time.Sleep(f * time.Millisecond)
}
