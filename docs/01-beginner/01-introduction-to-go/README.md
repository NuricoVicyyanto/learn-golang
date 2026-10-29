# 01 · Introduction to Go

> **Level:** Beginner · **Status:** ⬜ Not started · **Estimated time:** 3 hours

## Why this module

Before learning syntax, you need a mental model of what Go is good at and what it deliberately refuses to do. That model helps you choose Go for the right problems and write code that fits the language instead of fighting it.

## Learning objectives

By the end of this module you can:

- [ ] Explain what Go is and why it was created.
- [ ] List its main strengths, weaknesses, and design trade-offs.
- [ ] Name the main areas of the ecosystem and the tools used there.
- [ ] Decide whether Go is a good fit for a given project.
- [ ] Describe how Go differs from at least one language you already know.

---

## What is Go?

Go (often called Golang) is an open-source programming language developed at Google. It is designed to be easy to learn, efficient, safe, and suitable for large systems.

**Design philosophy:**

- Simplicity over cleverness
- Fast compilation
- Memory safety with garbage collection
- Concurrency built into the language (goroutines and channels)
- Strong tooling as part of the standard workflow
- A small language with a large standard library
- Backend services, APIs, CLIs, and infrastructure software

Go is statically typed and compiled. It produces a single binary by default, which makes deployment simple.

---

## A short history

| Year | Event |
| --- | --- |
| 2007 | Development starts at Google by Robert Griesemer, Rob Pike, and Ken Thompson |
| 2009 | Go is released as open source |
| 2012 | Go 1.0 is released, establishing the compatibility promise |
| 2018 | Go 1.11 introduces modules |
| 2022 | Go 1.18 introduces generics |
| Today | Go continues to evolve with a strong backwards-compatibility commitment |

The Go 1 compatibility promise is a major reason the ecosystem is stable: well-written Go code from years ago usually still compiles.

---

## Why Go was created

The creators were frustrated with existing choices for large-scale software:

- C and C++ compiled slowly and had complex build systems.
- Java was verbose and had heavy runtimes.
- Dynamic languages such as Python were productive but harder to scale safely in large teams.
- Concurrency in many languages was an afterthought.

Go aimed to combine:

- The performance and predictability of a compiled language.
- The simplicity and readability of a small language.
- First-class concurrency.
- Tooling that works the same way for everyone.

The result is a language optimized for **engineering teams and maintainability**, not for expressing every possible abstraction.

---

## Strengths

| Strength | What it means in practice |
| --- | --- |
| Fast compilation | Tight feedback loops; large codebases build quickly |
| Small, readable syntax | New team members can read code sooner |
| Garbage collection | No manual memory management in ordinary code |
| Built-in concurrency | Goroutines and channels are part of the language |
| Static typing | Many bugs are caught at compile time |
| Single binary | Easy deployment; no runtime dependency |
| Excellent standard library | HTTP, JSON, SQL, crypto, testing, and more |
| Strong tooling | `gofmt`, `go test`, `go vet`, `pprof`, race detector |
| Cross-compilation | Build for another OS/architecture from one machine |
| Compatibility promise | Upgrades are usually boring, in a good way |

---

## Weaknesses and trade-offs

| Weakness | Trade-off to understand |
| --- | --- |
| No inheritance | Composition and interfaces replace class hierarchies |
| Verbose error handling | Errors are values; there are no exceptions |
| Generics arrived late | Some generic abstractions are still less expressive than in other languages |
| Limited metaprogramming | No macros; reflection is powerful but costly |
| Not ideal for UI | The ecosystem focuses on backend, CLI, and infrastructure |
| Smaller ML/data ecosystem | Python still dominates those areas |
| Easy to write concurrent code badly | Goroutines are cheap, but correctness is still your job |

These are design choices, not accidents. Learn the idioms instead of importing patterns from other languages.

---

## Ecosystem

| Area | Common tools and libraries |
| --- | --- |
| HTTP APIs | `net/http`, `chi`, `gin`, `echo`, `fiber` |
| Databases | `database/sql`, `pgx`, `sqlc`, `sqlx`, `gorm` |
| Testing | `testing`, `testify`, `go-cmp`, `httptest`, `testcontainers-go` |
| CLI tools | `cobra`, `urfave/cli` |
| Configuration | `envconfig`, `viper`, `koanf` |
| Logging | `log/slog`, `zap`, `zerolog` |
| Observability | OpenTelemetry, Prometheus client |
| RPC | gRPC, Connect |
| Messaging | Kafka clients, NATS, RabbitMQ clients |
| Deployment | Docker, Kubernetes, GoReleaser |

Prefer the standard library until a dependency clearly earns its place.

---

## Real projects built with Go

- **Docker** — container platform
- **Kubernetes** — container orchestration
- **Terraform** — infrastructure as code
- **Hugo** — static site generator
- **Prometheus** — monitoring
- **etcd** — distributed key-value store
- Companies such as Cloudflare, Uber, Twitch, Dropbox, and Netflix use Go in production.

---

## When to use Go

Choose Go when you need:

- Backend services and APIs with predictable performance.
- High-concurrency systems with many simultaneous connections.
- CLI tools and automation.
- Network services, proxies, and infrastructure tooling.
- A small deployable binary with simple operations.
- A language that a large team can read consistently.

## When not to use Go

Consider another language when you need:

- Rich desktop or mobile UI.
- Heavy numerical computing, ML research, or data science.
- Deep metaprogramming or expressive type-level abstractions.
- A large existing ecosystem in another language (for example, Python for ML).
- Frontend browser code (use JavaScript/TypeScript, optionally compiled from Go via WebAssembly for niche cases).

---

## How Go compares to languages you may know

| Concern | Go | Typical alternative |
| --- | --- | --- |
| Inheritance | Composition + interfaces | Class hierarchies |
| Errors | Returned values | Exceptions |
| Concurrency | Goroutines + channels | Threads, async/await, actors |
| Build output | Single binary | Runtime or interpreter required |
| Formatting | One official formatter | Multiple style guides |
| Generics | Supported, intentionally simple | Often more expressive |
| Learning curve | Small language, subtle idioms | Varies |

---

## Exercises

1. Write a one-page summary of Go in your own words. Do not copy this page.
2. Compare Go with the language you know best: concurrency, error handling, deployment, and tooling.
3. Find three open-source Go projects you use or find interesting. Identify what they use Go for.
4. Write a short recommendation: for your next project, would you choose Go? Why or why not?

## Self-check

1. Why does Go compile quickly compared with C++?
2. What does "errors are values" mean?
3. Why is a single binary useful in deployment?
4. What is the cost of Go's simplicity?
5. When would you *not* choose Go?

## Lab starter

Runnable starter code for this module:

[`labs/beginner/01-introduction-to-go`](../../../labs/beginner/01-introduction-to-go/README.md)

```bash
cd labs
go run ./beginner/01-introduction-to-go
go test ./beginner/01-introduction-to-go -v
```

## Module checklist

- [ ] I can explain Go's design philosophy without notes.
- [ ] I can name three strengths and three weaknesses.
- [ ] I can map at least five ecosystem tools to their use cases.
- [ ] I completed the exercises.
- [ ] I completed the lab starter in `labs/beginner/01-introduction-to-go`.
- [ ] I updated `PROGRESS.md`.

## References

- [Go documentation](https://go.dev/doc/)
- [Frequently Asked Questions](https://go.dev/doc/faq)
- [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)
- [The Go Blog](https://go.dev/blog/)
- [Effective Go](https://go.dev/doc/effective_go)
