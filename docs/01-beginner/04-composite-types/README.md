# 04 · Composite Types

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 12 hours

## Why this module

Real programs group data. Slices, maps, structs, and pointers are the vocabulary of Go data modeling, and their copy and aliasing rules explain many surprising bugs.

## Prerequisites

- Module 03 completed.
- Comfort with functions and loops.

## Learning outcomes

- [ ] Choose arrays, slices, maps, structs, or pointers for a data model.
- [ ] Explain slice headers, aliasing, capacity, and append behavior.
- [ ] Use maps safely, including the comma-ok idiom.
- [ ] Know when a pointer is appropriate and when a value is simpler.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 04.1 | [Arrays and Slices](01-arrays-and-slices.md) | 90–120 min | ⬜ |
| 04.2 | [Maps](02-maps.md) | 75–90 min | ⬜ |
| 04.3 | [Structs](03-structs.md) | 75–90 min | ⬜ |
| 04.4 | [Pointers](04-pointers.md) | 75–90 min | ⬜ |

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

[`labs/beginner/04-composite-types`](../../../labs/beginner/04-composite-types/README.md)

```bash
cd labs
go run ./beginner/04-composite-types
go test ./beginner/04-composite-types -v
```

## Mini project — Contact Book

Build a command-line contact book backed by in-memory structs and maps. Support add, search, update, delete, and list. Keep storage behind a small interface so it can later be replaced by a JSON file.

**Deliverables:**

- [ ] Define a `Contact` struct with validation rules.
- [ ] Use a map keyed by ID and keep list output sorted.
- [ ] Return typed errors for duplicate, missing, and invalid contacts.
- [ ] Write tests for each operation and for empty input.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What are the three fields of a slice header?
2. When does `append` allocate a new array?
3. Why is map iteration order not guaranteed?
4. When should a method or field use a pointer receiver?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go slices: usage and internals](https://go.dev/blog/slices-intro)
- [ ] [Go maps in action](https://go.dev/blog/maps)
- [ ] [Go by Example: slices, maps, structs](https://gobyexample.com/)
- [ ] [The Go specification: composite types](https://go.dev/ref/spec#Composite_types)
