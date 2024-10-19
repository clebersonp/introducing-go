package main

import "fmt"

func main() {
	avoidMistakes()
}

func avoidMistakes() {
	// a map needs to be initialized before use it, otherwise, it will result in a runtime error (panic).
	//var m map[string]float64

	m := make(map[string]float64)
	m["pi"] = 3.1416

	fmt.Println(m)
}
