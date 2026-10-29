# Go Command Cheatsheet

## Toolchain

```bash
go version                  # installed version
go env                      # environment
go env GOROOT GOPATH        # key paths
go help <command>           # built-in help
```

## Modules and dependencies

```bash
go mod init example.com/app # create go.mod
go get example.com/pkg@v1.2.3
go get -u ./...             # update dependencies (review changes)
go mod tidy                 # add missing, remove unused
go mod download             # download modules
go mod verify               # verify checksums
go list -m all              # module graph
go list -m -u all           # available updates
go mod vendor               # vendor dependencies
```

## Build and run

```bash
go run .                    # compile and run the current package
go run ./cmd/api            # run a specific package
go build -o bin/app .       # build a binary
go build ./...              # build all packages
go install example.com/cmd/tool@latest
go clean -cache             # clear build cache
```

## Format, vet, and static checks

```bash
gofmt -w .                  # format
go fmt ./...                # format packages
go vet ./...                # suspicious constructs
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

## Tests

```bash
go test ./...                       # all packages
go test -v ./...                    # verbose
go test -run TestName ./pkg         # specific test
go test -race ./...                 # race detector
go test -cover ./...                # coverage summary
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out       # coverage report
go test -bench=. -benchmem ./...    # benchmarks
go test -bench=. -benchtime=2s -count=5 ./... | tee bench.txt
benchstat old.txt new.txt           # compare benchmark runs
go test -fuzz=FuzzName ./pkg        # fuzzing
go test -fuzz=FuzzName -fuzztime=30s ./pkg
```

## Profiling and tracing

```bash
go test -cpuprofile=cpu.out -bench=BenchmarkX ./pkg
go tool pprof -http=:8080 cpu.out
go test -memprofile=mem.out -bench=BenchmarkX ./pkg
go tool pprof -http=:8080 mem.out
go test -trace=trace.out ./pkg
go tool trace trace.out
```

## Documentation and analysis

```bash
go doc fmt.Println
go doc -all ./internal/store
go list ./...
go list -deps ./...
go build -gcflags='-m' ./...        # escape analysis
go tool nm ./bin/app | head         # symbols
```

## Cross-compilation

```bash
GOOS=linux   GOARCH=amd64 go build -o app-linux .
GOOS=windows GOARCH=amd64 go build -o app.exe .
GOOS=darwin  GOARCH=arm64 go build -o app-mac .
```

## Common patterns

### Error wrapping

```go
if err != nil {
    return fmt.Errorf("load config: %w", err)
}

if errors.Is(err, os.ErrNotExist) {
    // handle missing file
}
```

### Table-driven test

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 1, 2, 3},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Fatalf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### HTTP server with timeouts

```go
srv := &http.Server{
    Addr:              ":8080",
    Handler:           mux,
    ReadHeaderTimeout: 5 * time.Second,
    ReadTimeout:       10 * time.Second,
    WriteTimeout:      15 * time.Second,
    IdleTimeout:       60 * time.Second,
}
```

### Graceful shutdown

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()

go func() {
    if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}()

<-ctx.Done()
shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
_ = srv.Shutdown(shutdownCtx)
```

### Bounded concurrency

```go
g, ctx := errgroup.WithContext(ctx)
g.SetLimit(8)

for _, item := range items {
    item := item
    g.Go(func() error {
        return process(ctx, item)
    })
}

if err := g.Wait(); err != nil {
    return err
}
```

### Context timeout

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()

if err := doWork(ctx); err != nil {
    return fmt.Errorf("do work: %w", err)
}
```

## Git and repository hygiene

```bash
git status
git add -A
git commit -m "docs(beginner): complete functions note"
git log --oneline --graph --decorate -20
```

Check before committing:

```bash
gofmt -l .
go vet ./...
go test ./...
```
