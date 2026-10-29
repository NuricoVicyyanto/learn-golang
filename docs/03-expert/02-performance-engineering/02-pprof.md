# 02.2 pprof

> **Level:** Expert · **Module:** 02 Performance Engineering · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Profiles show where CPU time, memory, and blocking actually go. They turn a vague slowdown into a ranked list of causes.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: CPU, heap, goroutine, mutex, and block profiles; Collecting profiles in tests and services.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. CPU, heap, goroutine, mutex, and block profiles

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "CPU, heap, goroutine, mutex, and block profiles"
```

### 2. Collecting profiles in tests and services

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Collecting profiles in tests and services"
```

### 3. Flame graphs and top lists

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Flame graphs and top lists"
```

### 4. Symbolization and inlining

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Symbolization and inlining"
```

### 5. Production profiling safely

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Production profiling safely"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for CPU, heap, goroutine, mutex, and block profiles. Predict the output before running.
2. **Change** - modify one detail in Collecting profiles in tests and services. Explain the difference.
3. **Break** - break Production profiling safely on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use CPU, heap, goroutine, mutex, and block profiles and Collecting profiles in tests and services in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Profiling for too short a window. Explain the cause and the fix in your own words.
- **Mistake:** Leaving `net/http/pprof` publicly exposed. Explain the cause and the fix in your own words.
- **Mistake:** Optimizing a function that is not on the hot path. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates CPU, heap, goroutine, mutex, and block profiles.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Production profiling safely.
3. Write a short ADR-style note explaining a trade-off in pprof.

## Self-check

1. Explain CPU, heap, goroutine, mutex, and block profiles without opening this note.
2. Compare Collecting profiles in tests and services with an alternative. When is each appropriate?
3. What happens at the edge case of Production profiling safely, and how do you detect it?
4. Where does pprof appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go docs: profiling](https://go.dev/doc/diagnostics)
- [ ] [Package runtime/pprof](https://pkg.go.dev/runtime/pprof)
- [ ] [Package runtime/trace](https://pkg.go.dev/runtime/trace)
- [ ] [Go blog: pprof](https://go.dev/blog/pprof)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
