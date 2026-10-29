# Lab 01 — Concurrency Fundamentals

> **Level:** Intermediate · **Module guide:** [01 Concurrency Fundamentals](../../../docs/02-intermediate/01-concurrency-fundamentals/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./intermediate/01-concurrency-fundamentals

# run one exercise
go run ./intermediate/01-concurrency-fundamentals goroutines

# run the package tests
go test ./intermediate/01-concurrency-fundamentals -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 01.1 | `goroutines` | [Goroutines](../../../docs/02-intermediate/01-concurrency-fundamentals/01-goroutines.md) |
| 01.2 | `channels` | [Channels](../../../docs/02-intermediate/01-concurrency-fundamentals/02-channels.md) |
| 01.3 | `select-and-timeouts` | [Select and Timeouts](../../../docs/02-intermediate/01-concurrency-fundamentals/03-select-and-timeouts.md) |
| 01.4 | `sync-primitives` | [Sync Primitives](../../../docs/02-intermediate/01-concurrency-fundamentals/04-sync-primitives.md) |
| 01.5 | `race-detector` | [The Race Detector](../../../docs/02-intermediate/01-concurrency-fundamentals/05-race-detector.md) |
| 01.6 | `concurrency-patterns` | [Concurrency Patterns](../../../docs/02-intermediate/01-concurrency-fundamentals/06-concurrency-patterns.md) |

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
