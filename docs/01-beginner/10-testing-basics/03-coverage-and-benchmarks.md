# 10.3 Coverage and Benchmarks

> **Level:** Beginner · **Module:** 10 Testing Basics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Coverage shows what ran, not what is correct. Benchmarks show cost, not speed in production. Both are useful when interpreted carefully.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `go test -cover` and coverage profiles; Reading a coverage report and finding blind spots.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `go test -cover` and coverage profiles

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`go test -cover` and coverage profiles"
```

### 2. Reading a coverage report and finding blind spots

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Reading a coverage report and finding blind spots"
```

### 3. Benchmark functions and `b.N`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Benchmark functions and `b.N`"
```

### 4. `b.ResetTimer` and avoiding benchmark noise

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`b.ResetTimer` and avoiding benchmark noise"
```

### 5. Interpreting ns/op and allocations

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Interpreting ns/op and allocations"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `go test -cover` and coverage profiles. Predict the output before running.
2. **Change** — modify one detail in Reading a coverage report and finding blind spots. Explain the difference.
3. **Break** — break Interpreting ns/op and allocations on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `go test -cover` and coverage profiles and Reading a coverage report and finding blind spots in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Chasing 100% coverage instead of meaningful cases. Explain the cause and the fix in your own words.
- **Mistake:** Benchmarking with I/O or setup inside the measured loop. Explain the cause and the fix in your own words.
- **Mistake:** Concluding performance from one run. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `go test -cover` and coverage profiles and Reading a coverage report and finding blind spots. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Interpreting ns/op and allocations.
3. Explain Coverage and Benchmarks aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `go test -cover` and coverage profiles without opening this note.
2. Compare Reading a coverage report and finding blind spots with an alternative. When is each appropriate?
3. What happens at the edge case of Interpreting ns/op and allocations, and how do you detect it?
4. Where does Coverage and Benchmarks appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package testing](https://pkg.go.dev/testing)
- [ ] [Go blog: table-driven tests](https://go.dev/blog/subtests)
- [ ] [Go docs: coverage](https://go.dev/doc/build-cover)
- [ ] [Go by Example: testing](https://gobyexample.com/testing)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
