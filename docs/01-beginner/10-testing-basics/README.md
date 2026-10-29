# 10 · Testing Basics

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 8 hours

## Why this module

Tests are executable documentation and your safety net for refactoring. Go's testing tools are built in; you only need discipline and good test design.

## Prerequisites

- Module 09 completed.
- A project with at least three packages.

## Learning outcomes

- [ ] Write focused unit tests with the `testing` package.
- [ ] Use table-driven tests and subtests.
- [ ] Measure coverage and interpret it honestly.
- [ ] Write a benchmark and read the result.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 10.1 | [Your First Test](01-first-test.md) | 60 min | ⬜ |
| 10.2 | [Table-Driven Tests](02-table-driven-tests.md) | 75–90 min | ⬜ |
| 10.3 | [Coverage and Benchmarks](03-coverage-and-benchmarks.md) | 75–90 min | ⬜ |

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

[`labs/beginner/10-testing-basics`](../../../labs/beginner/10-testing-basics/README.md)

```bash
cd labs
go run ./beginner/10-testing-basics
go test ./beginner/10-testing-basics -v
```

## Mini project — Test Suite for the Task Store

Take the JSON task store from module 09 and give it a real test suite: unit tests for business rules, filesystem tests using `t.TempDir`, and a benchmark for load and save. Measure coverage and document one gap you intentionally accept.

**Deliverables:**

- [ ] Use `t.TempDir` for all filesystem tests.
- [ ] Cover corrupt input, empty input, and duplicate IDs.
- [ ] Add a benchmark for a 10,000-task file.
- [ ] Write a short note on what coverage misses.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. When should a test use `t.Fatal` instead of `t.Error`?
2. Why is high coverage not proof of correctness?
3. What does `b.ResetTimer` do?
4. How do you make a table-driven test parallel safely?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package testing](https://pkg.go.dev/testing)
- [ ] [Go blog: table-driven tests](https://go.dev/blog/subtests)
- [ ] [Go docs: coverage](https://go.dev/doc/build-cover)
- [ ] [Go by Example: testing](https://gobyexample.com/testing)
