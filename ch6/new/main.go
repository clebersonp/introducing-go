package main

import (
	"fmt"
	"reflect"
)

// another way to get a pointer is to use the built-in 'new' function
// In Go is a garbage-collected programming language,
// which means memory is cleaned up automatically when nothing refers to it anymore.
// Pointers are extremely useful when paired with 'structs'.

func one(x *int) {
	*x = 1
}

func main() {
	x := new(int)
	one(x)
	fmt.Println("x is type of", reflect.TypeOf(x), ", address memory:", &x, ", pointer address:", x, ", value itself:", *x)
}
