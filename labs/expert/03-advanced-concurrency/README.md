# Lab 03 — Advanced Concurrency

> **Level:** Expert · **Module guide:** [03 Advanced Concurrency](../../../docs/03-expert/03-advanced-concurrency/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./expert/03-advanced-concurrency

# run one exercise
go run ./expert/03-advanced-concurrency worker-pools

# run the package tests
go test ./expert/03-advanced-concurrency -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 03.1 | `worker-pools` | [Worker Pools](../../../docs/03-expert/03-advanced-concurrency/01-worker-pools.md) |
| 03.2 | `pipelines` | [Pipelines](../../../docs/03-expert/03-advanced-concurrency/02-pipelines.md) |
| 03.3 | `backpressure` | [Backpressure and Load Shedding](../../../docs/03-expert/03-advanced-concurrency/03-backpressure.md) |
| 03.4 | `errgroup-and-semaphores` | [errgroup and Semaphores](../../../docs/03-expert/03-advanced-concurrency/04-errgroup-and-semaphores.md) |
| 03.5 | `deadlocks-and-leaks` | [Deadlocks and Leaks](../../../docs/03-expert/03-advanced-concurrency/05-deadlocks-and-leaks.md) |

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
