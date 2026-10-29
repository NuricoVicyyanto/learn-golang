# Lab 05 - Databases

> **Level:** Intermediate · **Module guide:** [05 Databases](../../../docs/02-intermediate/05-databases/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./intermediate/05-databases

# run one exercise
go run ./intermediate/05-databases database-sql

# run the package tests
go test ./intermediate/05-databases -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 05.1 | `database-sql` | [database/sql](../../../docs/02-intermediate/05-databases/01-database-sql.md) |
| 05.2 | `migrations` | [Migrations](../../../docs/02-intermediate/05-databases/02-migrations.md) |
| 05.3 | `transactions` | [Transactions](../../../docs/02-intermediate/05-databases/03-transactions.md) |
| 05.4 | `repository-pattern` | [Repository Pattern](../../../docs/02-intermediate/05-databases/04-repository-pattern.md) |

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
