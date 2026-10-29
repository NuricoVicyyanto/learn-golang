# 02 · Installation and Setup

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 3 hours

## Why this module

A reliable toolchain removes friction from every later module. You should know how to install Go, inspect the environment, create a module, run code, format it, and fix the most common setup errors.

## Learning objectives

By the end of this module you can:

- [ ] Install Go on Windows, macOS, Linux, or Android/Termux.
- [ ] Explain `GOROOT`, `GOPATH`, and `GOMODCACHE`.
- [ ] Create and run a Go module.
- [ ] Use the core commands: `go run`, `go build`, `go test`, `go fmt`, `go vet`.
- [ ] Configure an editor with formatting, linting, and debugging.
- [ ] Diagnose common installation problems.

---

## 1. Download Go

Go to <https://go.dev/dl/> and choose the installer for your system:

- **Windows:** `.msi`
- **macOS:** `.pkg` or Homebrew
- **Linux:** `.tar.gz`
- **Android:** Termux package

Verify the version after installing:

```bash
go version
go env
```

---

## 2. Install on Windows

1. Download and run the `.msi` installer.
2. Follow the wizard. It sets `GOROOT` and updates the PATH.
3. Open a new terminal and run:

```bash
go version
```

If the command is not found, open a new terminal or check the PATH entry for `C:\Program Files\Go\bin`.

---

## 3. Install on macOS

With Homebrew (recommended):

```bash
brew install go
```

Or install the `.pkg` from the download page. If needed, add Go to your PATH:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Verify:

```bash
go version
```

---

## 4. Install on Linux

Manual installation, which works across distributions:

```bash
wget https://go.dev/dl/go1.26.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz
```

Add Go to your PATH:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

Package-manager alternatives:

```bash
# Ubuntu / Debian
sudo apt update && sudo apt install golang-go

# Arch
sudo pacman -S go

# Fedora
sudo dnf install golang
```

Distribution packages can lag behind the latest release. For current Go, use the official archive.

---

## 5. Install on Android with Termux

1. Install Termux from F-Droid, not the Play Store.
2. Update packages:

```bash
pkg update && pkg upgrade
pkg install golang
```

3. Verify:

```bash
go version
```

Termux is fine for exercises, but some tooling (Docker, heavy profiling) will not be available.

---

## 6. Environment: GOROOT, GOPATH, and modules

| Variable | Meaning |
| --- | --- |
| `GOROOT` | Where the Go toolchain is installed |
| `GOPATH` | Legacy workspace; also stores downloaded modules and binaries |
| `GOMODCACHE` | Where module dependencies are cached |
| `GOBIN` | Where `go install` puts binaries |
| `GOOS` / `GOARCH` | Target operating system and architecture |

Inspect them:

```bash
go env GOROOT
go env GOPATH
go env GOMODCACHE
go env GOBIN
```

Ensure installed binaries are on your PATH:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

> **Modern Go uses modules.** You no longer need to put projects inside `$GOPATH/src`. That layout is legacy; understand it, but do not build new projects around it.

---

## 7. Create your first module

```bash
mkdir hello
cd hello
go mod init example.com/hello
```

Create `main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run .
```

Build a binary:

```bash
go build -o hello .
./hello
```

On Windows:

```bash
hello.exe
```

`go mod init` creates `go.mod`. Commit it. `go.sum` appears when you add dependencies and must also be committed.

---

## 8. Recommended project workflow

```bash
go mod init example.com/project   # once
go get example.com/some/module    # add a dependency
go mod tidy                       # remove unused, add missing
gofmt -w .                        # format
go vet ./...                      # static checks
go test ./...                     # tests
go build ./...                    # compile everything
```

Useful extras:

```bash
go doc fmt.Println                 # documentation in the terminal
go list ./...                      # list packages
go env -w GOFLAGS=-mod=readonly    # optional: fail if go.mod needs changes
```

---

## 9. Editor setup

Use VS Code with the official Go extension, or GoLand. Enable:

- Format on save (`gofmt` or `goimports`).
- Inline diagnostics (`go vet`, staticcheck optional).
- Test running and debugging (`dlv`).
- `gopls` language server.

Minimum expectation: save a file and have it formatted automatically.

---

## 10. Troubleshooting

| Problem | Fix |
| --- | --- |
| `go: command not found` | Add the Go `bin` directory to PATH and open a new terminal |
| `go: cannot find main module` | Run `go mod init` in the project root, or run from the module directory |
| `missing go.sum entry` | Run `go mod tidy` |
| `permission denied` | Fix file permissions or use `sudo` only for installation, not for normal builds |
| Wrong Go version | Check `go version` and `go env GOROOT`; remove old installations from PATH |
| Proxy/network errors | Check `GOPROXY` with `go env GOPROXY` |
| Editor cannot find packages | Restart the language server; check that the workspace root contains `go.mod` |

---

## Exercises

1. Install Go and record the output of `go version` and `go env`.
2. Create a module named `example.com/hello`, print your name, build it, and run the binary.
3. Add a second package to the module and import it from `main`.
4. Run `gofmt` on a deliberately misformatted file and inspect the diff.
5. Trigger two errors from the troubleshooting table and fix them.

## Self-check

1. What is the difference between `GOROOT` and `GOPATH`?
2. Why is `go.sum` committed?
3. What does `go mod tidy` do?
4. When would you use `go run` instead of `go build`?
5. Where should a new project live now that modules exist?

## Lab starter

Runnable starter code for this module:

[`labs/beginner/02-installation-and-setup`](../../../labs/beginner/02-installation-and-setup/README.md)

```bash
cd labs
go run ./beginner/02-installation-and-setup
go test ./beginner/02-installation-and-setup -v
```

## Module checklist

- [ ] Go is installed and on PATH.
- [ ] I created and ran a module.
- [ ] I built a binary and executed it.
- [ ] My editor formats on save.
- [ ] I completed the lab starter in `labs/beginner/02-installation-and-setup`.
- [ ] I can explain `GOROOT`, `GOPATH`, and `GOMODCACHE`.
- [ ] I updated `PROGRESS.md`.

## References

- [Download Go](https://go.dev/dl/)
- [Getting started](https://go.dev/doc/tutorial/getting-started)
- [Managing dependencies](https://go.dev/doc/modules/managing-dependencies)
- [Go environment variables](https://go.dev/doc/install/source#environment)
- [VS Code Go extension](https://github.com/golang/vscode-go)
