package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Project Architecture"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"layout-and-packages", exercise01},
	{"dependency-injection", exercise02},
	{"configuration", exercise03},
	{"logging", exercise04},
}
