package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Data at Scale"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"caching", exercise01},
	{"query-tuning", exercise02},
	{"sharding-and-replication", exercise03},
}
