# Capstone: Production Platform

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 40 hours

## The problem

Teams integrating with your product need reliable webhook delivery, replay, and visibility. A naive sender loses events, duplicates deliveries, and hides failures. You will design a multi-tenant delivery platform with explicit reliability guarantees.

## Users and main flows

Tenant developers who register webhook endpoints and inspect delivery attempts; platform operators who monitor queue health and incidents.

## What you will build

A multi-tenant webhook delivery platform with an HTTP API, durable event ingestion, at-least-once delivery, idempotent consumers, retries with backoff, dead-letter handling, replay, observability, and Kubernetes deployment.

## Functional requirements

- [ ] Register tenants and endpoints; authenticate API requests with scoped keys
- [ ] Ingest events idempotently and persist them before acknowledging
- [ ] Deliver events at least once with signed payloads
- [ ] Retry transient failures with bounded exponential backoff and jitter
- [ ] Move exhausted deliveries to a dead-letter queue
- [ ] Expose delivery history and a manual replay endpoint
- [ ] Provide queue depth, age, success rate, and latency metrics

## Non-functional requirements

- **Reliability:** No acknowledged event is lost; duplicate delivery is safe for consumers
- **Latency:** p95 ingestion under 50 ms; p95 first delivery under 5 s
- **Throughput:** Sustain 500 events/s locally with bounded memory
- **Observability:** Every delivery has a trace and a structured log with event and tenant IDs
- **Security:** Signed payloads, scoped API keys, SSRF protection for endpoint URLs
- **Operations:** Graceful shutdown drains workers; DLQ depth is alerted
- **Documentation:** Architecture decision records, runbook, and SLO document

## Suggested architecture

```text
tenant API -> ingestion (validate, idempotency, persist event + outbox)
                                  |
                           dispatcher/relay
                                  |
                         delivery queue (bounded)
                                  |
                     worker pool -> HTTP endpoint (signed)
                                  |
                 attempts store + DLQ + metrics/traces

PostgreSQL: tenants, endpoints, events, deliveries, attempts, idempotency_keys
Kubernetes: API Deployment, worker Deployment, HPA, PDB, probes, ConfigMap/Secret
```

## Milestones

| # | Milestone | Definition of done |
| --- | --- | --- |
| 1 | Walking skeleton | Ingest event, persist, deliver to a local receiver |
| 2 | Reliability core | Idempotency, retries, backoff, attempt history |
| 3 | DLQ and replay | Exhausted deliveries replay safely |
| 4 | Multi-tenancy and security | Scoped keys, signed payloads, SSRF protection |
| 5 | Observability | Metrics, traces, logs, SLOs, alerts |
| 6 | Scale and operations | Load test, HPA, graceful shutdown, rollout under load |

## Deliverables

- [ ] Working repository in `labs/expert/capstone-platform/`
- [ ] README with setup, run, and test instructions
- [ ] Tests: unit + at least one integration or end-to-end test
- [ ] `gofmt`, `go vet`, and `go test -race` pass
- [ ] Design decision log
- [ ] Retrospective

## Start here

1. Read [requirements.md](requirements.md) and write the acceptance criteria.
2. Copy the milestone table into your project plan.
3. Build the walking skeleton before adding features.
4. Commit after every working milestone.

## Files

- [Requirements](requirements.md)
- [Milestones](milestones.md)
- [Rubric](rubric.md)
- [Retrospective](retrospective.md)
