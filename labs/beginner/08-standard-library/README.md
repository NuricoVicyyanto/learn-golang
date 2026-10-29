# Lab 08 — Standard Library Essentials

> **Level:** Beginner · **Module guide:** [08 Standard Library Essentials](../../../docs/01-beginner/08-standard-library/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/08-standard-library

# run one exercise
go run ./beginner/08-standard-library strings-and-strconv

# run the package tests
go test ./beginner/08-standard-library -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 08.1 | `strings-and-strconv` | [Strings and strconv](../../../docs/01-beginner/08-standard-library/01-strings-and-strconv.md) |
| 08.2 | `time-and-math` | [Time and Math](../../../docs/01-beginner/08-standard-library/02-time-and-math.md) |
| 08.3 | `sort-and-slices` | [Sort and Slices](../../../docs/01-beginner/08-standard-library/03-sort-and-slices.md) |
| 08.4 | `os-and-io` | [OS and I/O](../../../docs/01-beginner/08-standard-library/04-os-and-io.md) |

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
