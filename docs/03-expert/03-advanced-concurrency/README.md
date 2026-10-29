# 03 · Advanced Concurrency

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

At scale, concurrency bugs become outages. You will design bounded, cancellable systems and learn to find the leaks and deadlocks that tests miss.

## Prerequisites

- Intermediate concurrency module completed.
- Module 01 of this level.

## Learning outcomes

- [ ] Design bounded worker pools and pipelines with clean shutdown.
- [ ] Apply backpressure and load shedding.
- [ ] Use `errgroup` and semaphores correctly.
- [ ] Detect deadlocks, leaks, and goroutine buildup.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 03.1 | [Worker Pools](01-worker-pools.md) | 90–120 min | ⬜ |
| 03.2 | [Pipelines](02-pipelines.md) | 90 min | ⬜ |
| 03.3 | [Backpressure and Load Shedding](03-backpressure.md) | 90 min | ⬜ |
| 03.4 | [errgroup and Semaphores](04-errgroup-and-semaphores.md) | 75–90 min | ⬜ |
| 03.5 | [Deadlocks and Leaks](05-deadlocks-and-leaks.md) | 90–120 min | ⬜ |

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

[`labs/expert/03-advanced-concurrency`](../../../labs/expert/03-advanced-concurrency/README.md)

```bash
cd labs
go run ./expert/03-advanced-concurrency
go test ./expert/03-advanced-concurrency -v
```

## Mini project - Bounded Job Runner

Build a reusable job runner with a fixed worker budget, per-job timeouts, retry policy, graceful shutdown, and metrics. It should survive one slow job without blocking the rest.

**Deliverables:**

- [ ] Define ownership and shutdown semantics in writing.
- [ ] Add per-job timeout and bounded retries.
- [ ] Expose queue depth, in-flight, failures, and latency.
- [ ] Prove no leaks with `-race` and a leak test.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is an unbounded queue not backpressure?
2. What does `errgroup.SetLimit` guarantee?
3. How do you prove a pipeline has no leaked goroutines?
4. What is a retry budget, and why does it matter under load?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [ ] [Package semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- [ ] [Go blog: advanced concurrency patterns](https://go.dev/blog/advanced-go-concurrency-patterns)
- [ ] [goleak](https://github.com/uber-go/goleak)
