# 01 · Concurrency Fundamentals

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

Concurrency is Go's signature feature and its most dangerous one. You will learn the primitives, the memory model, and the tooling that proves a design is correct.

## Prerequisites

- Beginner level completed.
- Comfort with closures, interfaces, and testing.

## Learning outcomes

- [ ] Design concurrent work with goroutines, channels, and synchronization.
- [ ] Explain happens-before and identify data races.
- [ ] Use the race detector and goroutine profiles to find bugs.
- [ ] Apply fan-out/fan-in, worker pool, and pipeline patterns.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 01.1 | [Goroutines](01-goroutines.md) | 90–120 min | ⬜ |
| 01.2 | [Channels](02-channels.md) | 90–120 min | ⬜ |
| 01.3 | [Select and Timeouts](03-select-and-timeouts.md) | 75–90 min | ⬜ |
| 01.4 | [Sync Primitives](04-sync-primitives.md) | 90 min | ⬜ |
| 01.5 | [The Race Detector](05-race-detector.md) | 60–75 min | ⬜ |
| 01.6 | [Concurrency Patterns](06-concurrency-patterns.md) | 90–120 min | ⬜ |

## Module checklist

- [ ] Read every topic note and explain it without notes.
- [ ] Implement every exercise in the module lab.
- [ ] Complete the [exercises](exercises.md).
- [ ] Record your solutions in [solutions.md](solutions.md).
- [ ] Build the mini project.
- [ ] Pass the self-assessment.
- [ ] Run `gofmt`, `go vet`, and `go test`.
- [ ] Update `PROGRESS.md`.

## Exercises

Do the exercises without notes first. Then compare with `solutions.md` and write down what you missed.

- [Exercises](exercises.md)
- [Solutions](solutions.md)

## Lab starter

Runnable starter code for every topic in this module:

[`labs/intermediate/01-concurrency-fundamentals`](../../../labs/intermediate/01-concurrency-fundamentals/README.md)

```bash
cd labs
go run ./intermediate/01-concurrency-fundamentals
go test ./intermediate/01-concurrency-fundamentals -v
```

## Mini project - Concurrent File Indexer

Build a CLI that walks a directory tree and computes per-file hashes concurrently. Bound the number of workers, support cancellation with Ctrl+C, and print deterministic results even though work finishes out of order.

**Deliverables:**

- [ ] Use a worker pool with a configurable limit.
- [ ] Propagate cancellation and clean up all workers.
- [ ] Preserve input order in the final report.
- [ ] Add tests and run them with `-race`.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What is a goroutine leak, and how do you detect one?
2. When should you close a channel, and who should do it?
3. Why does the race detector not prove the absence of races?
4. Compare a mutex and a channel for protecting a counter.

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go blog: share memory by communicating](https://go.dev/blog/codelab-share)
- [ ] [The Go memory model](https://go.dev/ref/mem)
- [ ] [Package sync](https://pkg.go.dev/sync)
- [ ] [Go blog: advanced concurrency patterns](https://go.dev/blog/advanced-go-concurrency-patterns)
