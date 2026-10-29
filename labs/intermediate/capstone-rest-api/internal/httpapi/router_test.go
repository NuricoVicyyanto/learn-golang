package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-golang/labs/intermediate/capstone-rest-api/internal/store"
)

func TestHealthz(t *testing.T) {
	t.Skip("TODO: write handler tests, then remove this Skip")

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	NewRouter(store.NewMemoryStore()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
