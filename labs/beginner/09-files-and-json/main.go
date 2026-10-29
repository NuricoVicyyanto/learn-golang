package main

import (
	"errors"
	"fmt"
	"os"
)

var errNotImplemented = errors.New("not implemented")

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Lab: %s\n\n", labTitle)
		fmt.Println("Usage: go run . <exercise>")
		fmt.Println()
		fmt.Println("Exercises:")
		for _, t := range topics {
			fmt.Printf("  %-42s\n", t.name)
		}
		return
	}

	for _, t := range topics {
		if t.name == os.Args[1] {
			if err := t.run(); err != nil {
				fmt.Fprintf(os.Stderr, "exercise %q: %v\n", t.name, err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Fprintf(os.Stderr, "unknown exercise %q\n", os.Args[1])
	fmt.Fprintln(os.Stderr, "Run without arguments to list exercises.")
	os.Exit(2)
}
