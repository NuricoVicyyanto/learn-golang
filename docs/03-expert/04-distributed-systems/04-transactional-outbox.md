# 04.4 Transactional Outbox

> **Level:** Expert · **Module:** 04 Distributed Systems · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Writing to a database and publishing an event are two operations that can fail independently. The outbox pattern makes the state change and the event atomic.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The dual-write problem; Outbox table and relay worker.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The dual-write problem

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The dual-write problem"
```

### 2. Outbox table and relay worker

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Outbox table and relay worker"
```

### 3. At-least-once publishing and consumer idempotency

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "At-least-once publishing and consumer idempotency"
```

### 4. Schema evolution and event versioning

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Schema evolution and event versioning"
```

### 5. Cleanup and retention

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Cleanup and retention"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for The dual-write problem. Predict the output before running.
2. **Change** - modify one detail in Outbox table and relay worker. Explain the difference.
3. **Break** - break Cleanup and retention on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use The dual-write problem and Outbox table and relay worker in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Publishing before the transaction commits. Explain the cause and the fix in your own words.
- **Mistake:** Deleting outbox rows before publish is confirmed. Explain the cause and the fix in your own words.
- **Mistake:** Treating the outbox as a general-purpose queue. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates The dual-write problem.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Cleanup and retention.
3. Write a short ADR-style note explaining a trade-off in Transactional Outbox.

## Self-check

1. Explain The dual-write problem without opening this note.
2. Compare Outbox table and relay worker with an alternative. When is each appropriate?
3. What happens at the edge case of Cleanup and retention, and how do you detect it?
4. Where does Transactional Outbox appear in a production system?

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
