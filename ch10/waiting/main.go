package main

import (
	"fmt"
	"sync"
	"time"
)

// A sync.WaitGroup waits for a group of goroutines to finish.
// A WaitGroup must not be copied after first use.
func waitingGoroutines() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Done one goroutine.")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done another goroutine.")
		wg.Done()
	}()
	wg.Wait() // wait is used to block until all goroutines have finished. (called wg.Done())
	fmt.Println("Done all goroutines...")
}

func main() {
	waitingGoroutines()
}
