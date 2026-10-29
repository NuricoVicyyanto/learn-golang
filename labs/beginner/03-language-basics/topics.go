package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Language Basics"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"variables-constants-types", exercise01},
	{"operators-and-expressions", exercise02},
	{"control-flow", exercise03},
	{"functions", exercise04},
	{"scope-and-shadowing", exercise05},
}
