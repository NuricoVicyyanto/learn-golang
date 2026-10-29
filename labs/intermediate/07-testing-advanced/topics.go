package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Advanced Testing"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"test-doubles", exercise01},
	{"httptest", exercise02},
	{"integration-tests", exercise03},
	{"fuzzing", exercise04},
}
