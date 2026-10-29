# Learning Resources

## Official

- [The Go documentation](https://go.dev/doc/) — start here for anything language-related.
- [A Tour of Go](https://go.dev/tour/) — interactive introduction.
- [Effective Go](https://go.dev/doc/effective_go) — idioms and style.
- [The Go specification](https://go.dev/ref/spec) — the authoritative language reference.
- [Go by Example](https://gobyexample.com/) — runnable examples by topic.
- [The Go Blog](https://go.dev/blog/) — language, runtime, and standard library deep dives.
- [Go standard library reference](https://pkg.go.dev/std) — package documentation.
- [Go modules reference](https://go.dev/ref/mod) — dependency and versioning rules.
- [Go memory model](https://go.dev/ref/mem) — required reading for concurrency.
- [Go security documentation](https://go.dev/doc/security/) — vulnerability and supply-chain guidance.

## Books

| Book | Best for | Notes |
| --- | --- | --- |
| *The Go Programming Language* — Donovan & Kernighan | Beginner to intermediate | Dense, precise, still excellent |
| *Learning Go* — Jon Bodner | Beginner to intermediate | Modern coverage, practical |
| *Concurrency in Go* — Katherine Cox-Buday | Intermediate | Patterns and reasoning |
| *100 Go Mistakes and How to Avoid Them* — Teiva Harsanyi | Intermediate | Excellent for code review habits |
| *Designing Data-Intensive Applications* — Martin Kleppmann | Expert | Not Go-specific; essential for distributed systems |
| *Site Reliability Engineering* — Google | Expert | Operations, SLOs, incident response |

## Courses and videos

- [Go.dev learn](https://go.dev/learn/) — official tutorials and guides.
- [GopherCon talks](https://www.youtube.com/@GopherAcademy) — conference talks, especially runtime and production stories.
- [Ardan Labs blog and courses](https://www.ardanlabs.com/) — deep dives into internals and design.
- [Exercism Go track](https://exercism.org/tracks/go) — small exercises with mentoring.
- [Boot.dev Go course](https://www.boot.dev/courses/learn-golang) — guided, project-based.

## Blogs and articles

- [Dave Cheney](https://dave.cheney.net/) — practical Go, performance, API design.
- [Eli Bendersky](https://eli.thegreenplace.net/) — compilers, runtime, Go internals.
- [Bitfield Consulting](https://bitfieldconsulting.com/) — testing and language deep dives.
- [The Go Blog: execution traces](https://go.dev/blog/execution-traces-2024) — scheduler and trace analysis.
- [Go wiki](https://go.dev/wiki/) — community-maintained knowledge.

## Tools

| Tool | Purpose |
| --- | --- |
| `gofmt` / `goimports` | Formatting and import management |
| `go vet` | Suspicious constructs |
| `staticcheck` | Advanced static analysis |
| `golangci-lint` | Meta-linter for CI |
| `delve` | Debugger |
| `air` | Live reload for development |
| `govulncheck` | Known vulnerability scanning |
| `benchstat` | Benchmark comparison |
| `pprof` | CPU, heap, goroutine profiling |
| `go tool trace` | Execution tracing |
| `testcontainers-go` | Disposable integration-test infrastructure |
| `sqlc` | Generate type-safe code from SQL |
| `buf` | Protobuf linting, generation, and compatibility |
| `goleak` | Goroutine leak detection in tests |
| `mockery` | Mock generation (use sparingly) |

## Communities

- [Gophers Slack](https://gophers.slack.com/) — `#general`, `#getting-help`, `#performance`.
- [r/golang](https://www.reddit.com/r/golang/) — news and discussion.
- [Go forum](https://forum.golangbridge.org/) — questions and answers.
- [Stack Overflow: Go](https://stackoverflow.com/questions/tagged/go) — search before asking.

## Reading production code

- [`net/http`](https://github.com/golang/go/tree/master/src/net/http) — server and client design.
- [`database/sql`](https://github.com/golang/go/tree/master/src/database/sql) — pooling and interfaces.
- [`log/slog`](https://github.com/golang/go/tree/master/src/log/slog) — structured logging design.
- [`errgroup`](https://github.com/golang/sync/tree/master/errgroup) — small, readable concurrency.
- [Prometheus](https://github.com/prometheus/prometheus) — large Go codebase.
- [Kubernetes](https://github.com/kubernetes/kubernetes) — very large system; read selectively.

## How to choose a resource

1. Start with official docs for the language or package.
2. Use a book for structure and depth.
3. Use a talk or blog for production experience.
4. Use exercises and projects for skill.
5. Avoid collecting resources instead of finishing one.

## Resource log

Track what you actually used:

| Resource | Module | Useful? | Key takeaway |
| --- | --- | --- | --- |
|  |  |  |  |
|  |  |  |  |
|  |  |  |  |
