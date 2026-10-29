# 09.1 File I/O

> **Level:** Beginner · **Module:** 09 Files and JSON · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

File operations fail for many reasons: missing paths, permissions, disk space, and partial writes. Correct Go code treats those failures as normal.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `os.ReadFile` and `os.WriteFile` for small files; `os.Open`, `os.Create`, and file handles.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `os.ReadFile` and `os.WriteFile` for small files

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`os.ReadFile` and `os.WriteFile` for small files"
```

### 2. `os.Open`, `os.Create`, and file handles

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`os.Open`, `os.Create`, and file handles"
```

### 3. `bufio.Scanner` and `bufio.Reader`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`bufio.Scanner` and `bufio.Reader`"
```

### 4. File permissions and path handling

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "File permissions and path handling"
```

### 5. `defer Close`, sync, and atomic writes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`defer Close`, sync, and atomic writes"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `os.ReadFile` and `os.WriteFile` for small files. Predict the output before running.
2. **Change** — modify one detail in `os.Open`, `os.Create`, and file handles. Explain the difference.
3. **Break** — break `defer Close`, sync, and atomic writes on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `os.ReadFile` and `os.WriteFile` for small files and `os.Open`, `os.Create`, and file handles in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Ignoring the error from `Close` when writing. Explain the cause and the fix in your own words.
- **Mistake:** Assuming a rename or write is atomic across filesystems. Explain the cause and the fix in your own words.
- **Mistake:** Building paths with string concatenation. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `os.ReadFile` and `os.WriteFile` for small files and `os.Open`, `os.Create`, and file handles. Keep it under 40 lines.
2. Add a table-driven test for the edge case around `defer Close`, sync, and atomic writes.
3. Explain File I/O aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `os.ReadFile` and `os.WriteFile` for small files without opening this note.
2. Compare `os.Open`, `os.Create`, and file handles with an alternative. When is each appropriate?
3. What happens at the edge case of `defer Close`, sync, and atomic writes, and how do you detect it?
4. Where does File I/O appear in a production system?

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
