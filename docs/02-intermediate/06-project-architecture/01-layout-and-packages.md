# 06.1 Layout and Packages

> **Level:** Intermediate · **Module:** 06 Project Architecture · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

A layout is not a template; it is a dependency graph. Good layouts make ownership and change visible.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `cmd/`, `internal/`, and when `pkg/` helps; Feature-oriented vs layer-oriented packages.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `cmd/`, `internal/`, and when `pkg/` helps

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`cmd/`, `internal/`, and when `pkg/` helps"
```

### 2. Feature-oriented vs layer-oriented packages

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Feature-oriented vs layer-oriented packages"
```

### 3. Dependency direction and import cycles

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Dependency direction and import cycles"
```

### 4. Package APIs and exported surface

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Package APIs and exported surface"
```

### 5. Documenting architecture decisions

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Documenting architecture decisions"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `cmd/`, `internal/`, and when `pkg/` helps. Predict the output before running.
2. **Change** - modify one detail in Feature-oriented vs layer-oriented packages. Explain the difference.
3. **Break** - break Documenting architecture decisions on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `cmd/`, `internal/`, and when `pkg/` helps and Feature-oriented vs layer-oriented packages in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Copying a layout template without understanding it. Explain the cause and the fix in your own words.
- **Mistake:** Creating a package per file. Explain the cause and the fix in your own words.
- **Mistake:** Letting `internal` packages import each other in cycles. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `cmd/`, `internal/`, and when `pkg/` helps and Feature-oriented vs layer-oriented packages. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Documenting architecture decisions.
3. Document one trade-off and one failure mode of Layout and Packages.

## Self-check

1. Explain `cmd/`, `internal/`, and when `pkg/` helps without opening this note.
2. Compare Feature-oriented vs layer-oriented packages with an alternative. When is each appropriate?
3. What happens at the edge case of Documenting architecture decisions, and how do you detect it?
4. Where does Layout and Packages appear in a production system?

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
