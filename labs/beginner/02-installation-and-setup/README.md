# Lab 02 — Installation and Setup

> **Level:** Beginner · **Module guide:** [02 Installation and Setup](../../../docs/01-beginner/02-installation-and-setup/README.md)

## Goal

Prove your toolchain works: build, run, format, test, and inspect the environment.

## TODO

1. Complete `greet()` so it handles an empty name.
2. Run the program with and without an argument.
3. Remove the `t.Skip` in `main_test.go` and add cases.
4. Record the output of `go version` and `go env GOROOT GOPATH` in the lab notes.

```bash
cd labs
go run ./beginner/02-installation-and-setup
go run ./beginner/02-installation-and-setup Alice
gofmt -l ./beginner/02-installation-and-setup
go test ./beginner/02-installation-and-setup -v
```

## Definition of done

- [ ] The program handles zero, one, and multiple arguments.
- [ ] `gofmt -l` prints nothing for this folder.
- [ ] The test suite passes.
- [ ] You can explain `GOROOT`, `GOPATH`, and `GOMODCACHE`.
