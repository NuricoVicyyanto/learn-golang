# 03 · Generics

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 8 hours

## Why this module

Generics remove duplication without sacrificing type safety. They are a tool for data structures and algorithms, not a license to abstract everything.

## Prerequisites

- Modules 01 and 02 completed.

## Learning outcomes

- [ ] Write generic functions and types with clear constraints.
- [ ] Understand type inference and zero values under generics.
- [ ] Design generic data structures that remain readable.
- [ ] Know when an interface is a better choice.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 03.1 | [Type Parameters](01-type-parameters.md) | 75–90 min | ⬜ |
| 03.2 | [Constraints](02-constraints.md) | 75–90 min | ⬜ |
| 03.3 | [Generic Data Structures](03-generic-data-structures.md) | 90 min | ⬜ |

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

[`labs/intermediate/03-generics`](../../../labs/intermediate/03-generics/README.md)

```bash
cd labs
go run ./intermediate/03-generics
go test ./intermediate/03-generics -v
```

## Mini project - Generic Collections Library

Build a small, well-documented generic collections package: set, stack, queue, and a `Result[T]` type. Focus on usable zero values, clear constraints, and tests that prove type safety at compile time.

**Deliverables:**

- [ ] Design APIs before implementation; write usage examples first.
- [ ] Keep constraints minimal and documented.
- [ ] Test with at least three concrete types.
- [ ] Write a short note on when callers should use interfaces instead.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. When is a generic function better than an interface?
2. What does the `~` operator do in a constraint?
3. Why can you not use `==` on every type parameter?
4. What is the zero value of `T`?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go blog: generics](https://go.dev/blog/intro-generics)
- [ ] [Go spec: type parameters](https://go.dev/ref/spec#Type_parameter_declarations)
- [ ] [When to use generics](https://go.dev/blog/when-generics)
- [ ] [Package slices](https://pkg.go.dev/slices)
