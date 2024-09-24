package main

import (
	"fmt"
	"time"
)

// Channels provide a way for two goroutines to communicate with each other and synchronize their execution.
// A channel type is represented with the keyword 'chan' followed by the type of the things that are passed on the channel
// The left arrow operator (<-) is used to send and receive messages on the channel.
// c <- "ping" means send "ping" to channel c.
// msg := <-c means receive a message from channel c and assign the value to msg.
// Using channels like this synchronizes the two goroutines.
// When pinger attempts to send a message on the channel, it will lock until printer is ready to receive the message.
// (this is known as blocking).

func pinger(c chan<- string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
	}
}

// We can specify a direction on a channel type, thus restricting it to either sending or receiving.
// c <-chan string means that it's only a receiving channel.
// c chan<- string means that it's only a sending channel.
// c chan string means that it's both a sending and receiving channel (bidirectional).

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
