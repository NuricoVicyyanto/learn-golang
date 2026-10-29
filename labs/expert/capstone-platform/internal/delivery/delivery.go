package delivery

import (
	"context"
	"errors"
	"fmt"

	"learn-golang/labs/expert/capstone-platform/internal/store"
)

// Worker delivers pending events to endpoints.
type Worker struct {
	store   store.Store
	workers int
}

func NewWorker(store store.Store, workers int) *Worker {
	return &Worker{store: store, workers: workers}
}

// Run processes deliveries until the context is cancelled.
//
// TODO: claim batches, deliver with timeout, retry with jitter, and handle DLQ.
func (w *Worker) Run(ctx context.Context) error {
	return errors.New("TODO: implement worker loop")
}

// Deliver sends one signed payload.
//
// TODO: sign the body, set headers, enforce a timeout, and classify retryable errors.
func Deliver(ctx context.Context, endpointURL, secret string, payload []byte) error {
	return fmt.Errorf("TODO: deliver to %s", endpointURL)
}
