package main

import (
	"fmt"
	"strings"
)

// To show doc from command line:
// go doc strings Contains

func main() {
	// Contains function
	test := "test"
	fmt.Println(strings.Contains(test, "es"))

	// Count
	fmt.Println(strings.Count(test, "t"))

	// HasPrefix
	fmt.Println(strings.HasPrefix(test, "te"))

	// HasSuffix
	fmt.Println(strings.HasSuffix(test, "st"))

	// Index
	fmt.Println(strings.Index(test, "e"))

	// Join
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	// Repeat
	fmt.Println(strings.Repeat("a", 5))

	// Replace
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))

	// Split
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// ToLower
	fmt.Println(strings.ToLower("TEST"))

	// ToUpper
	fmt.Println(strings.ToUpper(test))

	// String to []byte
	fmt.Println([]byte(test))

	// []byte to String
	fmt.Println(string([]byte{'t', 'e', 's', 't'}))

}
