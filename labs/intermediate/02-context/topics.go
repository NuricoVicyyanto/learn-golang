package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Context and Cancellation"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"context-basics", exercise01},
	{"cancellation-and-deadlines", exercise02},
	{"context-in-http-and-db", exercise03},
}
