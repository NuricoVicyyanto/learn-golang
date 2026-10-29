package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Methods and Interfaces"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"methods-and-receivers", exercise01},
	{"interfaces", exercise02},
	{"embedding-and-composition", exercise03},
	{"type-assertions-and-switches", exercise04},
}
