# Lab 01 - Runtime Internals

> **Level:** Expert · **Module guide:** [01 Runtime Internals](../../../docs/03-expert/01-runtime-internals/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./expert/01-runtime-internals

# run one exercise
go run ./expert/01-runtime-internals scheduler-and-gmp

# run the package tests
go test ./expert/01-runtime-internals -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 01.1 | `scheduler-and-gmp` | [Scheduler and GMP](../../../docs/03-expert/01-runtime-internals/01-scheduler-and-gmp.md) |
| 01.2 | `memory-model` | [The Memory Model](../../../docs/03-expert/01-runtime-internals/02-memory-model.md) |
| 01.3 | `garbage-collector` | [Garbage Collector](../../../docs/03-expert/01-runtime-internals/03-garbage-collector.md) |
| 01.4 | `escape-analysis` | [Escape Analysis](../../../docs/03-expert/01-runtime-internals/04-escape-analysis.md) |

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
