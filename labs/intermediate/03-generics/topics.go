package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Generics"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"type-parameters", exercise01},
	{"constraints", exercise02},
	{"generic-data-structures", exercise03},
}
