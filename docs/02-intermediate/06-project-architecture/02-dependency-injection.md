# 06.2 Dependency Injection

> **Level:** Intermediate · **Module:** 06 Project Architecture · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Explicit dependencies make code testable and startup failures obvious. Go needs no framework for this.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Constructor injection; Wiring in `main`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Constructor injection

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Constructor injection"
```

### 2. Wiring in `main`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Wiring in `main`"
```

### 3. Interfaces only at boundaries

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Interfaces only at boundaries"
```

### 4. Avoiding global state

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Avoiding global state"
```

### 5. Functional options for configuration

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Functional options for configuration"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Constructor injection. Predict the output before running.
2. **Change** - modify one detail in Wiring in `main`. Explain the difference.
3. **Break** - break Functional options for configuration on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Constructor injection and Wiring in `main` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using a DI container to hide dependencies. Explain the cause and the fix in your own words.
- **Mistake:** Depending on interfaces you do not own. Explain the cause and the fix in your own words.
- **Mistake:** Creating globals because wiring feels tedious. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Constructor injection and Wiring in `main`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Functional options for configuration.
3. Document one trade-off and one failure mode of Dependency Injection.

## Self-check

1. Explain Constructor injection without opening this note.
2. Compare Wiring in `main` with an alternative. When is each appropriate?
3. What happens at the edge case of Functional options for configuration, and how do you detect it?
4. Where does Dependency Injection appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Organizing a Go module](https://go.dev/doc/modules/layout)
- [ ] [Package log/slog](https://pkg.go.dev/log/slog)
- [ ] [Go blog: structured logging](https://go.dev/blog/slog)
- [ ] [Twelve-Factor App](https://12factor.net/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
