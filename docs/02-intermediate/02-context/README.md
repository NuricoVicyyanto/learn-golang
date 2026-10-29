# 02 · Context and Cancellation

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 8 hours

## Why this module

Production work must stop when the client disconnects, the deadline expires, or the process shuts down. `context` is the standard way to carry that signal.

## Prerequisites

- Module 01 completed.

## Learning outcomes

- [ ] Propagate context through every blocking call.
- [ ] Use deadlines and cancellation without leaking resources.
- [ ] Integrate context with HTTP and database calls.
- [ ] Test cancellation behavior.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 02.1 | [Context Basics](01-context-basics.md) | 60–75 min | ⬜ |
| 02.2 | [Cancellation and Deadlines](02-cancellation-and-deadlines.md) | 75–90 min | ⬜ |
| 02.3 | [Context in HTTP and Databases](03-context-in-http-and-db.md) | 75–90 min | ⬜ |

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

[`labs/intermediate/02-context`](../../../labs/intermediate/02-context/README.md)

```bash
cd labs
go run ./intermediate/02-context
go test ./intermediate/02-context -v
```

## Mini project — Deadline-Aware Aggregator

Build a small service that aggregates data from three upstreams. Each upstream has its own timeout, the request has an overall budget, and slow upstreams degrade gracefully instead of failing the whole response.

**Deliverables:**

- [ ] Implement a total request budget and per-upstream budgets.
- [ ] Return partial results with a structured warning.
- [ ] Cancel remaining work when the client disconnects.
- [ ] Test timeouts with controllable fake upstreams.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why should context be the first parameter?
2. What happens if you forget to call `cancel`?
3. How do you distinguish cancellation from a real error?
4. Where should timeouts be configured?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go blog: context](https://go.dev/blog/context)
- [ ] [Package context](https://pkg.go.dev/context)
- [ ] [Go blog: context and cancellation](https://go.dev/blog/context-and-structs)
- [ ] [net/http server timeouts](https://pkg.go.dev/net/http#Server)
