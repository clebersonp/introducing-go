package main

import "fmt"

func main() {
	// Go has a special statement called 'defer' that schedules a function call to be run after the function completes.
	// will be invoked before the main function completes.
	// In this case, after first() call and before main function completes
	defer second()
	first() // will be invoked instantly

	// 'defer' is often used when resources need to be freed in some way. For example, when we open a file, we
	// need to make sure to close it later. Deferred functions are run even if a runtime panic occurs.
	// Example:
	// f, _ := os.Open(filename)
	// defer f.Close()
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
