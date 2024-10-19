package main

import (
	"fmt"
	"time"
)

func main() {
	rand := make(chan int)

	go printBits(rand)

	for { // Send random sequence of bits to rand.
		select {
		case rand <- 0: // note: no statement
		case rand <- 1:
		}
	}
}

func printBits(c <-chan int) {
	for {
		// Receive random sequence of bits from c.
		time.Sleep(300 * time.Millisecond)
		fmt.Println(<-c)
	}
}
