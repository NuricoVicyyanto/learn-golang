# Lab 06 — Errors and Panics

> **Level:** Beginner · **Module guide:** [06 Errors and Panics](../../../docs/01-beginner/06-errors-and-panics/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/06-errors-and-panics

# run one exercise
go run ./beginner/06-errors-and-panics the-error-interface

# run the package tests
go test ./beginner/06-errors-and-panics -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 06.1 | `the-error-interface` | [The Error Interface](../../../docs/01-beginner/06-errors-and-panics/01-the-error-interface.md) |
| 06.2 | `wrapping-and-sentinel-errors` | [Wrapping and Sentinel Errors](../../../docs/01-beginner/06-errors-and-panics/02-wrapping-and-sentinel-errors.md) |
| 06.3 | `custom-errors` | [Custom Errors](../../../docs/01-beginner/06-errors-and-panics/03-custom-errors.md) |
| 06.4 | `panic-and-recover` | [Panic and Recover](../../../docs/01-beginner/06-errors-and-panics/04-panic-and-recover.md) |

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
