package main

import (
	"fmt"
	"github.com/clebersonp/introducing-go/ch10/server_channel/server"
	"time"
)

func main() {
	s := server.New()

	divideByZero := &server.Work{
		Op:    func(a, b int) int { return a / b },
		A:     100,
		B:     0,
		Reply: make(chan int),
	}
	// send work to the server channel
	s <- divideByZero

	select {
	// receive the result from the server by the Reply channel
	case res := <-divideByZero.Reply:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("No result in one second.")
	}
	// Output: No result in one second.
}
