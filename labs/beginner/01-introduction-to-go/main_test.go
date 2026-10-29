package main

import "testing"

func TestProfile(t *testing.T) {
	t.Skip("TODO: fill in profile, then remove this Skip")

	lines := profile()
	if len(lines) != 5 {
		t.Fatalf("profile has %d lines, want 5", len(lines))
	}
	for i, line := range lines {
		if line == "" || len(line) < 10 {
			t.Errorf("line %d is too short or empty: %q", i, line)
		}
	}
}
