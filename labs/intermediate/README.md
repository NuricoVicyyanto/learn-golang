# Intermediate Labs

Suggested folders:

```text
intermediate/
├── 01-worker-pool/
├── 02-context/
├── 03-generics/
├── 04-http/
├── 05-database/
├── 06-architecture/
├── 07-testing/
├── 08-observability/
├── 09-docker/
└── capstone-api/
```

## Working rules

- Services should run with one command from their folder.
- Use `context` for every blocking call.
- Add timeouts before adding retries.
- Tests must pass with `-race`.
- Never commit real credentials or `.env` files.

## Local dependencies

Use Docker Compose for PostgreSQL, Redis, or a message broker. Put the compose file at the root of the project folder:

```text
intermediate/capstone-api/
├── cmd/api/
├── internal/
├── migrations/
├── docker-compose.yml
├── Dockerfile
└── README.md
```

## Before moving to expert

- [ ] Capstone runs with one command.
- [ ] `go test -race ./...` passes.
- [ ] Container image builds and runs as non-root.
- [ ] Health, metrics, and logs are present.
- [ ] You can explain the request lifecycle end to end.
