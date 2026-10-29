# 06 · Errors and Panics

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 8 hours

## Why this module

Error handling is a core Go design decision, not a nuisance. You will learn to return, wrap, inspect, and present errors, and to reserve panic for truly unrecoverable situations.

## Prerequisites

- Module 05 completed.
- Comfort with functions and interfaces.

## Learning outcomes

- [ ] Return and check errors idiomatically.
- [ ] Wrap errors with context using `%w`.
- [ ] Use `errors.Is` and `errors.As` correctly.
- [ ] Define custom error types and recover at process boundaries.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 06.1 | [The Error Interface](01-the-error-interface.md) | 60–75 min | ⬜ |
| 06.2 | [Wrapping and Sentinel Errors](02-wrapping-and-sentinel-errors.md) | 75–90 min | ⬜ |
| 06.3 | [Custom Errors](03-custom-errors.md) | 60 min | ⬜ |
| 06.4 | [Panic and Recover](04-panic-and-recover.md) | 60–75 min | ⬜ |

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

[`labs/beginner/06-errors-and-panics`](../../../labs/beginner/06-errors-and-panics/README.md)

```bash
cd labs
go run ./beginner/06-errors-and-panics
go test ./beginner/06-errors-and-panics -v
```

## Mini project - Resilient CSV Importer

Build a CSV importer that validates rows, reports all row-level errors with line numbers, and continues processing valid rows. Distinguish parse errors, validation errors, and I/O errors so the caller can react differently.

**Deliverables:**

- [ ] Define typed errors for parse, validation, and I/O failures.
- [ ] Return a report containing valid rows and per-row failures.
- [ ] Use `errors.Is`/`errors.As` in tests to prove the error contract.
- [ ] Add a recover boundary only where a panic would otherwise crash the CLI.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is `error` an interface rather than a special type?
2. When should you wrap an error?
3. What is the difference between `errors.Is` and `errors.As`?
4. Name three situations where panic is acceptable.

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go blog: error handling and Go](https://go.dev/blog/error-handling-and-go)
- [ ] [Working with errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [ ] [Go by Example: errors](https://gobyexample.com/errors)
- [ ] [The Go specification: errors](https://go.dev/ref/spec#Errors)
