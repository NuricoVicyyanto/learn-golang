# 02.2 Cancellation and Deadlines

> **Level:** Intermediate · **Module:** 02 Context and Cancellation · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Timeouts prevent a slow dependency from taking down the whole service. Cancellation prevents wasted work after the caller has gone.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `context.WithCancel`; `context.WithTimeout` and `WithDeadline`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `context.WithCancel`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`context.WithCancel`"
```

### 2. `context.WithTimeout` and `WithDeadline`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`context.WithTimeout` and `WithDeadline`"
```

### 3. Checking `ctx.Done()` and `ctx.Err()`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Checking `ctx.Done()` and `ctx.Err()`"
```

### 4. Cleanup and `cancel()`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Cleanup and `cancel()`"
```

### 5. Deadline propagation across services

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Deadline propagation across services"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `context.WithCancel`. Predict the output before running.
2. **Change** - modify one detail in `context.WithTimeout` and `WithDeadline`. Explain the difference.
3. **Break** - break Deadline propagation across services on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `context.WithCancel` and `context.WithTimeout` and `WithDeadline` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Forgetting to call `cancel`. Explain the cause and the fix in your own words.
- **Mistake:** Treating deadline exceeded as a programming error. Explain the cause and the fix in your own words.
- **Mistake:** Adding a timeout at every layer without a budget. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `context.WithCancel` and `context.WithTimeout` and `WithDeadline`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Deadline propagation across services.
3. Document one trade-off and one failure mode of Cancellation and Deadlines.

## Self-check

1. Explain `context.WithCancel` without opening this note.
2. Compare `context.WithTimeout` and `WithDeadline` with an alternative. When is each appropriate?
3. What happens at the edge case of Deadline propagation across services, and how do you detect it?
4. Where does Cancellation and Deadlines appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go blog: context](https://go.dev/blog/context)
- [ ] [Package context](https://pkg.go.dev/context)
- [ ] [Go blog: context and cancellation](https://go.dev/blog/context-and-structs)
- [ ] [net/http server timeouts](https://pkg.go.dev/net/http#Server)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
