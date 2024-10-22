package main

import "fmt"

// A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.
func deadlock() {
	ch := make(chan int)
	ch <- 1           // goroutine 1 send the value one to the ch 'channel'
	fmt.Println(<-ch) // this line never reached because it is waiting for a channel (the above line will be stuck)
	// the result will be: all goroutines are asleep - deadlock!
	// channel only works to synchronize and or sharing value between 2 or more goroutines.
}

func main() {
	deadlock()
}
