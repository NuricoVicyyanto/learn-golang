# Capstone Lab — REST API

> **Level:** Intermediate · **Spec:** [Capstone REST API](../../../docs/02-intermediate/10-capstone-rest-api/README.md)

## Goal

Turn this walking skeleton into the bookmark API described in the capstone requirements.

## Layout

```text
capstone-rest-api/
├── cmd/api/             wiring and startup
├── internal/config/     environment configuration
├── internal/httpapi/    routes, middleware, handlers
├── internal/store/      persistence boundary
├── migrations/          PostgreSQL schema
├── Dockerfile
└── docker-compose.yml
```

## Run

```bash
cd labs
go run ./intermediate/capstone-rest-api/cmd/api
curl localhost:8080/healthz
```

The starter is in-memory so it compiles without a database. Replace the store with PostgreSQL when you reach the database module.

## TODO order

1. Configure the server with timeouts and graceful shutdown.
2. Implement `GET /healthz` and `GET /readyz`.
3. Add bookmark CRUD with validation and consistent errors.
4. Add tags, filtering, and pagination.
5. Replace the in-memory store with PostgreSQL and migrations.
6. Add auth, request IDs, logs, metrics, and traces.
7. Write handler and integration tests.
8. Build the container and Compose stack.

## Definition of done

- [ ] The API satisfies every functional requirement.
- [ ] Migrations run on a clean database.
- [ ] `go test -race ./...` passes.
- [ ] The container runs as non-root with health checks.
- [ ] You can explain one failure mode and how the API handles it.
