package main

import "fmt"

// Panic function causes a runtime error.
// Panic generally indicates a programmer error or an exceptional condition that there's no easy way to recover from
// We can handle a runtime panic with the built-in recover function.
// Recover stops the panic and returns the value that was passed to the call to panic.

func main() {
	// version 1
	//panic("PANIC")
	//str := recover() // this will never happen, because panic immediately stops execution of the function
	//fmt.Println(str)

	// version 2
	// This way will run normally.
	defer func() {
		str := recover()
		fmt.Println("Recovering str:", str)
	}()
	panic("PANIC")
}
