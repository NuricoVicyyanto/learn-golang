# 08 · Observability

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 10 hours

## Why this module

You cannot operate what you cannot see. Logs, metrics, traces, and health checks turn a black box into a debuggable system.

## Prerequisites

- Modules 04 and 06 completed.

## Learning outcomes

- [ ] Emit structured logs with request context.
- [ ] Instrument a service with useful metrics.
- [ ] Add distributed tracing with OpenTelemetry.
- [ ] Design liveness, readiness, and dependency health checks.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 08.1 | [Structured Logging](01-structured-logging.md) | 60–75 min | ⬜ |
| 08.2 | [Metrics](02-metrics.md) | 75–90 min | ⬜ |
| 08.3 | [Tracing](03-tracing.md) | 75–90 min | ⬜ |
| 08.4 | [Health Checks](04-health-checks.md) | 60 min | ⬜ |

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

[`labs/intermediate/08-observability`](../../../labs/intermediate/08-observability/README.md)

```bash
cd labs
go run ./intermediate/08-observability
go test ./intermediate/08-observability -v
```

## Mini project - Observable Service

Take the API from module 04 and make it operable: structured logs, RED metrics, traces across handler and database, and separate liveness/readiness endpoints. Include a small dashboard description.

**Deliverables:**

- [ ] Add correlation IDs to logs and traces.
- [ ] Expose metrics and document each one.
- [ ] Implement dependency-aware readiness.
- [ ] Write three alerting rules with thresholds and rationale.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. When would you use a metric instead of a log?
2. Why are histograms better than averages for latency?
3. What is the difference between liveness and readiness?
4. What makes a trace useful during an incident?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [Package log/slog](https://pkg.go.dev/log/slog)
- [ ] [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)
- [ ] [Prometheus client_golang](https://pkg.go.dev/github.com/prometheus/client_golang)
- [ ] [Google SRE book: monitoring](https://sre.google/sre-book/monitoring-distributed-systems/)
