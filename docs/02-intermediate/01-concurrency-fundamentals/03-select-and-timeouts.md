# 01.3 Select and Timeouts

> **Level:** Intermediate · **Module:** 01 Concurrency Fundamentals · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

`select` lets a goroutine wait on multiple events. It is the foundation of cancellation, deadlines, and non-blocking communication.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `select` statement and ready cases; Timeouts with `time.After` and `context`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `select` statement and ready cases

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `select` statement and ready cases"
```

### 2. Timeouts with `time.After` and `context`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Timeouts with `time.After` and `context`"
```

### 3. Non-blocking operations with `default`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Non-blocking operations with `default`"
```

### 4. Tickers and periodic work

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Tickers and periodic work"
```

### 5. Avoiding leaks from timers in loops

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Avoiding leaks from timers in loops"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for The `select` statement and ready cases. Predict the output before running.
2. **Change** — modify one detail in Timeouts with `time.After` and `context`. Explain the difference.
3. **Break** — break Avoiding leaks from timers in loops on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use The `select` statement and ready cases and Timeouts with `time.After` and `context` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using `time.After` in a hot loop and leaking timers. Explain the cause and the fix in your own words.
- **Mistake:** Busy-looping with `default`. Explain the cause and the fix in your own words.
- **Mistake:** Forgetting the cancellation case. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies The `select` statement and ready cases and Timeouts with `time.After` and `context`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Avoiding leaks from timers in loops.
3. Document one trade-off and one failure mode of Select and Timeouts.

## Self-check

1. Explain The `select` statement and ready cases without opening this note.
2. Compare Timeouts with `time.After` and `context` with an alternative. When is each appropriate?
3. What happens at the edge case of Avoiding leaks from timers in loops, and how do you detect it?
4. Where does Select and Timeouts appear in a production system?

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
