package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Docker and Deployment"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"dockerfile", exercise01},
	{"compose", exercise02},
	{"ci-cd", exercise03},
}
