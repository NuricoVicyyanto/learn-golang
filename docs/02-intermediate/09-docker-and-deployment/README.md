# 09 · Docker and Deployment

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 8 hours

## Why this module

A service is not done when it compiles. You need a reproducible image, a local stack, and an automated path to production.

## Prerequisites

- Modules 04–08 completed.

## Learning outcomes

- [ ] Build small, secure Go container images.
- [ ] Run a local stack with Compose.
- [ ] Create a CI/CD pipeline with tests and releases.
- [ ] Explain health checks and graceful shutdown in containers.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 09.1 | [Dockerfile for Go](01-dockerfile.md) | 75–90 min | ⬜ |
| 09.2 | [Docker Compose](02-compose.md) | 75–90 min | ⬜ |
| 09.3 | [CI/CD](03-ci-cd.md) | 75–90 min | ⬜ |

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

[`labs/intermediate/09-docker-and-deployment`](../../../labs/intermediate/09-docker-and-deployment/README.md)

```bash
cd labs
go run ./intermediate/09-docker-and-deployment
go test ./intermediate/09-docker-and-deployment -v
```

## Mini project - Deployable API

Package the URL shortener or bookmark API for deployment: multi-stage image, Compose stack with database, migrations on startup, health checks, and a CI pipeline that builds and tests everything.

**Deliverables:**

- [ ] Keep the runtime image minimal and non-root.
- [ ] Add graceful shutdown on SIGTERM.
- [ ] Run migrations safely in the deployment path.
- [ ] Document local and production-like run instructions.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why does a multi-stage build produce a smaller image?
2. What does `depends_on` not guarantee?
3. How should migrations run in a multi-replica deployment?
4. What should a CI pipeline verify before publishing an image?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Docker: multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
- [ ] [Docker Compose documentation](https://docs.docker.com/compose/)
- [ ] [Go docs: containerize an app](https://docs.docker.com/language/golang/)
- [ ] [GitHub Actions documentation](https://docs.github.com/actions)
