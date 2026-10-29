package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "HTTP Clients and Servers"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"net-http-basics", exercise01},
	{"routing-and-middleware", exercise02},
	{"json-apis", exercise03},
	{"http-clients", exercise04},
}
