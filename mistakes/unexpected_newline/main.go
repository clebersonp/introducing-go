package main

import "fmt"

func main() {
	fruits := []string{
		"apple",
		"banana",
		"cherry", // missing the last comma. Every line must end with a comma.
	}
	fmt.Println(fruits)
}
