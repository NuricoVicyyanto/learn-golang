# 01.1 Goroutines

> **Level:** Intermediate · **Module:** 01 Concurrency Fundamentals · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Goroutines are cheap concurrent functions, but cheap does not mean free. You must control their lifetime or they become leaks.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Creating goroutines and the `go` statement; The scheduler and why goroutines are cheap.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Creating goroutines and the `go` statement

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Creating goroutines and the `go` statement"
```

### 2. The scheduler and why goroutines are cheap

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The scheduler and why goroutines are cheap"
```

### 3. `sync.WaitGroup` and waiting for completion

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`sync.WaitGroup` and waiting for completion"
```

### 4. Goroutine leaks and ownership

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Goroutine leaks and ownership"
```

### 5. Closure capture and loop variables

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Closure capture and loop variables"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Creating goroutines and the `go` statement. Predict the output before running.
2. **Change** - modify one detail in The scheduler and why goroutines are cheap. Explain the difference.
3. **Break** - break Closure capture and loop variables on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Creating goroutines and the `go` statement and The scheduler and why goroutines are cheap in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Starting a goroutine with no way to wait for it. Explain the cause and the fix in your own words.
- **Mistake:** Assuming a goroutine will stop when the caller returns. Explain the cause and the fix in your own words.
- **Mistake:** Capturing a loop variable by reference. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Creating goroutines and the `go` statement and The scheduler and why goroutines are cheap. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Closure capture and loop variables.
3. Document one trade-off and one failure mode of Goroutines.

## Self-check

1. Explain Creating goroutines and the `go` statement without opening this note.
2. Compare The scheduler and why goroutines are cheap with an alternative. When is each appropriate?
3. What happens at the edge case of Closure capture and loop variables, and how do you detect it?
4. Where does Goroutines appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go blog: share memory by communicating](https://go.dev/blog/codelab-share)
- [ ] [The Go memory model](https://go.dev/ref/mem)
- [ ] [Package sync](https://pkg.go.dev/sync)
- [ ] [Go blog: advanced concurrency patterns](https://go.dev/blog/advanced-go-concurrency-patterns)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
