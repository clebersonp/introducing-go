package main

import (
	"fmt"
	"time"
)

// Timers and Ticker let you execute code in the future, once or repeatedly

func timeout(ch <-chan string) {
	for alive := true; alive; {
		timer := time.NewTimer(10 * time.Second)
		select {
		case news := <-ch:
			timer.Stop()
			fmt.Println(news)
		case <-timer.C:
			alive = false
			fmt.Println("No news in ten seconds. Service aborting.")
		}
	}
}

// time.Tick returns a channel that delivers clock ticks at even intervals
func repeat() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		ticker := time.NewTicker(time.Second)
		n := 0
		defer ticker.Stop()
		defer close(ch)
		for now := range ticker.C {
			n++
			fmt.Println(now, "doing something important")
			if n > 8 {
				break
			}
		}
		fmt.Println("Finished ticker")
	}()
	return ch
}

func foo() {
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("Foo run for more than 1 second.")
	})
	defer timer.Stop()
	fmt.Println("Do the last thing before finished")
	time.Sleep(3 * time.Second)
}

func main() {
	ch := make(chan string)
	go timeout(ch)
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprint("Message", i)
		time.Sleep(100 * time.Millisecond)
	}
	<-time.After(12 * time.Second)

	fmt.Println()
	<-repeat()

	fmt.Println()
	foo()
}
