# Capstone: REST API

> **Level:** Intermediate · **Status:** ⬜ Not started · **Estimated time:** 30 hours

## The problem

A small team needs a shared bookmark manager with tags, full-text search, and reliable access from scripts and a future web client. The service must be deployable, observable, and safe under failure.

## Users and main flows

Developers and researchers who save links, organize them with tags, and search them later. Operators who deploy and monitor the service.

## What you will build

A production-style REST API with PostgreSQL persistence, authentication, validation, pagination, search, metrics, traces, Docker packaging, and a full test suite.

## Functional requirements

- [ ] Register and authenticate users with hashed passwords and expiring tokens
- [ ] Create, read, update, delete, and list bookmarks owned by the caller
- [ ] Add and remove tags; filter bookmarks by tag
- [ ] Search bookmarks by title and URL with pagination
- [ ] Return consistent JSON errors with request IDs
- [ ] Expose liveness, readiness, and Prometheus metrics
- [ ] Run migrations and start with validated configuration

## Non-functional requirements

- **Performance:** p95 read latency under 100 ms for 10,000 bookmarks per user
- **Reliability:** Graceful shutdown; no request loses data after a 200 response
- **Security:** Passwords hashed with Argon2id or bcrypt; no secrets in logs
- **Observability:** Every request has a log line, a metric, and a trace
- **Testability:** Core logic covered; one end-to-end test against a real database
- **Portability:** Runs with one Compose command on a clean machine

## Suggested architecture

```text
client -> HTTP API (middleware: request ID, auth, rate limit)
                 |
                 +-> service layer -> repository -> PostgreSQL
                 +-> slog + Prometheus + OpenTelemetry
                 +-> migrations on startup (single instance) / job (multi-instance)
```

## Milestones

| # | Milestone | Definition of done |
| --- | --- | --- |
| 1 | Walking skeleton | Server starts, health endpoint works, DB connects |
| 2 | Auth | Register, login, protected route, password hashing |
| 3 | Bookmark CRUD | Ownership enforced, validation, consistent errors |
| 4 | Tags and search | Tag filters and paginated search work |
| 5 | Operations | Logs, metrics, traces, readiness, graceful shutdown |
| 6 | Packaging | Multi-stage image and CI pipeline pass |

## Deliverables

- [ ] Working repository in `labs/intermediate/capstone-api/`
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
