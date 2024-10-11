package main

import (
	"fmt"
	"time"
)

func main() {
	wait := publisher("Hello, World!", 1*time.Minute)
	fmt.Println("Waiting for message...")
	// Do some more work.
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Still waiting...")
		case <-wait: // only will receive the signal when the channel is closed
			fmt.Println("Message received!")
			return
		}
	}
}

func publisher(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}
