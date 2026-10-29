# Lab 02 - Context and Cancellation

> **Level:** Intermediate · **Module guide:** [02 Context and Cancellation](../../../docs/02-intermediate/02-context/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./intermediate/02-context

# run one exercise
go run ./intermediate/02-context context-basics

# run the package tests
go test ./intermediate/02-context -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 02.1 | `context-basics` | [Context Basics](../../../docs/02-intermediate/02-context/01-context-basics.md) |
| 02.2 | `cancellation-and-deadlines` | [Cancellation and Deadlines](../../../docs/02-intermediate/02-context/02-cancellation-and-deadlines.md) |
| 02.3 | `context-in-http-and-db` | [Context in HTTP and Databases](../../../docs/02-intermediate/02-context/03-context-in-http-and-db.md) |

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
