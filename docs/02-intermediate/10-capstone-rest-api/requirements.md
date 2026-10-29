# Requirements - Capstone: REST API

> Fill in acceptance criteria before coding. A requirement without a test is a wish.

## Scope

**In scope:**

- User registration and token-based authentication
- Bookmark CRUD with tags
- Search and pagination
- PostgreSQL with migrations
- Structured logs, metrics, traces, health checks
- Docker image and Compose stack
- Unit, handler, and integration tests

**Out of scope:**

- Web or mobile UI
- Social features or sharing between users
- Full-text search at internet scale
- Multi-region deployment

## Functional requirements

### FR-01 - Register and authenticate users with hashed passwords and expiring tokens

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-02 - Create, read, update, delete, and list bookmarks owned by the caller

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-03 - Add and remove tags; filter bookmarks by tag

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-04 - Search bookmarks by title and URL with pagination

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-05 - Return consistent JSON errors with request IDs

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-06 - Expose liveness, readiness, and Prometheus metrics

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-07 - Run migrations and start with validated configuration

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

## Non-functional requirements

| Area | Target | How I will verify it |
| --- | --- | --- |
| Performance | p95 read latency under 100 ms for 10,000 bookmarks per user | _test, benchmark, or manual check_ |
| Reliability | Graceful shutdown; no request loses data after a 200 response | _test, benchmark, or manual check_ |
| Security | Passwords hashed with Argon2id or bcrypt; no secrets in logs | _test, benchmark, or manual check_ |
| Observability | Every request has a log line, a metric, and a trace | _test, benchmark, or manual check_ |
| Testability | Core logic covered; one end-to-end test against a real database | _test, benchmark, or manual check_ |
| Portability | Runs with one Compose command on a clean machine | _test, benchmark, or manual check_ |

## Data model

| Entity | Key fields | Constraints |
| --- | --- | --- |
| users | id, email, password_hash, created_at | unique email |
| bookmarks | id, user_id, url, title, description, created_at, updated_at | FK user_id; index on (user_id, created_at) |
| tags | id, user_id, name | unique (user_id, name) |
| bookmark_tags | bookmark_id, tag_id | composite PK |

## Interface

```text
POST   /v1/auth/register
POST   /v1/auth/login
GET    /v1/bookmarks?limit=&cursor=&q=&tag=
POST   /v1/bookmarks
GET    /v1/bookmarks/{id}
PATCH  /v1/bookmarks/{id}
DELETE /v1/bookmarks/{id}
GET    /healthz
GET    /readyz
GET    /metrics
```

## Assumptions and open questions

1.
2.
3.
