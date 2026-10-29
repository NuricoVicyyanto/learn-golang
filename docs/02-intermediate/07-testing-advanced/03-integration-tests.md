# 07.3 Integration Tests

> **Level:** Intermediate · **Module:** 07 Advanced Testing · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Integration tests catch wiring and SQL mistakes that unit tests cannot. They must be isolated, repeatable, and fast enough to run often.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Testcontainers or disposable databases; Build tags and test separation.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Testcontainers or disposable databases

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Testcontainers or disposable databases"
```

### 2. Build tags and test separation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Build tags and test separation"
```

### 3. Schema setup and cleanup

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Schema setup and cleanup"
```

### 4. Deterministic fixtures

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Deterministic fixtures"
```

### 5. Running integration tests in CI

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Running integration tests in CI"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Testcontainers or disposable databases. Predict the output before running.
2. **Change** - modify one detail in Build tags and test separation. Explain the difference.
3. **Break** - break Running integration tests in CI on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Testcontainers or disposable databases and Build tags and test separation in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Sharing one database across parallel tests without isolation. Explain the cause and the fix in your own words.
- **Mistake:** Depending on data left by a previous run. Explain the cause and the fix in your own words.
- **Mistake:** Making integration tests so slow nobody runs them. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Testcontainers or disposable databases and Build tags and test separation. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Running integration tests in CI.
3. Document one trade-off and one failure mode of Integration Tests.

## Self-check

1. Explain Testcontainers or disposable databases without opening this note.
2. Compare Build tags and test separation with an alternative. When is each appropriate?
3. What happens at the edge case of Running integration tests in CI, and how do you detect it?
4. Where does Integration Tests appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package httptest](https://pkg.go.dev/net/http/httptest)
- [ ] [Go blog: fuzzing](https://go.dev/blog/fuzz-beta)
- [ ] [Go docs: fuzzing](https://go.dev/doc/security/fuzz/)
- [ ] [testcontainers-go](https://golang.testcontainers.org/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
