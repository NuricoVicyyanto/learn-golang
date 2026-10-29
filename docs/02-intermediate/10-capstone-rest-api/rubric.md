# Rubric — Capstone: REST API

Score each area 0–3. Target: **at least 18/24 with no zero in correctness or security**.

| Area | What good looks like | Score |
| --- | --- | --- |
| API design | Resources, status codes, validation, error shape, pagination |
| Correctness | Ownership, auth, edge cases, and data integrity |
| Persistence | Schema, migrations, indexes, transactions |
| Tests | Unit, handler, integration, and failure cases |
| Operations | Logs, metrics, traces, health checks, shutdown |
| Security | Password storage, token handling, input limits |
| Code quality | Packages, dependency direction, readability |
| Documentation | README, API examples, decision log |

## Scoring guide

- **0** — missing or broken.
- **1** — works partially; needs guidance.
- **2** — works; minor gaps in tests or clarity.
- **3** — correct, tested, documented, and defended.

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
