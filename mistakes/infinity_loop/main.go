package main

import (
	"fmt"
)

func main() {
	//var b byte
	// Infinity loop. Why?
	// After the b == 255 iteration, b++ is executed. This overflows (since the maximum value for byte is 255)
	// and results in b == 0.
	//for b = 250; b <= 255; b++ {
	//	time.Sleep(300 * time.Millisecond)
	//	fmt.Printf("%[1]d %[1]c\n", b)
	//}

	// If we use the standard loop idiom with a strict inequality, the compiler will catch the bug.
	//for b = 250; b < 256; b++ {
	//	time.Sleep(300 * time.Millisecond)
	//	fmt.Printf("%[1]d %[1]c\n", b)
	//}

	// One solution is to use a wider data type, such as an int.
	for i := 250; i < 256; i++ {
		fmt.Printf("%[1]d %[1]c\n", i)
	}

	fmt.Println()

	// Another option is to put the termination test at the end of the loop.
	for b := byte(250); ; b++ {
		fmt.Printf("%[1]d %[1]c\n", b)
		if b == 255 {
			break
		}
	}
}
