package main

import "fmt"

// Length calculate the size of given string and return it.
func Length(s string) int {
	return len(s)
}

// To see documentation in the terminal:
// go doc fmt Println
// go doc github.com/clebersonp/introducing-go/ch8/documentation Length
// This documentation is also available in web form by running this command:
// godoc -http=":6060"
// and open the browser at http://localhost:6060
// Ctrl+C in terminal to exit

func main() {
	s := "Hello, World!"
	fmt.Printf("Length of text '%s' is %d", s, Length(s))
}
