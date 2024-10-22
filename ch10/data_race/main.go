package main

import "fmt"

// Data race happens when two goroutines access the same variable concurrently, and at least one of the accesses
// is a 'write'
func race() {
	wait := make(chan struct{})
	n := 0
	// goroutine 1
	go func() {
		// goroutine 2
		n++ // read, increment and write
		close(wait)
	}()
	n++ // conflicting access. read, increment and write in goroutine 1 while goroutine 2 is doing the same thing at same time
	<-wait
	fmt.Println(n) // Output: <unspecified>
}

// How to avoid data races
// The only way to avoid data races is to synchronize access to all mutable data tha is shared between threads.
// There are several ways to achieve this. In Go, you would normally use a 'channel' or a 'lock'.
// (Lower-lever mechanisms are available in the sync and sync/atomic packages.)
// The motto is: "Don't communicate by sharing memory; share memory by communicating"
func sharingIsCaring() {
	// channel passes the data from one goroutine to another, and it acts as a point of synchronization.
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		ch <- n // the data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another.
	n++
	fmt.Println(n) // Output: 2
}

func main() {
	// How to detect data race?
	// Run: go run -race main.go
	// Run: go test -race ./... (packages)
	//race()

	sharingIsCaring()
}
