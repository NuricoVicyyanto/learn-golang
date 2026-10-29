# Capstone Lab — CLI Application

> **Level:** Beginner · **Spec:** [Capstone CLI](../../../docs/01-beginner/11-capstone-cli/README.md)

## Goal

Turn this skeleton into the tested task CLI described in the capstone requirements.

## Layout

```text
capstone-cli/
├── main.go
├── internal/app/       business rules and commands
├── internal/store/     JSON persistence
└── README.md
```

## Run

```bash
cd labs
go run ./beginner/capstone-cli add "Learn Go"
go run ./beginner/capstone-cli list
go test ./beginner/capstone-cli/... -v
```

## TODO order

1. Define the `Task` model and store interface in `internal/app`.
2. Implement load/save with atomic writes in `internal/store`.
3. Implement `add`, `list`, `done`, `delete`, and `stats` in `internal/app`.
4. Add `--json` output.
5. Replace every `t.Skip` with real tests.
6. Add help text and non-zero exit codes.

## Definition of done

- [ ] All commands work from the README examples.
- [ ] Data survives restarts and a corrupt file is reported safely.
- [ ] `--json` output is stable and testable.
- [ ] `gofmt`, `go vet`, and `go test` pass.
- [ ] You can explain the package boundaries.
