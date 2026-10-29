# 07.1 Packages and Visibility

> **Level:** Beginner · **Module:** 07 Packages and Modules · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Exported identifiers form your public API. Package names and boundaries affect readability more than almost any other design choice.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `package` clause and import paths; Exported vs unexported identifiers.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `package` clause and import paths

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `package` clause and import paths"
```

### 2. Exported vs unexported identifiers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Exported vs unexported identifiers"
```

### 3. Package-level initialization and `init`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Package-level initialization and `init`"
```

### 4. Import cycles and how to break them

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Import cycles and how to break them"
```

### 5. Package naming and API surface

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Package naming and API surface"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for The `package` clause and import paths. Predict the output before running.
2. **Change** — modify one detail in Exported vs unexported identifiers. Explain the difference.
3. **Break** — break Package naming and API surface on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use The `package` clause and import paths and Exported vs unexported identifiers in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Exporting everything by default. Explain the cause and the fix in your own words.
- **Mistake:** Using `init` for work that should be explicit. Explain the cause and the fix in your own words.
- **Mistake:** Creating packages by layer instead of by responsibility. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses The `package` clause and import paths and Exported vs unexported identifiers. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Package naming and API surface.
3. Explain Packages and Visibility aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain The `package` clause and import paths without opening this note.
2. Compare Exported vs unexported identifiers with an alternative. When is each appropriate?
3. What happens at the edge case of Package naming and API surface, and how do you detect it?
4. Where does Packages and Visibility appear in a production system?

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
