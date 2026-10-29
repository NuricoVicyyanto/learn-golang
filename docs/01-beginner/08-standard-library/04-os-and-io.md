# 08.4 OS and I/O

> **Level:** Beginner · **Module:** 08 Standard Library Essentials · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

CLI tools and services read and write streams. Understanding `io.Reader`, `io.Writer`, and buffering makes code reusable across files, sockets, and tests.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `os.Args`, environment variables, and exit codes; `io.Reader`, `io.Writer`, and `io.Copy`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `os.Args`, environment variables, and exit codes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`os.Args`, environment variables, and exit codes"
```

### 2. `io.Reader`, `io.Writer`, and `io.Copy`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`io.Reader`, `io.Writer`, and `io.Copy`"
```

### 3. `bufio.Scanner` and buffered I/O

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`bufio.Scanner` and buffered I/O"
```

### 4. Reading stdin and writing stdout

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Reading stdin and writing stdout"
```

### 5. Propagating I/O errors and closing resources

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Propagating I/O errors and closing resources"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `os.Args`, environment variables, and exit codes. Predict the output before running.
2. **Change** — modify one detail in `io.Reader`, `io.Writer`, and `io.Copy`. Explain the difference.
3. **Break** — break Propagating I/O errors and closing resources on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `os.Args`, environment variables, and exit codes and `io.Reader`, `io.Writer`, and `io.Copy` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Ignoring `Scanner.Err()` after the loop. Explain the cause and the fix in your own words.
- **Mistake:** Buffering without flushing. Explain the cause and the fix in your own words.
- **Mistake:** Using `os.Exit` and skipping deferred cleanup. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `os.Args`, environment variables, and exit codes and `io.Reader`, `io.Writer`, and `io.Copy`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Propagating I/O errors and closing resources.
3. Explain OS and I/O aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `os.Args`, environment variables, and exit codes without opening this note.
2. Compare `io.Reader`, `io.Writer`, and `io.Copy` with an alternative. When is each appropriate?
3. What happens at the edge case of Propagating I/O errors and closing resources, and how do you detect it?
4. Where does OS and I/O appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package strings](https://pkg.go.dev/strings)
- [ ] [Package strconv](https://pkg.go.dev/strconv)
- [ ] [Package time](https://pkg.go.dev/time)
- [ ] [Package slices](https://pkg.go.dev/slices)
- [ ] [Package io](https://pkg.go.dev/io)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
