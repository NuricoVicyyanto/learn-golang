# 07.2 Go Modules and Semantic Versioning

> **Level:** Beginner · **Module:** 07 Packages and Modules · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Modules make dependency versions explicit and reproducible. Semantic versioning communicates compatibility, and major-version paths keep incompatible versions importable side by side.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `go mod init`, `go get`, and `go mod tidy`; `go.mod` and `go.sum`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `go mod init`, `go get`, and `go mod tidy`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`go mod init`, `go get`, and `go mod tidy`"
```

### 2. `go.mod` and `go.sum`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`go.mod` and `go.sum`"
```

### 3. Semantic versioning and compatibility

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Semantic versioning and compatibility"
```

### 4. Major version suffixes in import paths

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Major version suffixes in import paths"
```

### 5. Vendoring and module proxies

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Vendoring and module proxies"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `go mod init`, `go get`, and `go mod tidy`. Predict the output before running.
2. **Change** - modify one detail in `go.mod` and `go.sum`. Explain the difference.
3. **Break** - break Vendoring and module proxies on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `go mod init`, `go get`, and `go mod tidy` and `go.mod` and `go.sum` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Editing `go.sum` by hand. Explain the cause and the fix in your own words.
- **Mistake:** Publishing a breaking change without a major version bump. Explain the cause and the fix in your own words.
- **Mistake:** Committing a `replace` directive that only works locally. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `go mod init`, `go get`, and `go mod tidy` and `go.mod` and `go.sum`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Vendoring and module proxies.
3. Explain Go Modules and Semantic Versioning aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `go mod init`, `go get`, and `go mod tidy` without opening this note.
2. Compare `go.mod` and `go.sum` with an alternative. When is each appropriate?
3. What happens at the edge case of Vendoring and module proxies, and how do you detect it?
4. Where does Go Modules and Semantic Versioning appear in a production system?

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
