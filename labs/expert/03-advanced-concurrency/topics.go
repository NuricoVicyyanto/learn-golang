package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Advanced Concurrency"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"worker-pools", exercise01},
	{"pipelines", exercise02},
	{"backpressure", exercise03},
	{"errgroup-and-semaphores", exercise04},
	{"deadlocks-and-leaks", exercise05},
}
