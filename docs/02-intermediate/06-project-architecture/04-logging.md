# 06.4 Structured Logging

> **Level:** Intermediate · **Module:** 06 Project Architecture · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Logs are an interface for operators. Structured fields and request IDs make incidents debuggable.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `log/slog` and handlers; Levels and when to use them.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `log/slog` and handlers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`log/slog` and handlers"
```

### 2. Levels and when to use them

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Levels and when to use them"
```

### 3. Contextual fields and request IDs

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Contextual fields and request IDs"
```

### 4. Redaction and PII

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Redaction and PII"
```

### 5. Logs vs metrics vs traces

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Logs vs metrics vs traces"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `log/slog` and handlers. Predict the output before running.
2. **Change** - modify one detail in Levels and when to use them. Explain the difference.
3. **Break** - break Logs vs metrics vs traces on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `log/slog` and handlers and Levels and when to use them in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Logging the same error at every layer. Explain the cause and the fix in your own words.
- **Mistake:** Using string formatting instead of fields. Explain the cause and the fix in your own words.
- **Mistake:** Logging full request bodies. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `log/slog` and handlers and Levels and when to use them. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Logs vs metrics vs traces.
3. Document one trade-off and one failure mode of Structured Logging.

## Self-check

1. Explain `log/slog` and handlers without opening this note.
2. Compare Levels and when to use them with an alternative. When is each appropriate?
3. What happens at the edge case of Logs vs metrics vs traces, and how do you detect it?
4. Where does Structured Logging appear in a production system?

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
