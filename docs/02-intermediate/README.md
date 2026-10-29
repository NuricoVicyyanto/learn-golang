# Intermediate Level

> **Status:** ⬜ Not started · **Estimated time:** 10–14 weeks · 120–160 hours

## Who this level is for

Developers who know the language and want to build, test, deploy, and operate real services. This level assumes you have shipped at least one small program.

## Prerequisites

- Beginner level completed and gate passed.
- Basic SQL and HTTP knowledge.
- Familiarity with the command line and Git.
- Docker installed (for the deployment module).

## Learning outcomes

- [ ] Design and implement concurrent Go services safely.
- [ ] Build HTTP APIs with validation, middleware, and resilient clients.
- [ ] Persist data with SQL, migrations, and transactions.
- [ ] Structure a codebase for testability and change.
- [ ] Instrument, containerize, and deploy a service.

## Modules

| # | Module | Est. |
| --- | --- | --- |
| 01 | [Concurrency Fundamentals](01-concurrency-fundamentals/README.md) | 16 hours |
| 02 | [Context and Cancellation](02-context/README.md) | 8 hours |
| 03 | [Generics](03-generics/README.md) | 8 hours |
| 04 | [HTTP Clients and Servers](04-http-clients-and-servers/README.md) | 16 hours |
| 05 | [Databases](05-databases/README.md) | 16 hours |
| 06 | [Project Architecture](06-project-architecture/README.md) | 10 hours |
| 07 | [Advanced Testing](07-testing-advanced/README.md) | 10 hours |
| 08 | [Observability](08-observability/README.md) | 10 hours |
| 09 | [Docker and Deployment](09-docker-and-deployment/README.md) | 8 hours |
| 10 | **[Capstone: REST API](10-capstone-rest-api/README.md)** | 30 hours |

## Capstone

Build and operate a production-style bookmark manager API with authentication, PostgreSQL, search, observability, Docker, and a full test suite. See [10 Capstone: REST API](10-capstone-rest-api/README.md).

## How to work through this level

1. Read this page and the module README before starting.
2. Follow the weekly rhythm in [`ROADMAP.md`](../../ROADMAP.md).
3. Fill every topic note in your own words.
4. Complete the exercises before the mini project.
5. Build the capstone over multiple sessions; do not wait until the end.
6. Run the exit gate before moving to the next level.

## Common pitfalls at this level

- **Adding concurrency without measurement.** More goroutines can make things slower and less correct.
- **Skipping context propagation.** A missing context turns a timeout into a stuck request.
- **Treating the database as an implementation detail.** SQL, indexes, and transactions are part of the design.
- **Over-abstracting too early.** Add interfaces when a second implementation or test boundary exists.
- **Testing only the happy path.** Most production bugs live in timeouts, partial failures, and invalid input.
- **Deploying without observability.** If you cannot see it, you cannot operate it.

## Exit gate

Complete the capstone, pass its rubric with at least 18/24, and demonstrate cancellation, migrations, a failing-dependency test, and a container run. Then run the gate review in `ROADMAP.md`.

## References

- [ ] [Effective Go](https://go.dev/doc/effective_go)
- [ ] [Go blog](https://go.dev/blog/)
- [ ] [Package net/http](https://pkg.go.dev/net/http)
- [ ] [Package database/sql](https://pkg.go.dev/database/sql)
- [ ] [Google SRE book](https://sre.google/sre-book/table-of-contents/)
