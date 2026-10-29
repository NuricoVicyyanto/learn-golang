package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Security"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"authentication-and-authorization", exercise01},
	{"owasp-for-apis", exercise02},
	{"secrets-and-tls", exercise03},
	{"supply-chain-security", exercise04},
}
