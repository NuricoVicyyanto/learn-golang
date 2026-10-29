# Rubric - Capstone: CLI Application

Score each area 0–3. Target: **at least 15/18 with no zero in correctness**.

| Area | What good looks like | Score |
| --- | --- | --- |
| Correctness | All commands work on happy path and edge cases |
| CLI design | Consistent flags, help text, exit codes, and output |
| Persistence | Atomic writes, schema version, corrupt-file handling |
| Tests | Unit and command-level tests; meaningful cases |
| Code quality | Small packages, clear names, no dead code |
| Documentation | README, help text, and decision notes |

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
