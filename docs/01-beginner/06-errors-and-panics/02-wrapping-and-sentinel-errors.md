# 06.2 Wrapping and Sentinel Errors

> **Level:** Beginner · **Module:** 06 Errors and Panics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Wrapping adds context without losing the original cause. Sentinel errors let callers branch on known conditions without string matching.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Sentinel errors such as `io.EOF`; `fmt.Errorf` with `%w`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Sentinel errors such as `io.EOF`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Sentinel errors such as `io.EOF`"
```

### 2. `fmt.Errorf` with `%w`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`fmt.Errorf` with `%w`"
```

### 3. `errors.Is` and error chains

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`errors.Is` and error chains"
```

### 4. `errors.As` and typed errors

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`errors.As` and typed errors"
```

### 5. When not to wrap

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When not to wrap"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Sentinel errors such as `io.EOF`. Predict the output before running.
2. **Change** — modify one detail in `fmt.Errorf` with `%w`. Explain the difference.
3. **Break** — break When not to wrap on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Sentinel errors such as `io.EOF` and `fmt.Errorf` with `%w` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using string comparison instead of `errors.Is`. Explain the cause and the fix in your own words.
- **Mistake:** Wrapping the same error twice and duplicating context. Explain the cause and the fix in your own words.
- **Mistake:** Returning `nil` inside a non-nil error interface. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Sentinel errors such as `io.EOF` and `fmt.Errorf` with `%w`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around When not to wrap.
3. Explain Wrapping and Sentinel Errors aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Sentinel errors such as `io.EOF` without opening this note.
2. Compare `fmt.Errorf` with `%w` with an alternative. When is each appropriate?
3. What happens at the edge case of When not to wrap, and how do you detect it?
4. Where does Wrapping and Sentinel Errors appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go blog: error handling and Go](https://go.dev/blog/error-handling-and-go)
- [ ] [Working with errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [ ] [Go by Example: errors](https://gobyexample.com/errors)
- [ ] [The Go specification: errors](https://go.dev/ref/spec#Errors)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
