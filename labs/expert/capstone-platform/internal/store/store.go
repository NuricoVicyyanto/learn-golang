package store

import (
	"context"
	"errors"

	"learn-golang/labs/expert/capstone-platform/internal/domain"
)

// Store is the persistence boundary for the platform.
type Store interface {
	CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error)
	ClaimDeliveries(ctx context.Context, limit int) ([]domain.Delivery, error)
	RecordAttempt(ctx context.Context, deliveryID string, status int, err error) error
	MoveToDeadLetter(ctx context.Context, deliveryID string, reason string) error
}

// MemoryStore is a starting point for the walking skeleton.
type MemoryStore struct{}

func NewMemoryStore() *MemoryStore { return &MemoryStore{} }

func (s *MemoryStore) CreateEvent(ctx context.Context, event domain.Event) (domain.Event, error) {
	return domain.Event{}, errors.New("TODO: implement CreateEvent with idempotency")
}

func (s *MemoryStore) ClaimDeliveries(ctx context.Context, limit int) ([]domain.Delivery, error) {
	return nil, errors.New("TODO: implement ClaimDeliveries")
}

func (s *MemoryStore) RecordAttempt(ctx context.Context, deliveryID string, status int, err error) error {
	return errors.New("TODO: implement RecordAttempt")
}

func (s *MemoryStore) MoveToDeadLetter(ctx context.Context, deliveryID string, reason string) error {
	return errors.New("TODO: implement MoveToDeadLetter")
}
