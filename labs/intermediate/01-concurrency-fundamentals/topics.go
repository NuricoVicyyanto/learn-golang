package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Concurrency Fundamentals"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"goroutines", exercise01},
	{"channels", exercise02},
	{"select-and-timeouts", exercise03},
	{"sync-primitives", exercise04},
	{"race-detector", exercise05},
	{"concurrency-patterns", exercise06},
}
