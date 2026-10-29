# Exercises - Concurrency Fundamentals

> **Level:** Intermediate · **Rule:** attempt without notes first. Record failures honestly in [solutions.md](solutions.md).

Work in `labs/`. Keep each exercise small enough to finish in one sitting. Write tests where behavior is non-trivial.

## 1. Build a worker pool that processes N jobs with a fixed number of workers, preserves result order, and shuts down cleanly on cancellation.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 2. Implement a pipeline with three stages. Cancel halfway through and prove there are no leaked goroutines using `pprof` or a leak test.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 3. Write a bounded parallel map function and benchmark it against a sequential version for CPU-bound and I/O-bound work.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 4. Rebuild from memory

Pick the topic you understood least. Close the notes and rebuild its main example from scratch. Compare with your note and record the gaps.

## 5. Test and explain

Add a table-driven test for one exercise above. Then write three sentences explaining why the test cases are sufficient - and one case you deliberately left out.

## Review

| Exercise | Attempts | Passed | Time | What I learned |
| --- | --- | --- | --- | --- |
| 1 |  |  |  |  |
| 2 |  |  |  |  |
| 3 |  |  |  |  |
| 4 |  |  |  |  |
| 5 |  |  |  |  |
