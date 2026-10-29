# Expert Level

> **Status:** ⬜ Not started · **Estimated time:** 12–16 weeks · 160–220 hours

## Who this level is for

Engineers who already build Go services and now need to reason about runtime behavior, performance, reliability, and production operations.

## Prerequisites

- Intermediate level completed and gate passed.
- Experience with Linux, HTTP, SQL, and containers.
- A service or codebase you can profile and deploy.

## Learning outcomes

- [ ] Explain scheduler, memory, GC, and allocation behavior.
- [ ] Diagnose performance with benchmarks, profiles, and traces.
- [ ] Design bounded, cancellable concurrent systems.
- [ ] Build reliable distributed workflows with explicit failure semantics.
- [ ] Secure, scale, and operate a production service.

## Modules

| # | Module | Est. |
| --- | --- | --- |
| 01 | [Runtime Internals](01-runtime-internals/README.md) | 16 hours |
| 02 | [Performance Engineering](02-performance-engineering/README.md) | 16 hours |
| 03 | [Advanced Concurrency](03-advanced-concurrency/README.md) | 16 hours |
| 04 | [Distributed Systems](04-distributed-systems/README.md) | 16 hours |
| 05 | [gRPC and Event-Driven Systems](05-grpc-and-event-driven/README.md) | 14 hours |
| 06 | [Security](06-security/README.md) | 14 hours |
| 07 | [Data at Scale](07-data-at-scale/README.md) | 12 hours |
| 08 | [Kubernetes and Cloud](08-kubernetes-and-cloud/README.md) | 12 hours |
| 09 | **[Capstone: Production Platform](09-capstone-platform/README.md)** | 40 hours |

## Capstone

Design and operate a multi-tenant webhook delivery platform with idempotent ingestion, at-least-once delivery, retries, DLQ, replay, observability, and Kubernetes deployment. See [09 Capstone: Production Platform](09-capstone-platform/README.md).

## How to work through this level

1. Read this page and the module README before starting.
2. Follow the weekly rhythm in [`ROADMAP.md`](../../ROADMAP.md).
3. Fill every topic note in your own words.
4. Complete the exercises before the mini project.
5. Build the capstone over multiple sessions; do not wait until the end.
6. Run the exit gate before moving to the next level.

## Common pitfalls at this level

- **Optimizing without a baseline.** A profile without a hypothesis is noise.
- **Assuming the network is reliable.** Timeouts, retries, and idempotency are mandatory.
- **Confusing at-least-once with exactly-once.** Design consumers to tolerate duplicates.
- **Unbounded concurrency and queues.** They turn overload into a crash.
- **Treating security as a final checklist.** Threat-model before implementation.
- **Deploying without SLOs.** You cannot decide whether an incident matters without user-impact targets.

## Exit gate

Complete the capstone, pass its rubric with at least 22/27, present an architecture review, demonstrate a failure injection, and write an incident postmortem. Then run the expert gate review in `ROADMAP.md`.

## References

- [ ] [Go documentation](https://go.dev/doc/)
- [ ] [Google SRE book](https://sre.google/sre-book/table-of-contents/)
- [ ] [Designing Data-Intensive Applications](https://dataintensive.net/)
- [ ] [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/)
- [ ] [Kubernetes documentation](https://kubernetes.io/docs/home/)
