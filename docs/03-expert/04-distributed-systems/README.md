# 04 · Distributed Systems

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 16 hours

## Why this module

Networks fail partially, messages duplicate, and clocks lie. Reliable Go services are designed around those facts rather than pretending they do not exist.

## Prerequisites

- Module 03 completed.
- Experience with an HTTP service and a database.

## Learning outcomes

- [ ] Design idempotent operations and safe retries.
- [ ] Apply timeouts, circuit breakers, and bulkheads.
- [ ] Use message queues with explicit delivery semantics.
- [ ] Implement the transactional outbox pattern.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 04.1 | [Idempotency and Retries](01-idempotency-and-retries.md) | 90–120 min | ⬜ |
| 04.2 | [Timeouts and Circuit Breakers](02-timeouts-and-circuit-breakers.md) | 90 min | ⬜ |
| 04.3 | [Message Queues](03-message-queues.md) | 90–120 min | ⬜ |
| 04.4 | [Transactional Outbox](04-transactional-outbox.md) | 90–120 min | ⬜ |

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

[`labs/expert/04-distributed-systems`](../../../labs/expert/04-distributed-systems/README.md)

```bash
cd labs
go run ./expert/04-distributed-systems
go test ./expert/04-distributed-systems -v
```

## Mini project — Reliable Webhook Sender

Build a service that accepts events and delivers them to subscriber URLs with retries, idempotency, dead-letter handling, and visible delivery status. Assume the receiver may be slow, down, or duplicate deliveries.

**Deliverables:**

- [ ] Store events and delivery attempts durably.
- [ ] Retry with backoff and jitter; cap attempts.
- [ ] Make delivery idempotent for receivers.
- [ ] Expose queue depth, age, failures, and DLQ size.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is exactly-once delivery usually an illusion?
2. How does an idempotency key change retry safety?
3. What problem does the transactional outbox solve?
4. How do you prevent a poison message from blocking a queue?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Microservices.io: transactional outbox](https://microservices.io/patterns/data/transactional-outbox.html)
- [ ] [Google SRE book: handling overload](https://sre.google/sre-book/handling-overload/)
- [ ] [AWS: exponential backoff and jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/)
- [ ] [Go blog: context](https://go.dev/blog/context)
