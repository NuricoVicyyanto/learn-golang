# Lab 07 — Packages and Modules

> **Level:** Beginner · **Module guide:** [07 Packages and Modules](../../../docs/01-beginner/07-packages-and-modules/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/07-packages-and-modules

# run one exercise
go run ./beginner/07-packages-and-modules packages-and-visibility

# run the package tests
go test ./beginner/07-packages-and-modules -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 07.1 | `packages-and-visibility` | [Packages and Visibility](../../../docs/01-beginner/07-packages-and-modules/01-packages-and-visibility.md) |
| 07.2 | `go-modules-and-semver` | [Go Modules and Semantic Versioning](../../../docs/01-beginner/07-packages-and-modules/02-go-modules-and-semver.md) |
| 07.3 | `internal-and-layout` | [Internal Packages and Project Layout](../../../docs/01-beginner/07-packages-and-modules/03-internal-and-layout.md) |

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
