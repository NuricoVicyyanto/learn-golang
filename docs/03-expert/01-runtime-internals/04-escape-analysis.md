# 01.4 Escape Analysis

> **Level:** Expert · **Module:** 01 Runtime Internals · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Values that escape to the heap create allocation and GC pressure. Escape analysis explains why a seemingly local value is heap-allocated.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Stack vs heap allocation; `go build -gcflags='-m'` output.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Stack vs heap allocation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Stack vs heap allocation"
```

### 2. `go build -gcflags='-m'` output

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`go build -gcflags='-m'` output"
```

### 3. Pointers, interfaces, and boxing

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pointers, interfaces, and boxing"
```

### 4. Closures and captured variables

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Closures and captured variables"
```

### 5. Reducing allocations without hurting clarity

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Reducing allocations without hurting clarity"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Stack vs heap allocation. Predict the output before running.
2. **Change** - modify one detail in `go build -gcflags='-m'` output. Explain the difference.
3. **Break** - break Reducing allocations without hurting clarity on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Stack vs heap allocation and `go build -gcflags='-m'` output in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Assuming `new` or `&` always means heap. Explain the cause and the fix in your own words.
- **Mistake:** Returning pointers to large structs without measuring. Explain the cause and the fix in your own words.
- **Mistake:** Optimizing allocations that are not on the hot path. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Stack vs heap allocation.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Reducing allocations without hurting clarity.
3. Write a short ADR-style note explaining a trade-off in Escape Analysis.

## Self-check

1. Explain Stack vs heap allocation without opening this note.
2. Compare `go build -gcflags='-m'` output with an alternative. When is each appropriate?
3. What happens at the edge case of Reducing allocations without hurting clarity, and how do you detect it?
4. Where does Escape Analysis appear in a production system?

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
