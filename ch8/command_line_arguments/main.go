package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// The flag package allows us to parse arguments and flags sent to our program.
// Example to run:
// go run ./ch8/command_line_arguments/main.go -max 20 list of arguments

func main() {
	// define flags
	maxp := flag.Int("max", 6, "the max value")
	// parse
	flag.Parse()
	// generate a number between 0 and max
	fmt.Println("Max flag:", *maxp, "random:", rand.Intn(*maxp))

	fmt.Println("non-flags arguments:", flag.Args())

}
