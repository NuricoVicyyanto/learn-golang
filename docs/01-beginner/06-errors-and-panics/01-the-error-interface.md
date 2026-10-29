# 06.1 The Error Interface

> **Level:** Beginner · **Module:** 06 Errors and Panics · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

`error` is just an interface with an `Error() string` method. Understanding that makes the rest of Go error handling unsurprising.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `error` interface and `errors.New`; Returning errors as the last value.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `error` interface and `errors.New`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `error` interface and `errors.New`"
```

### 2. Returning errors as the last value

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Returning errors as the last value"
```

### 3. Checking errors immediately

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Checking errors immediately"
```

### 4. Error messages: lowercase, no punctuation, actionable

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Error messages: lowercase, no punctuation, actionable"
```

### 5. When to log vs when to return

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When to log vs when to return"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for The `error` interface and `errors.New`. Predict the output before running.
2. **Change** — modify one detail in Returning errors as the last value. Explain the difference.
3. **Break** — break When to log vs when to return on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use The `error` interface and `errors.New` and Returning errors as the last value in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Ignoring an error because the happy path works. Explain the cause and the fix in your own words.
- **Mistake:** Logging and returning the same error at every layer. Explain the cause and the fix in your own words.
- **Mistake:** Writing error messages that do not identify the operation. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses The `error` interface and `errors.New` and Returning errors as the last value. Keep it under 40 lines.
2. Add a table-driven test for the edge case around When to log vs when to return.
3. Explain The Error Interface aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain The `error` interface and `errors.New` without opening this note.
2. Compare Returning errors as the last value with an alternative. When is each appropriate?
3. What happens at the edge case of When to log vs when to return, and how do you detect it?
4. Where does The Error Interface appear in a production system?

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
