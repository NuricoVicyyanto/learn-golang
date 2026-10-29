# 04.2 Timeouts and Circuit Breakers

> **Level:** Expert · **Module:** 04 Distributed Systems · **Status:** ⬜ Not started
> **Estimated time:** 90 min

## Why it matters

One slow dependency can consume all workers. Timeouts, circuit breakers, and bulkheads isolate failure.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Deadline budgets and propagation; Circuit breaker states and thresholds.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Deadline budgets and propagation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Deadline budgets and propagation"
```

### 2. Circuit breaker states and thresholds

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Circuit breaker states and thresholds"
```

### 3. Bulkheads and isolation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Bulkheads and isolation"
```

### 4. Fallbacks and degraded responses

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Fallbacks and degraded responses"
```

### 5. Metrics for failures and rejected calls

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Metrics for failures and rejected calls"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Deadline budgets and propagation. Predict the output before running.
2. **Change** - modify one detail in Circuit breaker states and thresholds. Explain the difference.
3. **Break** - break Metrics for failures and rejected calls on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Deadline budgets and propagation and Circuit breaker states and thresholds in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Setting every timeout to the same value. Explain the cause and the fix in your own words.
- **Mistake:** Opening a circuit without a recovery path. Explain the cause and the fix in your own words.
- **Mistake:** Hiding failures behind a silent fallback. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Deadline budgets and propagation.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Metrics for failures and rejected calls.
3. Write a short ADR-style note explaining a trade-off in Timeouts and Circuit Breakers.

## Self-check

1. Explain Deadline budgets and propagation without opening this note.
2. Compare Circuit breaker states and thresholds with an alternative. When is each appropriate?
3. What happens at the edge case of Metrics for failures and rejected calls, and how do you detect it?
4. Where does Timeouts and Circuit Breakers appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Microservices.io: transactional outbox](https://microservices.io/patterns/data/transactional-outbox.html)
- [ ] [Google SRE book: handling overload](https://sre.google/sre-book/handling-overload/)
- [ ] [AWS: exponential backoff and jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/)
- [ ] [Go blog: context](https://go.dev/blog/context)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
