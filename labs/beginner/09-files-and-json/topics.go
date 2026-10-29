package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Files and JSON"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"file-io", exercise01},
	{"encoding-json", exercise02},
	{"configuration", exercise03},
}
