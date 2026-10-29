# 06.3 Custom Errors

> **Level:** Beginner · **Module:** 06 Errors and Panics · **Status:** ⬜ Not started
> **Estimated time:** 60 min

## Why it matters

Custom error types carry structured data such as field names, codes, and retryability. They make APIs easier to consume and test.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Defining an error type with fields; Implementing `Error()` on a value or pointer.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Defining an error type with fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Defining an error type with fields"
```

### 2. Implementing `Error()` on a value or pointer

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Implementing `Error()` on a value or pointer"
```

### 3. Unwrapping and joining errors

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Unwrapping and joining errors"
```

### 4. Structured errors for APIs

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Structured errors for APIs"
```

### 5. Error codes vs error types

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Error codes vs error types"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Defining an error type with fields. Predict the output before running.
2. **Change** — modify one detail in Implementing `Error()` on a value or pointer. Explain the difference.
3. **Break** — break Error codes vs error types on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Defining an error type with fields and Implementing `Error()` on a value or pointer in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Putting secrets or raw SQL in error messages. Explain the cause and the fix in your own words.
- **Mistake:** Returning a custom error by value and losing type information. Explain the cause and the fix in your own words.
- **Mistake:** Creating a new error type for every trivial case. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Defining an error type with fields and Implementing `Error()` on a value or pointer. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Error codes vs error types.
3. Explain Custom Errors aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Defining an error type with fields without opening this note.
2. Compare Implementing `Error()` on a value or pointer with an alternative. When is each appropriate?
3. What happens at the edge case of Error codes vs error types, and how do you detect it?
4. Where does Custom Errors appear in a production system?

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
