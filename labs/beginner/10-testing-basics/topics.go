package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Testing Basics"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"first-test", exercise01},
	{"table-driven-tests", exercise02},
	{"coverage-and-benchmarks", exercise03},
}
