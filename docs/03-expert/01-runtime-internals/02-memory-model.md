# 01.2 The Memory Model

> **Level:** Expert · **Module:** 01 Runtime Internals · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

The memory model defines when one goroutine is guaranteed to see another goroutine's writes. Without it, reasoning about concurrency is guesswork.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Happens-before and synchronization edges; Channel operations as synchronization.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Happens-before and synchronization edges

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Happens-before and synchronization edges"
```

### 2. Channel operations as synchronization

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Channel operations as synchronization"
```

### 3. Mutex and atomic operations

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Mutex and atomic operations"
```

### 4. Publication and safe initialization

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Publication and safe initialization"
```

### 5. Data races and undefined behavior

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Data races and undefined behavior"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Happens-before and synchronization edges. Predict the output before running.
2. **Change** — modify one detail in Channel operations as synchronization. Explain the difference.
3. **Break** — break Data races and undefined behavior on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Happens-before and synchronization edges and Channel operations as synchronization in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Assuming a write becomes visible without synchronization. Explain the cause and the fix in your own words.
- **Mistake:** Using `sync/atomic` values incorrectly. Explain the cause and the fix in your own words.
- **Mistake:** Relying on timing or `sleep` for ordering. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Happens-before and synchronization edges.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Data races and undefined behavior.
3. Write a short ADR-style note explaining a trade-off in The Memory Model.

## Self-check

1. Explain Happens-before and synchronization edges without opening this note.
2. Compare Channel operations as synchronization with an alternative. When is each appropriate?
3. What happens at the edge case of Data races and undefined behavior, and how do you detect it?
4. Where does The Memory Model appear in a production system?

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
