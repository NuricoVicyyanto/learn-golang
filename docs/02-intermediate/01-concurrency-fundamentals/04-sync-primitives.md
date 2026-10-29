# 01.4 Sync Primitives

> **Level:** Intermediate · **Module:** 01 Concurrency Fundamentals · **Status:** ⬜ Not started
> **Estimated time:** 90 min

## Why it matters

Channels are not always the right tool. Mutexes, atomics, and `sync.Once` are simpler and faster for shared state.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `sync.Mutex` and `sync.RWMutex`; `sync.Once` for one-time initialization.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `sync.Mutex` and `sync.RWMutex`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`sync.Mutex` and `sync.RWMutex`"
```

### 2. `sync.Once` for one-time initialization

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`sync.Once` for one-time initialization"
```

### 3. The `sync/atomic` package

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `sync/atomic` package"
```

### 4. `sync.Cond` and when to avoid it

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`sync.Cond` and when to avoid it"
```

### 5. Choosing channels vs shared memory

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Choosing channels vs shared memory"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `sync.Mutex` and `sync.RWMutex`. Predict the output before running.
2. **Change** - modify one detail in `sync.Once` for one-time initialization. Explain the difference.
3. **Break** - break Choosing channels vs shared memory on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `sync.Mutex` and `sync.RWMutex` and `sync.Once` for one-time initialization in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Copying a mutex after use. Explain the cause and the fix in your own words.
- **Mistake:** Locking and forgetting to unlock on an error path. Explain the cause and the fix in your own words.
- **Mistake:** Reading a variable outside the lock while another goroutine writes it. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `sync.Mutex` and `sync.RWMutex` and `sync.Once` for one-time initialization. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Choosing channels vs shared memory.
3. Document one trade-off and one failure mode of Sync Primitives.

## Self-check

1. Explain `sync.Mutex` and `sync.RWMutex` without opening this note.
2. Compare `sync.Once` for one-time initialization with an alternative. When is each appropriate?
3. What happens at the edge case of Choosing channels vs shared memory, and how do you detect it?
4. Where does Sync Primitives appear in a production system?

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
