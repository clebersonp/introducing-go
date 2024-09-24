package main

import (
	"fmt"
	"time"
)

// Go has a special statement called select that works like a switch but for channels.

func main() {
	// create two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// goroutine for sending message to channel 1
	go func() {
		for {
			// send a message on channel c1
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	// goroutine for sending message to channel 2
	go func() {
		for {
			// send a message on channel c2
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	// goroutine for receiving messages from both channels c1 and c2 using select statement.
	// Will receive only one message at a time, the one that is ready first.
	// If more than one of the channels are ready, then it randomly picks which one to receive from.
	// If none of the channels are ready, the statement blocks until one becomes available.
	// The select statement is often used to implement a timeout
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				// We can also specify a default case:
				// The default case happens immediately if none of the channels are ready.
				//default:
				//	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	// force program to wait
	fmt.Scanln(&input)
}
