# Requirements - Capstone: Production Platform

> Fill in acceptance criteria before coding. A requirement without a test is a wish.

## Scope

**In scope:**

- Tenant and endpoint management with API keys
- Event ingestion with idempotency keys
- Durable delivery queue and worker pool
- Retries with backoff, jitter, and DLQ
- Delivery attempt history and manual replay
- Metrics, traces, structured logs, SLOs
- Kubernetes manifests and load test

**Out of scope:**

- Billing and plan enforcement
- A full web UI
- Multi-region active-active deployment
- Exactly-once delivery claims

## Functional requirements

### FR-01 - Register tenants and endpoints; authenticate API requests with scoped keys

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-02 - Ingest events idempotently and persist them before acknowledging

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-03 - Deliver events at least once with signed payloads

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-04 - Retry transient failures with bounded exponential backoff and jitter

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-05 - Move exhausted deliveries to a dead-letter queue

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-06 - Expose delivery history and a manual replay endpoint

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-07 - Provide queue depth, age, success rate, and latency metrics

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

## Non-functional requirements

| Area | Target | How I will verify it |
| --- | --- | --- |
| Reliability | No acknowledged event is lost; duplicate delivery is safe for consumers | _test, benchmark, or manual check_ |
| Latency | p95 ingestion under 50 ms; p95 first delivery under 5 s | _test, benchmark, or manual check_ |
| Throughput | Sustain 500 events/s locally with bounded memory | _test, benchmark, or manual check_ |
| Observability | Every delivery has a trace and a structured log with event and tenant IDs | _test, benchmark, or manual check_ |
| Security | Signed payloads, scoped API keys, SSRF protection for endpoint URLs | _test, benchmark, or manual check_ |
| Operations | Graceful shutdown drains workers; DLQ depth is alerted | _test, benchmark, or manual check_ |
| Documentation | Architecture decision records, runbook, and SLO document | _test, benchmark, or manual check_ |

## Data model

| Entity | Key fields | Constraints |
| --- | --- | --- |
| tenants | id, name, api_key_hash, created_at | unique name |
| endpoints | id, tenant_id, url, secret, active | FK tenant; URL validated |
| events | id, tenant_id, type, payload, idempotency_key, created_at | unique (tenant_id, idempotency_key) |
| deliveries | id, event_id, endpoint_id, status, attempts, next_attempt_at | index on (status, next_attempt_at) |
| attempts | id, delivery_id, status_code, error, duration_ms, created_at | append-only |

## Interface

```text
POST   /v1/events                 (Idempotency-Key required)
GET    /v1/events/{id}
GET    /v1/events/{id}/deliveries
POST   /v1/deliveries/{id}/replay
GET    /v1/endpoints
POST   /v1/endpoints
DELETE /v1/endpoints/{id}
GET    /healthz
GET    /readyz
GET    /metrics
```

## Assumptions and open questions

1.
2.
3.
