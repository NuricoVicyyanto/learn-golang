package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Errors and Panics"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"the-error-interface", exercise01},
	{"wrapping-and-sentinel-errors", exercise02},
	{"custom-errors", exercise03},
	{"panic-and-recover", exercise04},
}
