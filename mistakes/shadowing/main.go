package main

import "fmt"

// Using Go tools to check 'shadowing issues' in the code
// go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
// go vet -vettool=$(which shadow)

func main() {
	shadowing()

	withoutShadowing()
}

func shadowing() {
	n := 0
	if true {
		// shadowing variable n and changing the n scope to the if block
		n := 1
		n++
	}
	fmt.Println(n)
}

func withoutShadowing() {
	n := 0
	if true {
		// reusing the variable n instead of creating a new one to the if block
		n = 1
		n++
	}
	fmt.Println(n)
}
