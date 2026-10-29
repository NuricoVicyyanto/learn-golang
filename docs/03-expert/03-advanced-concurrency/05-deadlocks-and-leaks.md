# 03.5 Deadlocks and Leaks

> **Level:** Expert · **Module:** 03 Advanced Concurrency · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Deadlocks stop progress; leaks silently consume memory. Both are easier to prevent with ownership rules and tooling.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Deadlock conditions and lock ordering; Goroutine leak patterns.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Deadlock conditions and lock ordering

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Deadlock conditions and lock ordering"
```

### 2. Goroutine leak patterns

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Goroutine leak patterns"
```

### 3. Goroutine profiles and stack dumps

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Goroutine profiles and stack dumps"
```

### 4. Leak tests with `goleak`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Leak tests with `goleak`"
```

### 5. Structured concurrency

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Structured concurrency"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Deadlock conditions and lock ordering. Predict the output before running.
2. **Change** - modify one detail in Goroutine leak patterns. Explain the difference.
3. **Break** - break Structured concurrency on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Deadlock conditions and lock ordering and Goroutine leak patterns in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Taking locks in inconsistent order. Explain the cause and the fix in your own words.
- **Mistake:** Waiting on a channel nobody can send to. Explain the cause and the fix in your own words.
- **Mistake:** Assuming `go test` exiting means no goroutines leaked. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Deadlock conditions and lock ordering.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Structured concurrency.
3. Write a short ADR-style note explaining a trade-off in Deadlocks and Leaks.

## Self-check

1. Explain Deadlock conditions and lock ordering without opening this note.
2. Compare Goroutine leak patterns with an alternative. When is each appropriate?
3. What happens at the edge case of Structured concurrency, and how do you detect it?
4. Where does Deadlocks and Leaks appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [ ] [Package semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- [ ] [Go blog: advanced concurrency patterns](https://go.dev/blog/advanced-go-concurrency-patterns)
- [ ] [goleak](https://github.com/uber-go/goleak)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
