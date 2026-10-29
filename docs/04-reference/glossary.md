# Glossary

Add terms as you meet them. The definitions below are starting points; rewrite them in your own words.

## Language

| Term | Working definition |
| --- | --- |
| **Goroutine** | A lightweight thread managed by the Go runtime. Cheap to create, but still needs an owner and a termination path. |
| **Channel** | A typed conduit for sending values between goroutines. Channels synchronize and transfer ownership. |
| **Interface** | A set of method signatures. A type satisfies it implicitly by implementing the methods. |
| **Method set** | The set of methods callable on a type. Value and pointer receivers produce different method sets. |
| **Zero value** | The default value of a type: `0`, `false`, `""`, `nil` for pointers, slices, maps, channels, functions, and interfaces. |
| **Composite literal** | A literal that constructs a struct, array, slice, or map. |
| **Slice header** | The three-word descriptor: pointer to the backing array, length, and capacity. |
| **Escape analysis** | The compiler's decision about whether a value can live on the stack or must be heap-allocated. |
| **Happens-before** | A guarantee that one memory operation is visible before another. The basis of the Go memory model. |
| **Data race** | Two goroutines access the same memory concurrently, at least one writes, and there is no synchronization. Undefined behavior. |
| **Panic** | A runtime failure that unwinds the stack. Recoverable only with `recover` in a deferred function. |
| **Recover** | A built-in that stops a panic in a deferred call. Not a general error-handling mechanism. |
| **Defer** | Schedules a function call to run when the surrounding function returns, in LIFO order. |
| **Closure** | A function value that captures variables from its surrounding scope. |
| **Type assertion** | Extracts a concrete type from an interface value, optionally with a comma-ok result. |
| **Type switch** | A switch over the dynamic type of an interface value. |
| **Embedding** | Including a type as an anonymous field so its methods and fields are promoted. Composition, not inheritance. |
| **Generics** | Type parameters and constraints that let code work across types while remaining statically typed. |
| **Constraint** | The interface that specifies which types and operations a type parameter supports. |

## Runtime and performance

| Term | Working definition |
| --- | --- |
| **GMP** | The scheduler model: G (goroutine), M (OS thread), P (processor context needed to run Go code). |
| **Work stealing** | An idle P takes runnable goroutines from another P's queue. |
| **GOMAXPROCS** | The maximum number of Ps that can execute Go code simultaneously. |
| **GC** | Garbage collection. Go uses a concurrent tri-color mark-and-sweep collector with write barriers. |
| **GOGC** | Target heap growth percentage between collections. |
| **GOMEMLIMIT** | A soft memory limit that makes the GC work harder as the heap approaches it. |
| **Write barrier** | Runtime code that records pointer writes during concurrent marking. |
| **Escape** | A value is allocated on the heap because its lifetime may outlive the stack frame. |
| **pprof** | Go's profiling tool for CPU, heap, goroutine, mutex, and block profiles. |
| **Execution trace** | A timeline of scheduler, GC, syscall, and goroutine events from `runtime/trace`. |
| **Benchmark** | A test function `BenchmarkXxx` measured by the testing package over increasing `b.N`. |
| **Allocation** | Reserving heap memory. A major source of GC pressure and latency. |

## Backend and distributed systems

| Term | Working definition |
| --- | --- |
| **Context** | A request-scoped value carrying cancellation, deadlines, and a small amount of metadata. |
| **Middleware** | A function that wraps an HTTP handler to add cross-cutting behavior. |
| **Repository** | A boundary that hides persistence details behind a domain-facing API. |
| **Transaction** | A group of database operations that commit or roll back together. |
| **Isolation level** | How much concurrent transaction interference the database allows. |
| **Migration** | A versioned, repeatable schema change. |
| **Idempotency** | Repeating an operation produces the same observable result as doing it once. |
| **At-least-once** | A delivery guarantee where messages may be duplicated but are not lost. |
| **Dead-letter queue** | A queue for messages that cannot be processed after the retry policy is exhausted. |
| **Outbox** | A table that stores events in the same transaction as the state change; a relay publishes them later. |
| **Backpressure** | A mechanism that slows producers when consumers cannot keep up. |
| **Circuit breaker** | A failure-isolation pattern that stops calling a failing dependency for a period. |
| **Bulkhead** | Isolating resources so one failing dependency cannot exhaust everything. |
| **SLO** | Service Level Objective: a target for a measurable service level indicator. |
| **Error budget** | The allowed amount of unreliability implied by an SLO. |
| **Cardinality** | The number of distinct label combinations in a metric. High cardinality is expensive. |
| **Trace** | A collection of spans representing one request across services. |
| **Span** | A timed operation within a trace. |

## Add your own

| Term | My definition | Where I used it |
| --- | --- | --- |
|  |  |  |
|  |  |  |
|  |  |  |
