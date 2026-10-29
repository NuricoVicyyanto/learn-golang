# Roadmap

This is the ordered path through the curriculum. It assumes no prior Go knowledge, but it does assume you can already program in some language.

---

## Level map

| Level | Modules | Time at 10 h/week | Time at 20 h/week | Capstone | Exit gate |
| --- | --- | --- | --- | --- | --- |
| Beginner | 11 | 8–10 weeks | 4–5 weeks | Tested CLI application | Write a small idiomatic Go program without copying |
| Intermediate | 10 | 10–14 weeks | 5–7 weeks | Production-style REST API | Design, test, containerize, and operate a service |
| Expert | 9 | 12–16 weeks | 6–8 weeks | Reliable distributed platform | Diagnose runtime/performance issues and make architecture trade-offs |

Total: roughly **30–40 weeks at 10 h/week**, or **15–20 weeks at 20 h/week**.

> Hours are estimates, not deadlines. Understanding is the metric; the schedule is only a constraint.

---

## Level 1 - Beginner

**Goal:** become comfortable with the language and its tooling. You should be able to read most ordinary Go code and write small programs without help.

**Prerequisites:** basic programming in any language (variables, loops, functions).

| # | Module | Key outcome | Est. |
| --- | --- | --- | --- |
| 01 | [Introduction to Go](docs/01-beginner/01-introduction-to-go/README.md) | Explain where Go fits, its trade-offs, and its ecosystem | 3 h |
| 02 | [Installation and setup](docs/01-beginner/02-installation-and-setup/README.md) | Working toolchain, editor, module workflow | 3 h |
| 03 | [Language basics](docs/01-beginner/03-language-basics/README.md) | Variables, types, control flow, functions, scope | 12 h |
| 04 | [Composite types](docs/01-beginner/04-composite-types/README.md) | Arrays, slices, maps, structs, pointers | 12 h |
| 05 | [Methods and interfaces](docs/01-beginner/05-methods-and-interfaces/README.md) | Receiver choice, implicit interfaces, composition | 10 h |
| 06 | [Errors and panics](docs/01-beginner/06-errors-and-panics/README.md) | Error wrapping, sentinel errors, panic boundaries | 8 h |
| 07 | [Packages and modules](docs/01-beginner/07-packages-and-modules/README.md) | Package APIs, `go.mod`, semantic import versioning | 6 h |
| 08 | [Standard library essentials](docs/01-beginner/08-standard-library/README.md) | `strings`, `strconv`, `time`, `sort`, `os`, `io` | 10 h |
| 09 | [Files and JSON](docs/01-beginner/09-files-and-json/README.md) | File I/O, JSON contracts, configuration | 8 h |
| 10 | [Testing basics](docs/01-beginner/10-testing-basics/README.md) | Table-driven tests, coverage, first benchmark | 8 h |
| 11 | [Capstone: CLI application](docs/01-beginner/11-capstone-cli/README.md) | Ship a tested command-line tool | 15 h |

**Exit gate - you can:**

- [ ] Write a 300–500 line Go program split into at least three packages.
- [ ] Choose slices vs maps vs structs deliberately.
- [ ] Define interfaces at the consumer side and explain method sets.
- [ ] Return and wrap errors instead of panicking.
- [ ] Write table-driven tests and run them with the race detector.
- [ ] Explain every line of your capstone without looking at notes.

**Do not advance if:** you still rely on copy-paste, cannot explain `nil` interface behavior, or have never debugged a failing test.

---

## Level 2 - Intermediate

**Goal:** build and operate real services. This is where most professional Go work happens.

**Prerequisites:** complete the beginner capstone and its exit gate.

| # | Module | Key outcome | Est. |
| --- | --- | --- | --- |
| 01 | [Concurrency fundamentals](docs/02-intermediate/01-concurrency-fundamentals/README.md) | Goroutines, channels, `select`, `sync`, race detector | 16 h |
| 02 | [Context and cancellation](docs/02-intermediate/02-context/README.md) | Propagate deadlines and cancellation correctly | 8 h |
| 03 | [Generics](docs/02-intermediate/03-generics/README.md) | Type parameters, constraints, generic helpers | 8 h |
| 04 | [HTTP clients and servers](docs/02-intermediate/04-http-clients-and-servers/README.md) | Handlers, middleware, JSON APIs, resilient clients | 16 h |
| 05 | [Databases](docs/02-intermediate/05-databases/README.md) | `database/sql`, transactions, migrations, repositories | 16 h |
| 06 | [Project architecture](docs/02-intermediate/06-project-architecture/README.md) | Layout, dependency injection, config, structured logging | 10 h |
| 07 | [Advanced testing](docs/02-intermediate/07-testing-advanced/README.md) | Fakes, `httptest`, integration tests, fuzzing | 10 h |
| 08 | [Observability](docs/02-intermediate/08-observability/README.md) | Logs, metrics, traces, health checks | 10 h |
| 09 | [Docker and deployment](docs/02-intermediate/09-docker-and-deployment/README.md) | Multi-stage builds, Compose, CI/CD | 8 h |
| 10 | [Capstone: REST API](docs/02-intermediate/10-capstone-rest-api/README.md) | Ship and operate a complete service | 30 h |

**Exit gate - you can:**

