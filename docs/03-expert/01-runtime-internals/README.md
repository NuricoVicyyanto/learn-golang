# 01 · Runtime Internals

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

Expert Go work requires a model of what the runtime does: scheduling, memory visibility, garbage collection, and stack vs heap allocation. That model explains performance and correctness bugs that source code alone cannot.

## Prerequisites

- Intermediate level completed.
- Comfort with concurrency and profiling basics.

## Learning outcomes

- [ ] Explain the G-M-P scheduler and its latency implications.
- [ ] Reason about happens-before and the Go memory model.
- [ ] Describe GC phases and tune `GOGC` / `GOMEMLIMIT`.
- [ ] Use escape analysis to reduce allocations.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 01.1 | [Scheduler and GMP](01-scheduler-and-gmp.md) | 90–120 min | ⬜ |
| 01.2 | [The Memory Model](02-memory-model.md) | 90–120 min | ⬜ |
| 01.3 | [Garbage Collector](03-garbage-collector.md) | 90–120 min | ⬜ |
| 01.4 | [Escape Analysis](04-escape-analysis.md) | 75–90 min | ⬜ |

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

[`labs/expert/01-runtime-internals`](../../../labs/expert/01-runtime-internals/README.md)

```bash
cd labs
go run ./expert/01-runtime-internals
go test ./expert/01-runtime-internals -v
```

## Mini project - Runtime Lab

Create a small lab with reproducible experiments for scheduler behavior, GC tuning, and escape analysis. Each experiment should include a hypothesis, measurement, result, and explanation.

**Deliverables:**

- [ ] Build three isolated experiments with commands to reproduce.
- [ ] Capture traces or profiles as evidence.
- [ ] Write a one-page explanation of each result.
- [ ] List one production implication per experiment.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What are G, M, and P responsible for?
2. Why is a data race undefined behavior even if it seems to work?
3. What triggers a GC cycle, and what does `GOMEMLIMIT` change?
4. Why does returning an interface often allocate?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [The Go memory model](https://go.dev/ref/mem)
- [ ] [Go runtime source](https://github.com/golang/go/tree/master/src/runtime)
- [ ] [Go blog: scheduler tracing](https://go.dev/blog/execution-traces-2024)
- [ ] [Go docs: GC guide](https://go.dev/doc/gc-guide)
