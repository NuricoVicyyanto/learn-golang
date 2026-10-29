package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Packages and Modules"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"packages-and-visibility", exercise01},
	{"go-modules-and-semver", exercise02},
	{"internal-and-layout", exercise03},
}
