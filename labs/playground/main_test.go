package main

import "testing"

func TestGreeting(t *testing.T) {
	if got := greeting(); got == "" {
		t.Fatal("greeting must not be empty")
	}
}
