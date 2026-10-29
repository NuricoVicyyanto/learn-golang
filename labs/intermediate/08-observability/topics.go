package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Observability"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"structured-logging", exercise01},
	{"metrics", exercise02},
	{"tracing", exercise03},
	{"health-checks", exercise04},
}
