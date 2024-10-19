package main

import (
	"fmt"
	"time"
)

func main() {
	// error: multiple-value time.Parse() in single-value context
	//t := time.Parse(time.RFC3339, "2012-11-01T22:08:41Z")

	// The time.Parse function returns two values, a time.Time and an error.
	t, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41Z")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// Blank identifier
	// You can use the blank identifier _ to ignore unwanted return values.
	const pi = "pi"
	m := map[string]float64{pi: 3.1416}
	_, exists := m[pi]
	fmt.Printf("Exists 'pi'?: %t\n", exists)
}
