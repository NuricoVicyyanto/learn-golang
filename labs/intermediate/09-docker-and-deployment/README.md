# Lab 09 - Docker and Deployment

> **Level:** Intermediate · **Module guide:** [09 Docker and Deployment](../../../docs/02-intermediate/09-docker-and-deployment/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./intermediate/09-docker-and-deployment

# run one exercise
go run ./intermediate/09-docker-and-deployment dockerfile

# run the package tests
go test ./intermediate/09-docker-and-deployment -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 09.1 | `dockerfile` | [Dockerfile for Go](../../../docs/02-intermediate/09-docker-and-deployment/01-dockerfile.md) |
| 09.2 | `compose` | [Docker Compose](../../../docs/02-intermediate/09-docker-and-deployment/02-compose.md) |
| 09.3 | `ci-cd` | [CI/CD](../../../docs/02-intermediate/09-docker-and-deployment/03-ci-cd.md) |

## Workflow

1. Read the linked topic note in `docs/`.
2. Open `NN_<topic>_lab.go` and implement the TODO.
3. Run the exercise and inspect the output.
4. Open `exercises_test.go`, remove the `t.Skip`, and add assertions.
5. Run `go test` and `go vet`.
6. Mark the topic done in the docs note and `PROGRESS.md`.

## Rules

- Keep each exercise small. One concept per function is enough.
- Print labels so the output is readable.
- Do not delete `errNotImplemented` from another exercise.
- Do not import a dependency unless the topic requires it.
- When the module needs PostgreSQL, Redis, or a broker, start it with Docker Compose and document the command in this README.

## Definition of done

- [ ] Every exercise returns `nil` and prints a useful result.
- [ ] Every test has its `t.Skip` removed and real assertions.
- [ ] `gofmt -l .` is clean.
- [ ] `go vet ./...` passes.
- [ ] `go test ./...` passes.
- [ ] You can explain each experiment without notes.
