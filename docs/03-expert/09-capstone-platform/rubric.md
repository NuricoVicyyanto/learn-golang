# Rubric - Capstone: Production Platform

Score each area 0–3. Target: **at least 22/27 with no zero in reliability, security, or testing**.

| Area | What good looks like | Score |
| --- | --- | --- |
| Reliability | No lost events, safe duplicates, retry and DLQ correctness |
| Architecture | Clear boundaries, justified trade-offs, failure isolation |
| Data model | Idempotency, indexes, transactions, retention |
| Concurrency | Bounded workers, backpressure, clean shutdown |
| Security | Tenant isolation, signing, SSRF, secret handling |
| Observability | Useful metrics, traces, logs, SLOs, alerts |
| Operations | Kubernetes manifests, HPA/PDB, runbook |
| Testing | Unit, integration, failure injection, load test |
| Documentation | ADRs, README, runbook, postmortem |

## Scoring guide

- **0** - missing or broken.
- **1** - works partially; needs guidance.
- **2** - works; minor gaps in tests or clarity.
- **3** - correct, tested, documented, and defended.

## Review questions

1. Walk through the architecture and the main request/command flow.
2. Show one failure case and how the system handles it.
3. Show the tests and explain what they do not cover.
4. Explain one trade-off you made and the alternative you rejected.
5. What would break first under 10× load?
6. What would you do differently with another week?

## Sign-off

- **Self-score:**
- **Peer score (optional):**
- **Date:**
- **Decision:** pass / revise
