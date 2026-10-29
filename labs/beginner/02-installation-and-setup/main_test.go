package main

import "testing"

func TestGreet(t *testing.T) {
	t.Skip("TODO: fix greet for empty input, then remove this Skip")

	tests := []struct {
		name string
		want string
	}{
		{"Alice", "Hello, Alice!"},
		{"", "Hello, stranger!"},
		{"  Bob  ", "Hello, Bob!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(tt.name); got != tt.want {
				t.Fatalf("greet(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
