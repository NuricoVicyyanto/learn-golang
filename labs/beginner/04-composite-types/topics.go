package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Composite Types"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"arrays-and-slices", exercise01},
	{"maps", exercise02},
	{"structs", exercise03},
	{"pointers", exercise04},
}
