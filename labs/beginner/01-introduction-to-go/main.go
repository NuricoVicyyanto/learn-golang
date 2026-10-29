package main

import "fmt"

// profile returns five short statements about Go.
//
// TODO: replace the placeholders with your own summary.
func profile() []string {
	return []string{
		"TODO: what Go is",
		"TODO: why Go was created",
		"TODO: one strength",
		"TODO: one weakness",
		"TODO: one project where you would use Go",
	}
}

func main() {
	for _, line := range profile() {
		fmt.Println(line)
	}
}
