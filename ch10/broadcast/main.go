package main

import (
	"fmt"
	"time"
)

// Print text after the given time has expired.
// When done, the wait channel is closed to communicate receiver
func publish(text string, delay time.Duration) (wait <-chan struct{}) {
	// Empty structs indicates that the channel will only be used for signalling, not for passing data.
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		// All readers receive zero values on a closed channel
		close(ch)
	}()
	return ch
}

func main() {
	wait := publish("Channels let goroutines communicate.", 5*time.Second)
	fmt.Println("Waiting for news...")
	<-wait
	fmt.Println("Time to leave.")
}
