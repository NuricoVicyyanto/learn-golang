# 07.3 Internal Packages and Project Layout

> **Level:** Beginner · **Module:** 07 Packages and Modules · **Status:** ⬜ Not started
> **Estimated time:** 60 min

## Why it matters

Layout is a communication tool. The `internal/` rule enforces encapsulation, while `cmd/` separates binaries from reusable code.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `internal/` import rule; `cmd/` for binaries and `internal/` for implementation.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `internal/` import rule

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `internal/` import rule"
```

### 2. `cmd/` for binaries and `internal/` for implementation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`cmd/` for binaries and `internal/` for implementation"
```

### 3. When `pkg/` helps and when it is noise

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When `pkg/` helps and when it is noise"
```

### 4. Package-oriented vs layer-oriented layout

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Package-oriented vs layer-oriented layout"
```

### 5. Documenting a package with `doc.go`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Documenting a package with `doc.go`"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for The `internal/` import rule. Predict the output before running.
2. **Change** - modify one detail in `cmd/` for binaries and `internal/` for implementation. Explain the difference.
3. **Break** - break Documenting a package with `doc.go` on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use The `internal/` import rule and `cmd/` for binaries and `internal/` for implementation in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Creating a `pkg/` directory just because a template did. Explain the cause and the fix in your own words.
- **Mistake:** Putting business logic in `cmd/`. Explain the cause and the fix in your own words.
- **Mistake:** Making packages too small to be meaningful. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses The `internal/` import rule and `cmd/` for binaries and `internal/` for implementation. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Documenting a package with `doc.go`.
3. Explain Internal Packages and Project Layout aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain The `internal/` import rule without opening this note.
2. Compare `cmd/` for binaries and `internal/` for implementation with an alternative. When is each appropriate?
3. What happens at the edge case of Documenting a package with `doc.go`, and how do you detect it?
4. Where does Internal Packages and Project Layout appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go modules reference](https://go.dev/ref/mod)
- [ ] [Organizing a Go module](https://go.dev/doc/modules/layout)
- [ ] [Package names](https://go.dev/blog/package-names)
- [ ] [Go by Example: modules](https://gobyexample.com/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
