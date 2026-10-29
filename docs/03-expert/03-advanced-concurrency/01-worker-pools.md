# 03.1 Worker Pools

> **Level:** Expert · **Module:** 03 Advanced Concurrency · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Unbounded concurrency exhausts memory, connections, and downstream capacity. Worker pools make concurrency a deliberate budget.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Fixed workers and a job queue; Result collection and ordering.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Fixed workers and a job queue

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Fixed workers and a job queue"
```

### 2. Result collection and ordering

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Result collection and ordering"
```

### 3. Graceful shutdown and draining

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Graceful shutdown and draining"
```

### 4. Error propagation and partial failure

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Error propagation and partial failure"
```

### 5. Pool sizing and metrics

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pool sizing and metrics"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Fixed workers and a job queue. Predict the output before running.
2. **Change** - modify one detail in Result collection and ordering. Explain the difference.
3. **Break** - break Pool sizing and metrics on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Fixed workers and a job queue and Result collection and ordering in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Forgetting to close the job channel. Explain the cause and the fix in your own words.
- **Mistake:** Blocking workers on a full result channel. Explain the cause and the fix in your own words.
- **Mistake:** Sizing the pool without knowing the bottleneck. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Fixed workers and a job queue.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Pool sizing and metrics.
3. Write a short ADR-style note explaining a trade-off in Worker Pools.

## Self-check

1. Explain Fixed workers and a job queue without opening this note.
2. Compare Result collection and ordering with an alternative. When is each appropriate?
3. What happens at the edge case of Pool sizing and metrics, and how do you detect it?
4. Where does Worker Pools appear in a production system?

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
