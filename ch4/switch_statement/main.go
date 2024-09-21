package main

func main() {
	printEnglishNumberNameIfStatement(7)
	printEnglishNumberNameSwitchStatement(3)

	for i := 1; i <= 10; i++ {
		printEvenOrOddWithSwitchStatement(i)
	}

	printVowelOrConsonantWithSwitchStatement('a')
	printVowelOrConsonantWithSwitchStatement('O')
	printVowelOrConsonantWithSwitchStatement('z')
}

func printEnglishNumberNameIfStatement(i int) {
	if i == 0 {
		println("zero")
	} else if i == 1 {
		println("one")
	} else if i == 2 {
		println("two")
	} else if i == 3 {
		println("three")
	} else if i == 4 {
		println("four")
	} else if i == 5 {
		println("five")
	} else if i == 6 {
		println("six")
	} else if i == 7 {
		println("seven")
	} else if i == 8 {
		println("eight")
	} else if i == 9 {
		println("nine")
	} else if i == 10 {
		println("ten")
	} else {
		println("unknown")
	}
}

func printEnglishNumberNameSwitchStatement(i int) {
	// switch statement starts with the keyword 'switch' followed by an expression (in this case, i)
	// and then a series of cases.
	// In Go don't need the 'break' statement
	switch i {
	case 0:
		// this statement only will be executed if 'case 0' is true (i == 0)
		println("zero")
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	case 4:
		println("four")
	case 5:
		println("five")
	case 6:
		println("six")
	case 7:
		println("seven")
	case 8:
		println("eight")
	case 9:
		println("nine")
	case 10:
		println("ten")
	default:
		// switch supports default case that will happen if none of the cases matches the value (similar to 'else')
		println("unknown")
	}
}

func printEvenOrOddWithSwitchStatement(i int) {
	switch i % 2 {
	case 0:
		println(i, "is even")
	default:
		println(i, "is odd")
	}
}

func printVowelOrConsonantWithSwitchStatement(c rune) {
	switch c {
	// we can group many cases together with the same output
	case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
		println(string(c), "is a vowel")
	default:
		println(string(c), "is a consonant")
	}
}
