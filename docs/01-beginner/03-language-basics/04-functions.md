# 03.4 Functions

> **Level:** Beginner · **Module:** 03 Language Basics · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Functions are Go's primary unit of behavior. Multiple returns, closures, and defer are used everywhere, including error handling, cleanup, and tests.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Signatures, parameters, and multiple return values; Named results and naked returns.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Signatures, parameters, and multiple return values

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Signatures, parameters, and multiple return values"
```

### 2. Named results and naked returns

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Named results and naked returns"
```

### 3. Variadic parameters and slice expansion

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Variadic parameters and slice expansion"
```

### 4. Function values, closures, and captured variables

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Function values, closures, and captured variables"
```

### 5. `defer`, argument evaluation, and LIFO order

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`defer`, argument evaluation, and LIFO order"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Signatures, parameters, and multiple return values. Predict the output before running.
2. **Change** - modify one detail in Named results and naked returns. Explain the difference.
3. **Break** - break `defer`, argument evaluation, and LIFO order on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Signatures, parameters, and multiple return values and Named results and naked returns in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Ignoring an error return. Explain the cause and the fix in your own words.
- **Mistake:** Deferring a function call inside a loop. Explain the cause and the fix in your own words.
- **Mistake:** Capturing a loop variable by reference unintentionally. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Signatures, parameters, and multiple return values and Named results and naked returns. Keep it under 40 lines.
2. Add a table-driven test for the edge case around `defer`, argument evaluation, and LIFO order.
3. Explain Functions aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Signatures, parameters, and multiple return values without opening this note.
2. Compare Named results and naked returns with an alternative. When is each appropriate?
3. What happens at the edge case of `defer`, argument evaluation, and LIFO order, and how do you detect it?
4. Where does Functions appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [The Go specification: declarations and types](https://go.dev/ref/spec)
- [ ] [A Tour of Go: basics](https://go.dev/tour/basics/1)
- [ ] [Effective Go](https://go.dev/doc/effective_go)
- [ ] [Go by Example](https://gobyexample.com/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
