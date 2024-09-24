package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte{'t', 'e', 's', 't'})
	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())

	stdout := os.Stdout
	stdout.ReadFrom(strings.NewReader("another test"))
}
