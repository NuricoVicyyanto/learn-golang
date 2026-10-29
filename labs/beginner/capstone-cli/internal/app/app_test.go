package app

import "testing"

type fakeStore struct {
	tasks []Task
	err   error
}

func (f *fakeStore) Load() ([]Task, error) { return f.tasks, f.err }
func (f *fakeStore) Save(tasks []Task) error {
	f.tasks = tasks
	return f.err
}

func TestAdd(t *testing.T) {
	t.Skip("TODO: implement add, then remove this Skip")

	st := &fakeStore{}
	a := New(st)

	if err := a.Run([]string{"add", "Learn Go"}); err != nil {
		t.Fatal(err)
	}
	if len(st.tasks) != 1 {
		t.Fatalf("got %d tasks, want 1", len(st.tasks))
	}
}
