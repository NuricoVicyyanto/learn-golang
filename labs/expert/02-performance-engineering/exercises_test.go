package main

import "testing"

// These tests are intentionally skipped. For each exercise:
//   1. Implement the exercise function.
//   2. Remove its t.Skip.
//   3. Add real assertions for the behavior you implemented.
//   4. Run: go test ./... -v
//
// A skipped test is a reminder, not a passing grade.

// 01 - Benchmarking
func TestExercise01(t *testing.T) {
	t.Skip("TODO: implement exercise01, then remove this Skip")
	if err := exercise01(); err != nil {
		t.Fatalf("exercise01: %v", err)
	}
}

// 02 - pprof
func TestExercise02(t *testing.T) {
	t.Skip("TODO: implement exercise02, then remove this Skip")
	if err := exercise02(); err != nil {
		t.Fatalf("exercise02: %v", err)
	}
}

// 03 - Execution Tracing
func TestExercise03(t *testing.T) {
	t.Skip("TODO: implement exercise03, then remove this Skip")
	if err := exercise03(); err != nil {
		t.Fatalf("exercise03: %v", err)
	}
}

// 04 - Allocation Optimization
func TestExercise04(t *testing.T) {
	t.Skip("TODO: implement exercise04, then remove this Skip")
	if err := exercise04(); err != nil {
		t.Fatalf("exercise04: %v", err)
	}
}
