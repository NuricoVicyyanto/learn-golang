# 08 · Standard Library Essentials

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 10 hours

## Why this module

The standard library is the default toolkit of every Go program. Knowing what exists prevents unnecessary dependencies and reinvention.

## Prerequisites

- Module 07 completed.

## Learning outcomes

- [ ] Use `strings`, `strconv`, `time`, `math`, `sort`, `slices`, `os`, and `io` confidently.
- [ ] Choose between `strings.Builder` and concatenation.
- [ ] Parse and format time and numbers correctly.
- [ ] Read from files, stdin, and generic readers.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 08.1 | [Strings and strconv](01-strings-and-strconv.md) | 75–90 min | ⬜ |
| 08.2 | [Time and Math](02-time-and-math.md) | 75–90 min | ⬜ |
| 08.3 | [Sort and Slices](03-sort-and-slices.md) | 60–75 min | ⬜ |
| 08.4 | [OS and I/O](04-os-and-io.md) | 75–90 min | ⬜ |

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

[`labs/beginner/08-standard-library`](../../../labs/beginner/08-standard-library/README.md)

```bash
cd labs
go run ./beginner/08-standard-library
go test ./beginner/08-standard-library -v
```

## Mini project - Log Analyzer

Build a CLI that reads a log file or stdin, parses timestamp and level, and reports counts per level, the busiest minute, and the slowest request. Stream input rather than loading everything into memory.

**Deliverables:**

- [ ] Support both a file argument and stdin.
- [ ] Parse without panicking on malformed lines; report skipped lines.
- [ ] Sort and format the report deterministically.
- [ ] Test with a golden file and malformed input.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. When should you use `strings.Builder`?
2. Why is a string index not a character index?
3. Why is `time.Time.Equal` safer than `==`?
4. How do you detect a scanner error?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package strings](https://pkg.go.dev/strings)
- [ ] [Package strconv](https://pkg.go.dev/strconv)
- [ ] [Package time](https://pkg.go.dev/time)
- [ ] [Package slices](https://pkg.go.dev/slices)
- [ ] [Package io](https://pkg.go.dev/io)
