# Lab 09 — Files and JSON

> **Level:** Beginner · **Module guide:** [09 Files and JSON](../../../docs/01-beginner/09-files-and-json/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./beginner/09-files-and-json

# run one exercise
go run ./beginner/09-files-and-json file-io

# run the package tests
go test ./beginner/09-files-and-json -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 09.1 | `file-io` | [File I/O](../../../docs/01-beginner/09-files-and-json/01-file-io.md) |
| 09.2 | `encoding-json` | [Encoding JSON](../../../docs/01-beginner/09-files-and-json/02-encoding-json.md) |
| 09.3 | `configuration` | [Configuration](../../../docs/01-beginner/09-files-and-json/03-configuration.md) |

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
