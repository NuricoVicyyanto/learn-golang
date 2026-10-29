# 07 · Packages and Modules

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 6 hours

## Why this module

Packages are the unit of compilation, naming, and API design; modules are the unit of versioning and distribution. Getting both right keeps a codebase navigable as it grows.

## Prerequisites

- Module 06 completed.
- A working Go module from module 02.

## Learning outcomes

- [ ] Design package boundaries and exported APIs.
- [ ] Explain `go.mod`, `go.sum`, and semantic import versioning.
- [ ] Use `internal/` and standard layout conventions deliberately.
- [ ] Avoid import cycles and unnecessary dependencies.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 07.1 | [Packages and Visibility](01-packages-and-visibility.md) | 60–75 min | ⬜ |
| 07.2 | [Go Modules and Semantic Versioning](02-go-modules-and-semver.md) | 60–75 min | ⬜ |
| 07.3 | [Internal Packages and Project Layout](03-internal-and-layout.md) | 60 min | ⬜ |

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

[`labs/beginner/07-packages-and-modules`](../../../labs/beginner/07-packages-and-modules/README.md)

```bash
cd labs
go run ./beginner/07-packages-and-modules
go test ./beginner/07-packages-and-modules -v
```

## Mini project — Multi-Package Refactor

Refactor an earlier project into at least three packages with a clear dependency direction. Add package documentation, an `internal/` boundary, and a small exported API. The goal is not more files; it is clearer ownership.

**Deliverables:**

- [ ] Draw the dependency graph before refactoring.
- [ ] Move code without changing behavior; keep tests green.
- [ ] Make only the necessary identifiers exported.
- [ ] Write a short ADR explaining the package boundaries.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What makes an identifier exported?
2. Why does Go disallow import cycles?
3. When is `internal/` the right tool?
4. What does a major version suffix in an import path mean?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Go modules reference](https://go.dev/ref/mod)
- [ ] [Organizing a Go module](https://go.dev/doc/modules/layout)
- [ ] [Package names](https://go.dev/blog/package-names)
- [ ] [Go by Example: modules](https://gobyexample.com/)
