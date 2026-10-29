package store

import (
	"context"
	"errors"
)

// Bookmark is the core domain type.
type Bookmark struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Store is the persistence boundary.
type Store interface {
	Create(ctx context.Context, b Bookmark) (Bookmark, error)
	Get(ctx context.Context, userID, id string) (Bookmark, error)
	List(ctx context.Context, userID string, limit int) ([]Bookmark, error)
	Update(ctx context.Context, b Bookmark) error
	Delete(ctx context.Context, userID, id string) error
}

// MemoryStore is a temporary implementation for the walking skeleton.
type MemoryStore struct {
	bookmarks map[string]Bookmark
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{bookmarks: make(map[string]Bookmark)}
}

func (s *MemoryStore) Create(ctx context.Context, b Bookmark) (Bookmark, error) {
	return Bookmark{}, errors.New("TODO: implement Create")
}

func (s *MemoryStore) Get(ctx context.Context, userID, id string) (Bookmark, error) {
	return Bookmark{}, errors.New("TODO: implement Get")
}

func (s *MemoryStore) List(ctx context.Context, userID string, limit int) ([]Bookmark, error) {
	return nil, errors.New("TODO: implement List")
}

func (s *MemoryStore) Update(ctx context.Context, b Bookmark) error {
	return errors.New("TODO: implement Update")
}

func (s *MemoryStore) Delete(ctx context.Context, userID, id string) error {
	return errors.New("TODO: implement Delete")
}
