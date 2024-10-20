package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// the problem of data race only occurs before Go 1.21 version.
		// See the proposal: https://go.dev/doc/go1.21
		// now we don't need to create a new variable per-iteration, Go will do that for us.
		// Before the change the for loop creates variables per-loop. What can cause data race because all iterations
		// share the same address variable per-loop.
		go func() {
			defer wg.Done()
			fmt.Print(i)
			fmt.Printf("%v:%d, ", &i, i)
		}()
	}
	wg.Wait()
	fmt.Println()
}
