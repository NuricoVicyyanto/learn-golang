package domain

import "time"

// Tenant owns endpoints and events.
type Tenant struct {
	ID     string
	Name   string
	APIKey string
}

// Endpoint receives webhook deliveries.
type Endpoint struct {
	ID       string
	TenantID string
	URL      string
	Secret   string
	Active   bool
}

// Event is an accepted event waiting for delivery.
type Event struct {
	ID             string
	TenantID       string
	Type           string
	Payload        []byte
	IdempotencyKey string
	CreatedAt      time.Time
}

// Delivery tracks one event-to-endpoint attempt sequence.
type Delivery struct {
	ID         string
	EventID    string
	EndpointID string
	Status     string
	Attempts   int
	NextTryAt  time.Time
}
