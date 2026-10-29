# 03 · Language Basics

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 12 hours

## Why this module

This module turns syntax into fluency. You will learn how Go declares values, controls flow, and structures functions, so later modules are about design rather than deciphering the language.

## Prerequisites

- Modules 01 and 02 completed.
- Basic programming experience in any language.

## Learning outcomes

- [ ] Declare, initialize, and convert values with the right types.
- [ ] Write control flow without off-by-one or shadowing bugs.
- [ ] Use multiple returns, variadics, closures, and defer correctly.
- [ ] Explain zero values, scope, and evaluation order.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 03.1 | [Variables, Constants, and Types](01-variables-constants-types.md) | 75–90 min | ⬜ |
| 03.2 | [Operators and Expressions](02-operators-and-expressions.md) | 60–75 min | ⬜ |
| 03.3 | [Control Flow](03-control-flow.md) | 75–90 min | ⬜ |
| 03.4 | [Functions](04-functions.md) | 90–120 min | ⬜ |
| 03.5 | [Scope and Shadowing](05-scope-and-shadowing.md) | 60 min | ⬜ |

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

[`labs/beginner/03-language-basics`](../../../labs/beginner/03-language-basics/README.md)

```bash
cd labs
go run ./beginner/03-language-basics
go test ./beginner/03-language-basics -v
```

## Mini project — Command-Line Calculator

Build a calculator that parses simple expressions from the command line, supports `+ - * /`, handles division by zero, and prints a readable error for invalid input. Keep parsing and evaluation in separate functions so they can be tested independently.

**Deliverables:**

- [ ] Support one operation per invocation, e.g. `calc 12 / 3`.
- [ ] Return a non-zero exit code and a clear message on invalid input or division by zero.
- [ ] Separate parsing, evaluation, and output into testable functions.
- [ ] Add table-driven tests for valid and invalid expressions.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What is the zero value of a string, an int, a bool, a pointer, and a slice?
2. When must you use `var` instead of `:=`?
3. Why does `defer` evaluate its arguments immediately?
4. Explain a shadowing bug you wrote and how you found it.

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [The Go specification: declarations and types](https://go.dev/ref/spec)
- [ ] [A Tour of Go: basics](https://go.dev/tour/basics/1)
- [ ] [Effective Go](https://go.dev/doc/effective_go)
- [ ] [Go by Example](https://gobyexample.com/)
