# Lab 08 - Kubernetes and Cloud

> **Level:** Expert · **Module guide:** [08 Kubernetes and Cloud](../../../docs/03-expert/08-kubernetes-and-cloud/README.md)

## What this lab is for

One runnable Go package for every topic in the module. Each topic has a starter file with a TODO. Your job is to turn the placeholder into a small, working experiment and then write a real test.

## Run it

From the `labs/` directory:

```bash
# list exercises
go run ./expert/08-kubernetes-and-cloud

# run one exercise
go run ./expert/08-kubernetes-and-cloud containers-and-kubernetes

# run the package tests
go test ./expert/08-kubernetes-and-cloud -v
```

## Exercises

| # | Command | Topic |
| --- | --- | --- |
| 08.1 | `containers-and-kubernetes` | [Containers and Kubernetes](../../../docs/03-expert/08-kubernetes-and-cloud/01-containers-and-kubernetes.md) |
| 08.2 | `graceful-shutdown-and-scaling` | [Graceful Shutdown and Scaling](../../../docs/03-expert/08-kubernetes-and-cloud/02-graceful-shutdown-and-scaling.md) |
| 08.3 | `observability-in-production` | [Observability in Production](../../../docs/03-expert/08-kubernetes-and-cloud/03-observability-in-production.md) |

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
