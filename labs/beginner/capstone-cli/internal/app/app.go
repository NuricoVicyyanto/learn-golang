package app

import (
	"errors"
	"fmt"
)

// Task is the core domain type.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Store is the persistence boundary. Keep it small and consumer-defined.
type Store interface {
	Load() ([]Task, error)
	Save(tasks []Task) error
}

// App contains the command handlers.
type App struct {
	store Store
}

func New(store Store) *App {
	return &App{store: store}
}

// Run dispatches a command.
//
// TODO: support add, list, done, delete, stats, and help.
func (a *App) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("TODO: print help")
	}

	switch args[0] {
	case "add":
		return fmt.Errorf("TODO: implement add")
	case "list":
		return fmt.Errorf("TODO: implement list")
	case "done":
		return fmt.Errorf("TODO: implement done")
	case "delete":
		return fmt.Errorf("TODO: implement delete")
	case "stats":
		return fmt.Errorf("TODO: implement stats")
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}
