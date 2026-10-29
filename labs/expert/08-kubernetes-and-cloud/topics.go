package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Kubernetes and Cloud"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"containers-and-kubernetes", exercise01},
	{"graceful-shutdown-and-scaling", exercise02},
	{"observability-in-production", exercise03},
}
