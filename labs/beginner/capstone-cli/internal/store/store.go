package store

import (
	"errors"

	"learn-golang/labs/beginner/capstone-cli/internal/app"
)

// JSONStore persists tasks in a JSON file.
type JSONStore struct {
	path string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

// Load reads tasks from disk.
//
// TODO: handle a missing file as an empty list, and reject corrupt JSON.
func (s *JSONStore) Load() ([]app.Task, error) {
	return nil, errors.New("TODO: implement Load")
}

// Save writes tasks atomically.
//
// TODO: write to a temporary file, sync, close, then rename.
func (s *JSONStore) Save(tasks []app.Task) error {
	return errors.New("TODO: implement Save")
}
