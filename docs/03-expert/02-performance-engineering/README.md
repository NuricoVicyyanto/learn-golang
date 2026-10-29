# 02 · Performance Engineering

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

Performance work without measurement is guessing. You will learn to build trustworthy benchmarks, read profiles and traces, and prove that a change helped.

## Prerequisites

- Module 01 completed.
- A service or library you can profile.

## Learning outcomes

- [ ] Design benchmarks that measure the right thing.
- [ ] Read CPU, heap, goroutine, and block profiles.
- [ ] Use execution traces to diagnose latency.
- [ ] Optimize allocations with measured before/after results.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 02.1 | [Benchmarking](01-benchmarking.md) | 90–120 min | ⬜ |
| 02.2 | [pprof](02-pprof.md) | 90–120 min | ⬜ |
| 02.3 | [Execution Tracing](03-execution-tracing.md) | 75–90 min | ⬜ |
| 02.4 | [Allocation Optimization](04-allocation-optimization.md) | 90 min | ⬜ |

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

[`labs/expert/02-performance-engineering`](../../../labs/expert/02-performance-engineering/README.md)

```bash
cd labs
go run ./expert/02-performance-engineering
go test ./expert/02-performance-engineering -v
```

## Mini project - Performance Investigation

Take a deliberately slow service or library, establish a baseline, find the bottleneck with profiles, make one change at a time, and document the measured impact. Include a regression benchmark in CI.

**Deliverables:**

- [ ] Write a reproducible load or benchmark harness.
- [ ] Capture baseline profiles and traces.
- [ ] Make three measured improvements with before/after data.
- [ ] Add a benchmark that would catch a regression.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What makes a benchmark trustworthy?
2. When would you use a trace instead of a CPU profile?
3. How do you distinguish allocation cost from GC cost?
4. What is the risk of optimizing for a microbenchmark?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go docs: profiling](https://go.dev/doc/diagnostics)
- [ ] [Package runtime/pprof](https://pkg.go.dev/runtime/pprof)
- [ ] [Package runtime/trace](https://pkg.go.dev/runtime/trace)
- [ ] [Go blog: pprof](https://go.dev/blog/pprof)
