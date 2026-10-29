# 05 · gRPC and Event-Driven Systems

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 14 hours

## Why this module

As systems grow, typed RPC and asynchronous events replace ad-hoc HTTP calls. You need schema discipline, streaming, and operational awareness.

## Prerequisites

- Module 04 completed.

## Learning outcomes

- [ ] Define protobuf schemas and generate Go servers and clients.
- [ ] Implement unary and streaming RPCs with deadlines.
- [ ] Choose between request/response and event-driven integration.
- [ ] Operate a Kafka or NATS consumer safely.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 05.1 | [Protobuf and gRPC](01-protobuf-and-grpc.md) | 90–120 min | ⬜ |
| 05.2 | [Streaming RPCs](02-streaming.md) | 75–90 min | ⬜ |
| 05.3 | [Kafka or NATS](03-kafka-or-nats.md) | 90–120 min | ⬜ |

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

[`labs/expert/05-grpc-and-event-driven`](../../../labs/expert/05-grpc-and-event-driven/README.md)

```bash
cd labs
go run ./expert/05-grpc-and-event-driven
go test ./expert/05-grpc-and-event-driven -v
```

## Mini project - Typed Event Service

Build a gRPC service with a companion event consumer. The service owns one aggregate, publishes events through an outbox, and exposes a streaming endpoint. Include schema compatibility checks.

**Deliverables:**

- [ ] Write proto definitions and generate code reproducibly.
- [ ] Implement unary and streaming RPCs with deadlines.
- [ ] Publish and consume one event type idempotently.
- [ ] Document compatibility rules and failure handling.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. Why is field numbering critical in protobuf?
2. When is streaming better than pagination?
3. What does consumer lag tell you?
4. How do you version an event schema without breaking consumers?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [gRPC Go documentation](https://grpc.io/docs/languages/go/)
- [ ] [Protocol Buffers language guide](https://protobuf.dev/programming-guides/proto3/)
- [ ] [buf.build documentation](https://buf.build/docs/)
- [ ] [NATS documentation](https://docs.nats.io/)
- [ ] [Kafka documentation](https://kafka.apache.org/documentation/)
