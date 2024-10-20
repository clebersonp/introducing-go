package main

import "fmt"

func main() {
	fmt.Println("H" + "i") // Hi
	fmt.Println('H' + 'i') // 177

	// The rune 'H' and 'i' are integers values identifying Unicode code points: 'H' is 72 and 'i' is 105

	// You can turn a code point into a string with a conversion
	fmt.Println(string(72) + string('i'))

	// Or using fmt.Printf function
	fmt.Printf("%c%c, world", 72, 'i') // Hi, world
}
