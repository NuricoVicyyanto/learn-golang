# Lab 10 — Testing Basics

> **Level:** Beginner · **Module guide:** [10 Testing Basics](../../../docs/01-beginner/10-testing-basics/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/10-testing-basics

# run one exercise
go run ./beginner/10-testing-basics first-test

# run the package tests
go test ./beginner/10-testing-basics -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 10.1 | `first-test` | [Your First Test](../../../docs/01-beginner/10-testing-basics/01-first-test.md) |
| 10.2 | `table-driven-tests` | [Table-Driven Tests](../../../docs/01-beginner/10-testing-basics/02-table-driven-tests.md) |
| 10.3 | `coverage-and-benchmarks` | [Coverage and Benchmarks](../../../docs/01-beginner/10-testing-basics/03-coverage-and-benchmarks.md) |

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
