# 09.3 Configuration

> **Level:** Beginner · **Module:** 09 Files and JSON · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Configuration is where code meets operations. Precedence, validation, and secret handling should be explicit rather than accidental.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Configuration precedence: defaults, file, environment, flags; Environment variables and parsing.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Configuration precedence: defaults, file, environment, flags

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Configuration precedence: defaults, file, environment, flags"
```

### 2. Environment variables and parsing

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Environment variables and parsing"
```

### 3. JSON or TOML configuration files

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "JSON or TOML configuration files"
```

### 4. Validation and fail-fast startup

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Validation and fail-fast startup"
```

### 5. Secrets and what must never be logged

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Secrets and what must never be logged"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Configuration precedence: defaults, file, environment, flags. Predict the output before running.
2. **Change** - modify one detail in Environment variables and parsing. Explain the difference.
3. **Break** - break Secrets and what must never be logged on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Configuration precedence: defaults, file, environment, flags and Environment variables and parsing in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Reading environment variables at package init. Explain the cause and the fix in your own words.
- **Mistake:** Logging the entire config struct. Explain the cause and the fix in your own words.
- **Mistake:** Allowing invalid configuration to start the program. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Configuration precedence: defaults, file, environment, flags and Environment variables and parsing. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Secrets and what must never be logged.
3. Explain Configuration aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Configuration precedence: defaults, file, environment, flags without opening this note.
2. Compare Environment variables and parsing with an alternative. When is each appropriate?
3. What happens at the edge case of Secrets and what must never be logged, and how do you detect it?
4. Where does Configuration appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package os](https://pkg.go.dev/os)
- [ ] [Package encoding/json](https://pkg.go.dev/encoding/json)
- [ ] [Go blog: JSON and Go](https://go.dev/blog/json)
- [ ] [Package bufio](https://pkg.go.dev/bufio)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
