# Standard Library Map

The Go standard library is large enough to build most services without frameworks. Use this map to find the right package before adding a dependency.

## Core language support

| Package | Use it for | Notes |
| --- | --- | --- |
| `fmt` | Formatting and printing | `%v`, `%+v`, `%q`, `%w` for errors |
| `errors` | Creating, wrapping, inspecting errors | `Is`, `As`, `Join`, `Unwrap` |
| `strings` | String search, split, join, trim | Immutable strings; use `Builder` for construction |
| `strconv` | String ↔ number/bool conversions | Always handle the error |
| `bytes` | Byte slices and buffers | `bytes.Buffer`, `Equal`, `Contains` |
| `unicode` / `unicode/utf8` | Rune classification and UTF-8 | Strings are bytes, not runes |
| `regexp` | Regular expressions | Compile once; avoid in hot paths if possible |
| `sort` | Legacy sorting interfaces | Prefer `slices.Sort` for new code |
| `slices` | Generic slice operations | `Sort`, `Contains`, `Index`, `Clone`, `Delete` |
| `maps` | Generic map helpers | `Keys`, `Values`, `Clone`, `Equal` |
| `cmp` | Comparison helpers | `Compare`, `Or`, `Less` |
| `iter` | Iterators (`iter.Seq`) | Used by `slices` and `maps` helpers |

## I/O and files

| Package | Use it for | Notes |
| --- | --- | --- |
| `io` | Reader/Writer interfaces and helpers | `Copy`, `ReadAll`, `EOF` |
| `bufio` | Buffered reading and writing | `Scanner` has a token size limit |
| `os` | Files, args, env, processes | Prefer `os.ReadFile` for small files |
| `io/fs` | Filesystem abstractions | Useful for embedded files and tests |
| `path/filepath` | OS-aware path manipulation | Never concatenate paths manually |
| `embed` | Embedding static files | `//go:embed` |
| `archive/zip`, `archive/tar` | Archives | Streaming-friendly APIs |

## Encoding and data

| Package | Use it for | Notes |
| --- | --- | --- |
| `encoding/json` | JSON | Tags, `Decoder`, `json.Number` |
| `encoding/csv` | CSV | Configure `Comma`, `FieldsPerRecord` |
| `encoding/xml` | XML | Less common for new APIs |
| `encoding/base64` | Base64 | Useful for keys and payloads |
| `encoding/hex` | Hex encoding | Hashes and IDs |
| `encoding/gob` | Go-specific binary encoding | Not for cross-language APIs |
| `database/sql` | SQL database access | Pair with a driver such as `pgx` |
| `text/template` / `html/template` | Templating | Use `html/template` for untrusted HTML |

## HTTP and networking

| Package | Use it for | Notes |
| --- | --- | --- |
| `net/http` | Servers, clients, middleware | Set timeouts; close bodies |
| `net/http/httptest` | Testing handlers and clients | No real port needed for handlers |
| `net/url` | URL parsing and query values | Validate before fetching |
| `net` | Low-level networking | Rarely needed directly |
| `mime/multipart` | File uploads | Enforce size limits |
| `crypto/tls` | TLS configuration | Do not disable verification |

## Concurrency

| Package | Use it for | Notes |
| --- | --- | --- |
| `context` | Cancellation, deadlines, request values | First parameter; never store in a struct |
| `sync` | Mutex, RWMutex, WaitGroup, Once, Pool, Cond | Do not copy a used mutex |
| `sync/atomic` | Lock-free counters and flags | Use the typed `atomic.Int64` etc. |
| `runtime` | Scheduler and runtime controls | `GOMAXPROCS`, `GC`, `Goexit` |
| `runtime/debug` | Build info, GC tuning, stack dumps | `SetMemoryLimit`, `ReadBuildInfo` |

## Time, math, and randomness

| Package | Use it for | Notes |
| --- | --- | --- |
| `time` | Time, durations, timers | Use `Equal`, not `==`; respect locations |
| `math` | Floating-point math | NaN comparisons are special |
| `math/rand/v2` | Non-cryptographic randomness | Seed is automatic; use `crypto/rand` for secrets |
| `crypto/rand` | Cryptographic randomness | Tokens, keys, salts |

## Logging, testing, and diagnostics

| Package | Use it for | Notes |
| --- | --- | --- |
| `log` | Minimal logging | Default logger is global; prefer `slog` for services |
| `log/slog` | Structured logging | Handlers, levels, attributes |
| `testing` | Tests, benchmarks, fuzzing | `t.Run`, `t.Parallel`, `b.N` |
| `testing/fstest` | In-memory filesystems | Useful for file-based code |
| `runtime/pprof` | CPU and memory profiles | Test and production profiling |
| `runtime/trace` | Execution traces | Scheduler and GC timeline |
| `net/http/pprof` | HTTP profiling endpoints | Never expose publicly without protection |

## CLI and configuration

| Package | Use it for | Notes |
| --- | --- | --- |
| `flag` | Simple command-line flags | Good default for small CLIs |
| `os` | Args and environment | `os.Getenv`, `os.LookupEnv` |
| `encoding/json` | Config files | Validate after decoding |

## Security

| Package | Use it for | Notes |
| --- | --- | --- |
| `crypto/sha256`, `crypto/hmac` | Hashing and signatures | Use `hmac.Equal` for comparison |
| `crypto/subtle` | Constant-time comparison | Tokens and secrets |
| `crypto/aes`, `crypto/cipher` | Symmetric encryption | Prefer high-level libraries for protocols |
| `crypto/x509` | Certificates | TLS and identity |
| `golang.org/x/crypto/argon2` | Password hashing | Argon2id; or `bcrypt` |
| `golang.org/x/vuln/cmd/govulncheck` | Vulnerability scanning | Run in CI |

## When to add a dependency

Add a module when the standard library:

- Does not cover the protocol or format.
- Has an API that would take weeks to build safely (SQL drivers, protobuf, OpenTelemetry).
- Lacks operational features you need (migrations, routing ergonomics).

Do not add a dependency for:

- A few helper functions you can write clearly.
- Logging wrappers around `slog`.
- Simple configuration that `flag` and `os.Getenv` handle.
- A framework when `net/http` is enough.
