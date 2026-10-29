# 09 · Files and JSON

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 8 hours

## Why this module

Most programs need to persist or exchange data. Files and JSON are the simplest durable boundary, and they teach error handling, encoding rules, and configuration design.

## Prerequisites

- Module 08 completed.

## Learning outcomes

- [ ] Read and write files safely with proper cleanup.
- [ ] Marshal and unmarshal JSON with controlled contracts.
- [ ] Design configuration with precedence and validation.
- [ ] Handle malformed input without losing data.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 09.1 | [File I/O](01-file-io.md) | 75–90 min | ⬜ |
| 09.2 | [Encoding JSON](02-encoding-json.md) | 75–90 min | ⬜ |
| 09.3 | [Configuration](03-configuration.md) | 60–75 min | ⬜ |

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

[`labs/beginner/09-files-and-json`](../../../labs/beginner/09-files-and-json/README.md)

```bash
cd labs
go run ./beginner/09-files-and-json
go test ./beginner/09-files-and-json -v
```

## Mini project - JSON Task Store

Build a persistent task store backed by a JSON file. Support add, list, complete, and delete. Writes must be atomic, and a corrupt file must produce a clear error rather than data loss.

**Deliverables:**

- [ ] Define the on-disk schema and version it.
- [ ] Implement atomic save and safe load.
- [ ] Handle missing, empty, and corrupt files distinctly.
- [ ] Test round-trips and concurrent-safe access within one process.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is writing directly to the target file risky?
2. What does `omitempty` do, and when is it wrong?
3. How do you reject unknown JSON fields?
4. What configuration precedence would you choose for a CLI, and why?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package os](https://pkg.go.dev/os)
- [ ] [Package encoding/json](https://pkg.go.dev/encoding/json)
- [ ] [Go blog: JSON and Go](https://go.dev/blog/json)
- [ ] [Package bufio](https://pkg.go.dev/bufio)
