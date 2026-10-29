package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Runtime Internals"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"scheduler-and-gmp", exercise01},
	{"memory-model", exercise02},
	{"garbage-collector", exercise03},
	{"escape-analysis", exercise04},
}
