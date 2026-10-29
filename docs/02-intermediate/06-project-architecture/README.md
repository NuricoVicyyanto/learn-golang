# 06 · Project Architecture

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 10 hours

## Why this module

Architecture is the set of decisions that make change cheap or expensive. You will learn layout, dependency injection, configuration, and logging without over-engineering.

## Prerequisites

- Modules 04 and 05 completed.

## Learning outcomes

- [ ] Choose a layout that matches the project size.
- [ ] Wire dependencies explicitly from `main`.
- [ ] Load, validate, and document configuration.
- [ ] Use structured logging with request context.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 06.1 | [Layout and Packages](01-layout-and-packages.md) | 75–90 min | ⬜ |
| 06.2 | [Dependency Injection](02-dependency-injection.md) | 60–75 min | ⬜ |
| 06.3 | [Configuration](03-configuration.md) | 60–75 min | ⬜ |
| 06.4 | [Structured Logging](04-logging.md) | 60–75 min | ⬜ |

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

[`labs/intermediate/06-project-architecture`](../../../labs/intermediate/06-project-architecture/README.md)

```bash
cd labs
go run ./intermediate/06-project-architecture
go test ./intermediate/06-project-architecture -v
```

## Mini project — Service Skeleton

Create a reusable service skeleton with explicit wiring, validated configuration, structured logging, health endpoints, and a clean shutdown path. It should be small enough to read in one sitting.

**Deliverables:**

- [ ] Write `main` as wiring only.
- [ ] Load configuration from env and flags with clear precedence.
- [ ] Add request-ID-aware structured logs.
- [ ] Write a short ADR for the package layout.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is `main` a good place for wiring?
2. When does an interface at a dependency boundary pay off?
3. What configuration errors should stop startup?
4. What is the difference between a log and a metric?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Organizing a Go module](https://go.dev/doc/modules/layout)
- [ ] [Package log/slog](https://pkg.go.dev/log/slog)
- [ ] [Go blog: structured logging](https://go.dev/blog/slog)
- [ ] [Twelve-Factor App](https://12factor.net/)
