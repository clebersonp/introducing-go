package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Why this regex matches a string with characters in it?
	// The function 'regexp.MatchString' does 'substring' matching.
	numberWithCharacters := "12three45"
	patternWithoutAnchor := `[0-9]*`
	matched, err := regexp.MatchString(patternWithoutAnchor, numberWithCharacters)
	if err != nil {
		panic(err)
	}
	fmt.Println(matched, err) // true

	// To check if a full string matches only [0-9]*, anchor the start and the end of the regular expression.
	patternWithAnchor := `^[0-9]*$`
	matched, err = regexp.MatchString(patternWithAnchor, numberWithCharacters)
	if err != nil {
		panic(err)
	}
	fmt.Println(matched, err) // false
}
