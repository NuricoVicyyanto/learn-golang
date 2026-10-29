# Lab 04 — Composite Types

> **Level:** Beginner · **Module guide:** [04 Composite Types](../../../docs/01-beginner/04-composite-types/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/04-composite-types

# run one exercise
go run ./beginner/04-composite-types arrays-and-slices

# run the package tests
go test ./beginner/04-composite-types -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 04.1 | `arrays-and-slices` | [Arrays and Slices](../../../docs/01-beginner/04-composite-types/01-arrays-and-slices.md) |
| 04.2 | `maps` | [Maps](../../../docs/01-beginner/04-composite-types/02-maps.md) |
| 04.3 | `structs` | [Structs](../../../docs/01-beginner/04-composite-types/03-structs.md) |
| 04.4 | `pointers` | [Pointers](../../../docs/01-beginner/04-composite-types/04-pointers.md) |

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
