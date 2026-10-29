# Lab 04 - Distributed Systems

> **Level:** Expert · **Module guide:** [04 Distributed Systems](../../../docs/03-expert/04-distributed-systems/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./expert/04-distributed-systems

# run one exercise
go run ./expert/04-distributed-systems idempotency-and-retries

# run the package tests
go test ./expert/04-distributed-systems -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 04.1 | `idempotency-and-retries` | [Idempotency and Retries](../../../docs/03-expert/04-distributed-systems/01-idempotency-and-retries.md) |
| 04.2 | `timeouts-and-circuit-breakers` | [Timeouts and Circuit Breakers](../../../docs/03-expert/04-distributed-systems/02-timeouts-and-circuit-breakers.md) |
| 04.3 | `message-queues` | [Message Queues](../../../docs/03-expert/04-distributed-systems/03-message-queues.md) |
| 04.4 | `transactional-outbox` | [Transactional Outbox](../../../docs/03-expert/04-distributed-systems/04-transactional-outbox.md) |

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
