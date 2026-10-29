# 01.5 The Race Detector

> **Level:** Intermediate · **Module:** 01 Concurrency Fundamentals · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Data races are undefined behavior. The race detector is the cheapest way to find them before production does.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: What a data race is; Running tests with `-race`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. What a data race is

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "What a data race is"
```

### 2. Running tests with `-race`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Running tests with `-race`"
```

### 3. Happens-before and synchronization edges

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Happens-before and synchronization edges"
```

### 4. Common race patterns

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Common race patterns"
```

### 5. Fixing a race without over-locking

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Fixing a race without over-locking"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for What a data race is. Predict the output before running.
2. **Change** - modify one detail in Running tests with `-race`. Explain the difference.
3. **Break** - break Fixing a race without over-locking on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use What a data race is and Running tests with `-race` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Assuming a test suite without `-race` proves safety. Explain the cause and the fix in your own words.
- **Mistake:** Fixing the symptom with `sleep`. Explain the cause and the fix in your own words.
- **Mistake:** Holding a lock while calling untrusted code. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies What a data race is and Running tests with `-race`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Fixing a race without over-locking.
3. Document one trade-off and one failure mode of The Race Detector.

## Self-check

1. Explain What a data race is without opening this note.
2. Compare Running tests with `-race` with an alternative. When is each appropriate?
3. What happens at the edge case of Fixing a race without over-locking, and how do you detect it?
4. Where does The Race Detector appear in a production system?

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
