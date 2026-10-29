package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Distributed Systems"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"idempotency-and-retries", exercise01},
	{"timeouts-and-circuit-breakers", exercise02},
	{"message-queues", exercise03},
	{"transactional-outbox", exercise04},
}
