package main

import (
	"fmt"
	"math/rand"
	"time"
)

// It's a function that is capable of running concurrently with other functions.
// To create a goroutine, we use the keyword 'go' followed by a 'function invocation'

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {

	// This program consists of two goroutines.
	// The first one calls the function 'f' and the second one waits for the user to press enter (function main)
	go f(0)

	// Goroutines are lightweight, and we can easily create thousands of them. We can modify program to run
	// 10 goroutines concurrently.

	for i := 1; i < 11; i++ {
		go f(i)
	}

	var input string
	// Without this line, the program will not wait for printing all numbers from goroutine f
	fmt.Scanln(&input)
}
