package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Performance Engineering"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"benchmarking", exercise01},
	{"pprof", exercise02},
	{"execution-tracing", exercise03},
	{"allocation-optimization", exercise04},
}
