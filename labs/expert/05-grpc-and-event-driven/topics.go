package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "gRPC and Event-Driven Systems"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"protobuf-and-grpc", exercise01},
	{"streaming", exercise02},
	{"kafka-or-nats", exercise03},
}
