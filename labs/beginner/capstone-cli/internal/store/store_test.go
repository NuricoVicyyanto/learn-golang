package store

import "testing"

func TestSaveAndLoad(t *testing.T) {
	t.Skip("TODO: implement Save and Load, then remove this Skip")

	path := t.TempDir() + "/tasks.json"
	s := NewJSONStore(path)
	if err := s.Save(nil); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Load(); err != nil {
		t.Fatal(err)
	}
}
