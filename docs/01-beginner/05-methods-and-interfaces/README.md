# 05 · Methods and Interfaces

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 10 hours

## Why this module

Interfaces are Go's main abstraction tool. Unlike inheritance, they are satisfied implicitly, which makes APIs flexible but requires care around method sets and nil values.

## Prerequisites

- Module 04 completed.
- Comfort with structs and pointers.

## Learning outcomes

- [ ] Choose value or pointer receivers deliberately.
- [ ] Design small interfaces at the consumer side.
- [ ] Use embedding and composition instead of inheritance.
- [ ] Handle nil interfaces and type assertions safely.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 05.1 | [Methods and Receivers](01-methods-and-receivers.md) | 75–90 min | ⬜ |
| 05.2 | [Interfaces](02-interfaces.md) | 90–120 min | ⬜ |
| 05.3 | [Embedding and Composition](03-embedding-and-composition.md) | 60–75 min | ⬜ |
| 05.4 | [Type Assertions and Switches](04-type-assertions-and-switches.md) | 60–75 min | ⬜ |

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

[`labs/beginner/05-methods-and-interfaces`](../../../labs/beginner/05-methods-and-interfaces/README.md)

```bash
cd labs
go run ./beginner/05-methods-and-interfaces
go test ./beginner/05-methods-and-interfaces -v
```

## Mini project - Notification System

Build a small notification service with a `Notifier` interface and email, SMS, and log implementations. A dispatcher selects notifiers from configuration and records delivery results. Keep each notifier independently testable.

**Deliverables:**

- [ ] Define the smallest useful `Notifier` interface.
- [ ] Implement at least three notifiers, including one fake for tests.
- [ ] Add a dispatcher that handles partial failure and retries once.
- [ ] Test the dispatcher with a table of success and failure cases.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What is the difference between a value and pointer receiver method set?
2. Why can an interface holding a nil pointer be non-nil?
3. Where should an interface be defined?
4. When is embedding the wrong tool?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Effective Go: interfaces](https://go.dev/doc/effective_go#interfaces)
- [ ] [Go blog: interfaces](https://go.dev/blog/)
- [ ] [Go by Example: interfaces](https://gobyexample.com/interfaces)
- [ ] [The Go specification: method sets](https://go.dev/ref/spec#Method_sets)
