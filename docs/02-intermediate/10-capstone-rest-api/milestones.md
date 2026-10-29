# Milestones — Capstone: REST API

> One milestone per 2–4 sessions. Do not start the next milestone until the current one is demonstrable.

| # | Milestone | Done when | Evidence | Status |
| --- | --- | --- | --- | --- |
| 1 | Walking skeleton | Server starts, health endpoint works, DB connects | Compose run + test | ⬜ |
| 2 | Auth | Register, login, protected route, password hashing | Handler tests | ⬜ |
| 3 | Bookmark CRUD | Ownership enforced, validation, consistent errors | Integration tests | ⬜ |
| 4 | Tags and search | Tag filters and paginated search work | Repository tests + API tests | ⬜ |
| 5 | Operations | Logs, metrics, traces, readiness, graceful shutdown | Dashboard + shutdown test | ⬜ |
| 6 | Packaging | Multi-stage image and CI pipeline pass | CI run + image scan | ⬜ |

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
