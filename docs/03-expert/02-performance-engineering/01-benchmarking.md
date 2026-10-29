# 02.1 Benchmarking

> **Level:** Expert · **Module:** 02 Performance Engineering · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

A bad benchmark produces confident wrong answers. Setup, noise, compiler optimizations, and statistics all matter.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Benchmark structure and `b.N`; `b.ResetTimer`, `b.StopTimer`, and setup.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Benchmark structure and `b.N`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Benchmark structure and `b.N`"
```

### 2. `b.ResetTimer`, `b.StopTimer`, and setup

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`b.ResetTimer`, `b.StopTimer`, and setup"
```

### 3. `benchstat` and run-to-run variance

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`benchstat` and run-to-run variance"
```

### 4. Avoiding dead-code elimination

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Avoiding dead-code elimination"
```

### 5. Benchmarking allocations with `-benchmem`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Benchmarking allocations with `-benchmem`"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Benchmark structure and `b.N`. Predict the output before running.
2. **Change** — modify one detail in `b.ResetTimer`, `b.StopTimer`, and setup. Explain the difference.
3. **Break** — break Benchmarking allocations with `-benchmem` on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Benchmark structure and `b.N` and `b.ResetTimer`, `b.StopTimer`, and setup in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Including setup or I/O in the measured loop. Explain the cause and the fix in your own words.
- **Mistake:** Concluding from a single run. Explain the cause and the fix in your own words.
- **Mistake:** Benchmarking a different workload than production. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Benchmark structure and `b.N`.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Benchmarking allocations with `-benchmem`.
3. Write a short ADR-style note explaining a trade-off in Benchmarking.

## Self-check

1. Explain Benchmark structure and `b.N` without opening this note.
2. Compare `b.ResetTimer`, `b.StopTimer`, and setup with an alternative. When is each appropriate?
3. What happens at the edge case of Benchmarking allocations with `-benchmem`, and how do you detect it?
4. Where does Benchmarking appear in a production system?

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
