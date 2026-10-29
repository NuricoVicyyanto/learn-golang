package main

import (
	"fmt"
	"os"

	"learn-golang/labs/beginner/capstone-cli/internal/app"
	"learn-golang/labs/beginner/capstone-cli/internal/store"
)

func main() {
	st := store.NewJSONStore("tasks.json")
	application := app.New(st)

	if err := application.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
