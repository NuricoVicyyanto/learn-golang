# 04 · HTTP Clients and Servers

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

HTTP is the default interface between services. You need to build handlers that are safe by default and clients that fail gracefully.

## Prerequisites

- Modules 01–03 completed.

## Learning outcomes

- [ ] Build HTTP servers with the standard library.
- [ ] Write middleware for logging, recovery, and authentication.
- [ ] Design JSON APIs with validation and consistent errors.
- [ ] Build HTTP clients with timeouts, retries, and tests.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 04.1 | [net/http Basics](01-net-http-basics.md) | 90–120 min | ⬜ |
| 04.2 | [Routing and Middleware](02-routing-and-middleware.md) | 90–120 min | ⬜ |
| 04.3 | [JSON APIs](03-json-apis.md) | 90–120 min | ⬜ |
| 04.4 | [HTTP Clients and Resilience](04-http-clients.md) | 90–120 min | ⬜ |

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

[`labs/intermediate/04-http-clients-and-servers`](../../../labs/intermediate/04-http-clients-and-servers/README.md)

```bash
cd labs
go run ./intermediate/04-http-clients-and-servers
go test ./intermediate/04-http-clients-and-servers -v
```

## Mini project — URL Shortener API

Build a URL shortener with create, redirect, stats, and delete endpoints. Include validation, collision handling, request limits, and a client package used by the integration tests.

**Deliverables:**

- [ ] Use only `net/http` and the standard library for the core server.
- [ ] Validate URLs and reject unsafe schemes.
- [ ] Handle short-code collisions and expired links.
- [ ] Write handler tests with `httptest` and a client timeout test.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is the default HTTP client unsafe for production?
2. How does middleware ordering affect behavior?
3. What should a JSON error response contain?
4. How do you drain and close a response body correctly?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package net/http](https://pkg.go.dev/net/http)
- [ ] [Go blog: writing web applications](https://go.dev/doc/articles/wiki/)
- [ ] [Go docs: HTTP servers](https://pkg.go.dev/net/http#Server)
- [ ] [Go blog: context](https://go.dev/blog/context)
