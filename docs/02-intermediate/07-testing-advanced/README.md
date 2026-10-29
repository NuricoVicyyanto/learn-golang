# 07 · Advanced Testing

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 10 hours

## Why this module

Unit tests are not enough for services. You need fakes, handler tests, real-database integration tests, and fuzzing for parsers.

## Prerequisites

- Modules 04–06 completed.
- Comfort with interfaces and HTTP.

## Learning outcomes

- [ ] Choose fakes, stubs, or mocks deliberately.
- [ ] Test HTTP handlers and middleware with `httptest`.
- [ ] Write reliable integration tests with disposable infrastructure.
- [ ] Use fuzzing to find edge cases in parsers.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 07.1 | [Test Doubles](01-test-doubles.md) | 75–90 min | ⬜ |
| 07.2 | [httptest](02-httptest.md) | 75–90 min | ⬜ |
| 07.3 | [Integration Tests](03-integration-tests.md) | 90–120 min | ⬜ |
| 07.4 | [Fuzzing](04-fuzzing.md) | 60–75 min | ⬜ |

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

[`labs/intermediate/07-testing-advanced`](../../../labs/intermediate/07-testing-advanced/README.md)

```bash
cd labs
go run ./intermediate/07-testing-advanced
go test ./intermediate/07-testing-advanced -v
```

## Mini project - Test Harness for the API

Build a test harness around the URL shortener or bookmark API: service tests with fakes, handler tests with `httptest`, and integration tests against a disposable database. Document the test pyramid you chose.

**Deliverables:**

- [ ] Separate unit, handler, and integration tests with build tags.
- [ ] Create reusable fixtures and a test server helper.
- [ ] Fuzz the input validator.
- [ ] Write a short note on what is intentionally not tested.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. When is a fake better than a mock?
2. How do you test a handler without starting a real server?
3. What makes an integration test reliable?
4. What invariant should your fuzz target check?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package httptest](https://pkg.go.dev/net/http/httptest)
- [ ] [Go blog: fuzzing](https://go.dev/blog/fuzz-beta)
- [ ] [Go docs: fuzzing](https://go.dev/doc/security/fuzz/)
- [ ] [testcontainers-go](https://golang.testcontainers.org/)