- [ ] Explain when a goroutine leaks and prove the fix with tooling.
- [ ] Cancel work through `context` from HTTP handler to database query.
- [ ] Design an API with validation, consistent errors, and pagination.
- [ ] Write unit, handler, and integration tests for the same feature.
- [ ] Run migrations and explain transaction isolation.
- [ ] Build a small, non-root container image and deploy it locally.
- [ ] Add logs, metrics, traces, and health endpoints to a service.

**Do not advance if:** you add concurrency "to make it faster" without measurement, or your service has no timeouts.

---

## Level 3 - Expert

**Goal:** reason about the runtime, performance, reliability, and security of Go systems at scale.

**Prerequisites:** intermediate capstone, comfort with Linux, HTTP, SQL, and containers.

| # | Module | Key outcome | Est. |
| --- | --- | --- | --- |
| 01 | [Runtime internals](docs/03-expert/01-runtime-internals/README.md) | Scheduler, memory model, GC, escape analysis | 16 h |
| 02 | [Performance engineering](docs/03-expert/02-performance-engineering/README.md) | Benchmarks, pprof, traces, allocation reduction | 16 h |
| 03 | [Advanced concurrency](docs/03-expert/03-advanced-concurrency/README.md) | Pools, pipelines, backpressure, leak hunting | 16 h |
| 04 | [Distributed systems](docs/03-expert/04-distributed-systems/README.md) | Retries, idempotency, queues, transactional outbox | 16 h |
| 05 | [gRPC and event-driven systems](docs/03-expert/05-grpc-and-event-driven/README.md) | Protobuf, streaming, Kafka/NATS | 14 h |
| 06 | [Security](docs/03-expert/06-security/README.md) | AuthN/AuthZ, OWASP, TLS, supply chain | 14 h |
| 07 | [Data at scale](docs/03-expert/07-data-at-scale/README.md) | Caching, query tuning, replication, sharding | 12 h |
| 08 | [Kubernetes and cloud](docs/03-expert/08-kubernetes-and-cloud/README.md) | Probes, graceful shutdown, autoscaling, SLOs | 12 h |
| 09 | [Capstone: production platform](docs/03-expert/09-capstone-platform/README.md) | Design and defend a reliable distributed system | 40 h |

**Exit gate - you can:**

- [ ] Read a scheduler or GC trace and explain a latency spike.
- [ ] Produce a benchmark, profile, and measurable before/after improvement.
- [ ] Design for partial failure: retries, timeouts, idempotency, backpressure.
- [ ] Justify consistency, availability, and cost trade-offs in writing.
- [ ] Threat-model a service and fix the top risks.
- [ ] Run a production readiness review and an incident postmortem.

**Do not advance if:** you cannot measure the problem, or your design depends on the network being reliable.

---

## Weekly rhythm

A sustainable week looks like this:

| Session | Activity | Time |
| --- | --- | --- |
| 1 | Read the module README and topic notes; write questions | 60–90 min |
| 2 | Code along with experiments; break things | 60–90 min |
| 3 | Exercises without looking at notes | 60–90 min |
| 4 | Mini project or capstone slice | 90–120 min |
| 5 | Review: re-explain from memory, update notes, run tests | 45–60 min |

At 10 h/week, aim for one module every one to two weeks. At 20 h/week, one module per week plus capstone time.

## Study cycle for every topic

1. **Preview** - read the objectives and skim the references.
2. **Learn** - read one authoritative source; take sparse notes.
3. **Do** - type every example, then change it.
4. **Recall** - close the notes and explain the concept aloud.
5. **Apply** - solve the practice task and write a test.
6. **Review** - revisit after 1 day, 1 week, and 1 month.

## Gate reviews

Before changing level, schedule a 60-minute review with yourself or a peer:

- Explain the level's core ideas without notes.
- Live-code a small problem.
- Walk through the capstone architecture and decisions.
- Show the test suite and one debugging story.
- List what you still do not understand; schedule it.

## Suggested 12-week beginner sprint

| Week | Modules |
| --- | --- |
| 1 | 01–02 |
| 2–3 | 03 |
| 4–5 | 04 |
| 6 | 05 |
| 7 | 06 |
| 8 | 07–08 |
| 9 | 09 |
| 10 | 10 |
| 11–12 | 11 (capstone) + review |

## Suggested 16-week intermediate sprint

| Week | Modules |
| --- | --- |
| 1–2 | 01 |
| 3 | 02 |
| 4 | 03 |
| 5–6 | 04 |
| 7–8 | 05 |
| 9 | 06 |
| 10 | 07 |
| 11 | 08 |
| 12 | 09 |
| 13–16 | 10 (capstone) + gate review |

## Suggested 18-week expert sprint

| Week | Modules |
| --- | --- |
| 1–2 | 01 |
| 3–4 | 02 |
| 5–6 | 03 |
| 7–8 | 04 |
| 9–10 | 05 |
| 11–12 | 06 |
| 13 | 07 |
| 14 | 08 |
| 15–18 | 09 (capstone) + production review |

## What to do when stuck

1. Reproduce the problem in the smallest possible program.
2. Read the official docs and the source of the function you are calling.
3. Add a test that fails for the right reason.
4. Search the error text with the Go version included.
5. Write the question you would ask; half the time, that solves it.
6. Ask for help with: goal, minimal code, expected result, actual result, what you tried.

## After the roadmap

- Contribute to an open-source Go project.
- Read production code: `net/http`, `database/sql`, `slog`, `errgroup`.
- Learn a focused domain: compilers, databases, networking, or distributed storage.
- Practice system design and incident response.
- Teach the material; teaching exposes gaps faster than another tutorial.
