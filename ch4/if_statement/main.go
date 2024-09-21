package main

import "fmt"

func main() {
	evenOrOdd()
}

func evenOrOdd() {
	for i := 1; i <= 10; i++ {
		// if condition is true, then the block of code will be executed
		if i%2 == 0 {
			fmt.Println(i, "is even")
			// else block will be executed if the condition is false
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
