# Lab 04 - HTTP Clients and Servers

> **Level:** Intermediate · **Module guide:** [04 HTTP Clients and Servers](../../../docs/02-intermediate/04-http-clients-and-servers/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./intermediate/04-http-clients-and-servers

# run one exercise
go run ./intermediate/04-http-clients-and-servers net-http-basics

# run the package tests
go test ./intermediate/04-http-clients-and-servers -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 04.1 | `net-http-basics` | [net/http Basics](../../../docs/02-intermediate/04-http-clients-and-servers/01-net-http-basics.md) |
| 04.2 | `routing-and-middleware` | [Routing and Middleware](../../../docs/02-intermediate/04-http-clients-and-servers/02-routing-and-middleware.md) |
| 04.3 | `json-apis` | [JSON APIs](../../../docs/02-intermediate/04-http-clients-and-servers/03-json-apis.md) |
| 04.4 | `http-clients` | [HTTP Clients and Resilience](../../../docs/02-intermediate/04-http-clients-and-servers/04-http-clients.md) |

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
