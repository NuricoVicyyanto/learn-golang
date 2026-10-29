# Milestones — Capstone: Production Platform

> One milestone per 2–4 sessions. Do not start the next milestone until the current one is demonstrable.

| # | Milestone | Done when | Evidence | Status |
| --- | --- | --- | --- | --- |
| 1 | Walking skeleton | Ingest event, persist, deliver to a local receiver | End-to-end test | ⬜ |
| 2 | Reliability core | Idempotency, retries, backoff, attempt history | Failure-injection tests | ⬜ |
| 3 | DLQ and replay | Exhausted deliveries replay safely | Integration test | ⬜ |
| 4 | Multi-tenancy and security | Scoped keys, signed payloads, SSRF protection | Security tests | ⬜ |
| 5 | Observability | Metrics, traces, logs, SLOs, alerts | Dashboard + runbook | ⬜ |
| 6 | Scale and operations | Load test, HPA, graceful shutdown, rollout under load | Load report + k8s demo | ⬜ |

## Walking skeleton checklist

- [ ] Repository initialized and runs with one command.
- [ ] Configuration is externalized.
- [ ] One end-to-end path works, even if hard-coded.
- [ ] Logging shows what happened.
- [ ] A test exercises the path.

## Per-milestone review

For each milestone, answer:

1. What works now that did not work before?
2. What did I learn?
3. What did I postpone, and why?
4. What is the next smallest step?
