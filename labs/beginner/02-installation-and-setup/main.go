package main

import (
	"fmt"
	"os"
	"strings"
)

// greet returns a greeting for the given name.
//
// TODO: handle empty names gracefully.
func greet(name string) string {
	name = strings.TrimSpace(name)
	return "Hello, " + name + "!"
}

func main() {
	name := "Go"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greet(name))
}
