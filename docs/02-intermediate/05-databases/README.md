# 05 · Databases

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

Most services are mostly a database with an API in front. You need to write safe queries, manage schema changes, and understand transactions.

## Prerequisites

- Module 04 completed.
- Basic SQL knowledge.

## Learning outcomes

- [ ] Use `database/sql` safely with context and prepared statements.
- [ ] Manage schema changes with migrations.
- [ ] Write correct transactions and handle isolation.
- [ ] Structure persistence behind testable repositories.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 05.1 | [database/sql](01-database-sql.md) | 90–120 min | ⬜ |
| 05.2 | [Migrations](02-migrations.md) | 75–90 min | ⬜ |
| 05.3 | [Transactions](03-transactions.md) | 90–120 min | ⬜ |
| 05.4 | [Repository Pattern](04-repository-pattern.md) | 75–90 min | ⬜ |

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

[`labs/intermediate/05-databases`](../../../labs/intermediate/05-databases/README.md)

```bash
cd labs
go run ./intermediate/05-databases
go test ./intermediate/05-databases -v
```

## Mini project — Bookmark Store

Build the persistence layer for a bookmark manager using PostgreSQL and `database/sql`. Include migrations, a repository API, pagination, full-text search, and transaction tests.

**Deliverables:**

- [ ] Design the schema and write up/down migrations.
- [ ] Implement a repository with context-aware queries.
- [ ] Add integration tests against a disposable database.
- [ ] Write a note explaining one index and why it exists.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is `sql.DB` not a single connection?
2. How do you detect no rows without treating it as an error?
3. What is a deadlock, and how should the application respond?
4. When is a repository interface worth the indirection?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package database/sql](https://pkg.go.dev/database/sql)
- [ ] [Go docs: accessing databases](https://go.dev/doc/database/)
- [ ] [Go blog: database/sql](https://go.dev/blog/database-sql)
- [ ] [PostgreSQL documentation](https://www.postgresql.org/docs/)
