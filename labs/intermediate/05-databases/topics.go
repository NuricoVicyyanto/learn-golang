package main

// topic is one runnable lab exercise.
type topic struct {
	name string
	run  func() error
}

const labTitle = "Databases"

// topics is the exercise registry used by main.go.
// Add a new exercise by adding a file and registering it here.
var topics = []topic{
	{"database-sql", exercise01},
	{"migrations", exercise02},
	{"transactions", exercise03},
	{"repository-pattern", exercise04},
}
