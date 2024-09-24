package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	// Go has a built-in type for errors (the error type).
	// We can create our own errors by using the 'New' function in the errors package:
	err := errors.New("error message")
	fmt.Println("Type of err:", reflect.TypeOf(err), ", value:", err)

}
