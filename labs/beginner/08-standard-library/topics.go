package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Standard Library Essentials"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"strings-and-strconv", exercise01},
	{"time-and-math", exercise02},
	{"sort-and-slices", exercise03},
	{"os-and-io", exercise04},
}
