# 02.3 Context in HTTP and Databases

> **Level:** Intermediate · **Module:** 02 Context and Cancellation · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

HTTP requests and database queries already support context. Connecting them correctly is the difference between responsive and stuck services.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `http.Request.Context`; Context-aware database methods.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `http.Request.Context`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`http.Request.Context`"
```

### 2. Context-aware database methods

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Context-aware database methods"
```

### 3. Timeouts at the server, handler, and query level

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Timeouts at the server, handler, and query level"
```

### 4. Graceful shutdown with context

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Graceful shutdown with context"
```

### 5. Testing cancellation with fake slow dependencies

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Testing cancellation with fake slow dependencies"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `http.Request.Context`. Predict the output before running.
2. **Change** - modify one detail in Context-aware database methods. Explain the difference.
3. **Break** - break Testing cancellation with fake slow dependencies on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `http.Request.Context` and Context-aware database methods in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Creating a new background context inside a handler. Explain the cause and the fix in your own words.
- **Mistake:** Setting a server timeout longer than the client timeout. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring context in retry loops. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `http.Request.Context` and Context-aware database methods. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Testing cancellation with fake slow dependencies.
3. Document one trade-off and one failure mode of Context in HTTP and Databases.

## Self-check

1. Explain `http.Request.Context` without opening this note.
2. Compare Context-aware database methods with an alternative. When is each appropriate?
3. What happens at the edge case of Testing cancellation with fake slow dependencies, and how do you detect it?
4. Where does Context in HTTP and Databases appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go blog: context](https://go.dev/blog/context)
- [ ] [Package context](https://pkg.go.dev/context)
- [ ] [Go blog: context and cancellation](https://go.dev/blog/context-and-structs)
- [ ] [net/http server timeouts](https://pkg.go.dev/net/http#Server)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
