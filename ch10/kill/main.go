package main

import (
	"fmt"
	"time"
)

// To make a goroutine stoppable, let it listen for a stop signal on a dedicated quit channel (bool value),
// and check this channel at suitable points in your goroutine.
func killGoroutine() (wait <-chan struct{}) {
	quit := make(chan bool)
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("\nQuitting...")
				close(ch)
				return
			case <-time.After(time.Second):
				fmt.Print(".")
			}
		}
	}()
	<-time.After(10 * time.Second)
	quit <- true
	return ch
}

// generator returns a channel that produces the numbers 1, 2, 3, ...
// To stop the underlying goroutine, send a number on this channel.
// We use a single channel for both data and signalling.
func generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				fmt.Println("Stopping...")
				close(ch)
				return
			}
		}
	}()
	return ch
}

func main() {
	<-killGoroutine()
	fmt.Println("Do something...")

	// generator
	number := generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 0           // stops underlying goroutine
	fmt.Println(<-number) // zero-value (0) because the channel is closed.
	// If we don't close the channel, would be an error, all goroutines are asleep - deadlock
}
