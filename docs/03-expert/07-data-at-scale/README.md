# 07 · Data at Scale

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 12 hours

## Why this module

As data and traffic grow, queries, caches, and replication strategy dominate performance. You need to reason about consistency and cost, not just syntax.

## Prerequisites

- Intermediate database module completed.
- Module 02 of this level recommended.

## Learning outcomes

- [ ] Design cache layers with explicit invalidation.
- [ ] Tune queries and indexes using execution plans.
- [ ] Reason about replication lag and sharding trade-offs.
- [ ] Measure the impact of data-layer changes.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 07.1 | [Caching](01-caching.md) | 90–120 min | ⬜ |
| 07.2 | [Query Tuning](02-query-tuning.md) | 90 min | ⬜ |
| 07.3 | [Sharding and Replication](03-sharding-and-replication.md) | 90–120 min | ⬜ |

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

[`labs/expert/07-data-at-scale`](../../../labs/expert/07-data-at-scale/README.md)

```bash
cd labs
go run ./expert/07-data-at-scale
go test ./expert/07-data-at-scale -v
```

## Mini project - Scaled Read Path

Take an API read path and make it scale: add caching with invalidation, optimize the query and index, and document the consistency trade-offs. Include a load test with before/after data.

**Deliverables:**

- [ ] Measure baseline latency and query counts.
- [ ] Add a cache with explicit failure behavior.
- [ ] Tune one query with an execution plan.
- [ ] Write an ADR on consistency and invalidation.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What is a cache stampede, and how do you prevent it?
2. When does an index hurt?
3. Why is offset pagination problematic at scale?
4. What consistency issue does replication lag create?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Use The Index, Luke](https://use-the-index-luke.com/)
- [ ] [Redis documentation](https://redis.io/docs/latest/)
- [ ] [PostgreSQL: indexes](https://www.postgresql.org/docs/current/indexes.html)
- [ ] [Google SRE book: data integrity](https://sre.google/sre-book/data-integrity/)
