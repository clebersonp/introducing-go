package main

// Using packages

// mathy is the name of a package
// When we import our mathy library, we use its full name ("github.com/clebersonp/introducing-go/ch8/chapter8/mathy"),
// but inside the mathy.go file, we only use the last part of the name (package mathy)
// If we wanted to use alias for package name, we put an alias just before its full name on import statement.

import (
	"fmt"
	m "github.com/clebersonp/introducing-go/ch8/chapter8/mathy"
)

// Run from the root directory: go run ./ch8/chapter8/main.go

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	average := m.Average(xs)
	fmt.Println("Values:", xs, ", Average:", average)
}
