# Lab 03 - Language Basics

> **Level:** Beginner · **Module guide:** [03 Language Basics](../../../docs/01-beginner/03-language-basics/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/03-language-basics

# run one exercise
go run ./beginner/03-language-basics variables-constants-types

# run the package tests
go test ./beginner/03-language-basics -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 03.1 | `variables-constants-types` | [Variables, Constants, and Types](../../../docs/01-beginner/03-language-basics/01-variables-constants-types.md) |
| 03.2 | `operators-and-expressions` | [Operators and Expressions](../../../docs/01-beginner/03-language-basics/02-operators-and-expressions.md) |
| 03.3 | `control-flow` | [Control Flow](../../../docs/01-beginner/03-language-basics/03-control-flow.md) |
| 03.4 | `functions` | [Functions](../../../docs/01-beginner/03-language-basics/04-functions.md) |
| 03.5 | `scope-and-shadowing` | [Scope and Shadowing](../../../docs/01-beginner/03-language-basics/05-scope-and-shadowing.md) |

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
