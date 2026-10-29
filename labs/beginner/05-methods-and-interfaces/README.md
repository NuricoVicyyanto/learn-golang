# Lab 05 - Methods and Interfaces

> **Level:** Beginner · **Module guide:** [05 Methods and Interfaces](../../../docs/01-beginner/05-methods-and-interfaces/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/05-methods-and-interfaces

# run one exercise
go run ./beginner/05-methods-and-interfaces methods-and-receivers

# run the package tests
go test ./beginner/05-methods-and-interfaces -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 05.1 | `methods-and-receivers` | [Methods and Receivers](../../../docs/01-beginner/05-methods-and-interfaces/01-methods-and-receivers.md) |
| 05.2 | `interfaces` | [Interfaces](../../../docs/01-beginner/05-methods-and-interfaces/02-interfaces.md) |
| 05.3 | `embedding-and-composition` | [Embedding and Composition](../../../docs/01-beginner/05-methods-and-interfaces/03-embedding-and-composition.md) |
| 05.4 | `type-assertions-and-switches` | [Type Assertions and Switches](../../../docs/01-beginner/05-methods-and-interfaces/04-type-assertions-and-switches.md) |

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
