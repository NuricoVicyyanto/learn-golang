# 03.4 errgroup and Semaphores

> **Level:** Expert · **Module:** 03 Advanced Concurrency · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

`errgroup` and semaphores express common concurrency constraints with less custom code and fewer mistakes.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `errgroup.Group` and `WithContext`; `SetLimit` for bounded concurrency.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `errgroup.Group` and `WithContext`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`errgroup.Group` and `WithContext`"
```

### 2. `SetLimit` for bounded concurrency

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`SetLimit` for bounded concurrency"
```

### 3. Weighted semaphores

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Weighted semaphores"
```

### 4. First error vs all errors

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "First error vs all errors"
```

### 5. Cancellation on first failure

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Cancellation on first failure"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `errgroup.Group` and `WithContext`. Predict the output before running.
2. **Change** — modify one detail in `SetLimit` for bounded concurrency. Explain the difference.
3. **Break** — break Cancellation on first failure on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `errgroup.Group` and `WithContext` and `SetLimit` for bounded concurrency in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Launching goroutines after the context is cancelled. Explain the cause and the fix in your own words.
- **Mistake:** Losing errors from other goroutines. Explain the cause and the fix in your own words.
- **Mistake:** Using a semaphore without a release path. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates `errgroup.Group` and `WithContext`.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Cancellation on first failure.
3. Write a short ADR-style note explaining a trade-off in errgroup and Semaphores.

## Self-check

1. Explain `errgroup.Group` and `WithContext` without opening this note.
2. Compare `SetLimit` for bounded concurrency with an alternative. When is each appropriate?
3. What happens at the edge case of Cancellation on first failure, and how do you detect it?
4. Where does errgroup and Semaphores appear in a production system?

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
