# 01.1 Scheduler and GMP

> **Level:** Expert · **Module:** 01 Runtime Internals · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Goroutines are multiplexed onto OS threads by the runtime scheduler. Understanding G, M, and P explains latency spikes, syscall blocking, and `GOMAXPROCS` behavior.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Goroutines (G), OS threads (M), and processors (P); Run queues and work stealing.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Goroutines (G), OS threads (M), and processors (P)

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Goroutines (G), OS threads (M), and processors (P)"
```

### 2. Run queues and work stealing

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Run queues and work stealing"
```

### 3. Blocking syscalls and handoff

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Blocking syscalls and handoff"
```

### 4. `GOMAXPROCS` and container CPU limits

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`GOMAXPROCS` and container CPU limits"
```

### 5. Scheduler traces and latency

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Scheduler traces and latency"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Goroutines (G), OS threads (M), and processors (P). Predict the output before running.
2. **Change** - modify one detail in Run queues and work stealing. Explain the difference.
3. **Break** - break Scheduler traces and latency on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Goroutines (G), OS threads (M), and processors (P) and Run queues and work stealing in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Setting `GOMAXPROCS` higher than the container CPU quota. Explain the cause and the fix in your own words.
- **Mistake:** Assuming a goroutine maps to a thread. Explain the cause and the fix in your own words.
- **Mistake:** Diagnosing latency without checking scheduler delays. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Goroutines (G), OS threads (M), and processors (P).
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Scheduler traces and latency.
3. Write a short ADR-style note explaining a trade-off in Scheduler and GMP.

## Self-check

1. Explain Goroutines (G), OS threads (M), and processors (P) without opening this note.
2. Compare Run queues and work stealing with an alternative. When is each appropriate?
3. What happens at the edge case of Scheduler traces and latency, and how do you detect it?
4. Where does Scheduler and GMP appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [The Go memory model](https://go.dev/ref/mem)
- [ ] [Go runtime source](https://github.com/golang/go/tree/master/src/runtime)
- [ ] [Go blog: scheduler tracing](https://go.dev/blog/execution-traces-2024)
- [ ] [Go docs: GC guide](https://go.dev/doc/gc-guide)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
