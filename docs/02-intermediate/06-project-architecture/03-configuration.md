# 06.3 Configuration

> **Level:** Intermediate · **Module:** 06 Project Architecture · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Configuration errors are operational outages. Parsing, defaults, precedence, and validation should fail fast at startup.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Environment variables and flags; Config structs and parsing.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Environment variables and flags

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Environment variables and flags"
```

### 2. Config structs and parsing

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Config structs and parsing"
```

### 3. Defaults and precedence

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Defaults and precedence"
```

### 4. Validation at startup

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Validation at startup"
```

### 5. Secrets and redaction

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Secrets and redaction"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Environment variables and flags. Predict the output before running.
2. **Change** - modify one detail in Config structs and parsing. Explain the difference.
3. **Break** - break Secrets and redaction on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Environment variables and flags and Config structs and parsing in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Parsing config at package init. Explain the cause and the fix in your own words.
- **Mistake:** Accepting zero values that are invalid. Explain the cause and the fix in your own words.
- **Mistake:** Logging secrets during startup. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Environment variables and flags and Config structs and parsing. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Secrets and redaction.
3. Document one trade-off and one failure mode of Configuration.

## Self-check

1. Explain Environment variables and flags without opening this note.
2. Compare Config structs and parsing with an alternative. When is each appropriate?
3. What happens at the edge case of Secrets and redaction, and how do you detect it?
4. Where does Configuration appear in a production system?

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
