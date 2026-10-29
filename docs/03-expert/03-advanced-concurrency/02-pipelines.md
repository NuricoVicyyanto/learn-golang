# 03.2 Pipelines

> **Level:** Expert · **Module:** 03 Advanced Concurrency · **Status:** ⬜ Not started
> **Estimated time:** 90 min

## Why it matters

Pipelines compose stages that transform streams. Correct cancellation and channel ownership are essential.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Stage functions and channel ownership; Cancellation propagation.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Stage functions and channel ownership

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Stage functions and channel ownership"
```

### 2. Cancellation propagation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Cancellation propagation"
```

### 3. Ordering and concurrency within a stage

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Ordering and concurrency within a stage"
```

### 4. Error channels and termination

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Error channels and termination"
```

### 5. Testing pipelines for leaks

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Testing pipelines for leaks"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Stage functions and channel ownership. Predict the output before running.
2. **Change** — modify one detail in Cancellation propagation. Explain the difference.
3. **Break** — break Testing pipelines for leaks on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Stage functions and channel ownership and Cancellation propagation in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Leaking upstream stages when a consumer exits. Explain the cause and the fix in your own words.
- **Mistake:** Closing a channel in the wrong stage. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring errors to keep the pipeline moving. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Stage functions and channel ownership.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Testing pipelines for leaks.
3. Write a short ADR-style note explaining a trade-off in Pipelines.

## Self-check

1. Explain Stage functions and channel ownership without opening this note.
2. Compare Cancellation propagation with an alternative. When is each appropriate?
3. What happens at the edge case of Testing pipelines for leaks, and how do you detect it?
4. Where does Pipelines appear in a production system?

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
