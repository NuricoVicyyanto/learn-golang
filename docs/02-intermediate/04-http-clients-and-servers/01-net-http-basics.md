# 04.1 net/http Basics

> **Level:** Intermediate · **Module:** 04 HTTP Clients and Servers · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

The standard library HTTP server is production-grade. Understanding handlers, muxes, and server timeouts removes the need for a framework in many services.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `http.Handler` and `http.HandlerFunc`; `http.ServeMux` and path patterns.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `http.Handler` and `http.HandlerFunc`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`http.Handler` and `http.HandlerFunc`"
```

### 2. `http.ServeMux` and path patterns

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`http.ServeMux` and path patterns"
```

### 3. Request parsing and `ResponseWriter`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Request parsing and `ResponseWriter`"
```

### 4. Status codes and headers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Status codes and headers"
```

### 5. Server timeouts and graceful shutdown

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Server timeouts and graceful shutdown"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `http.Handler` and `http.HandlerFunc`. Predict the output before running.
2. **Change** — modify one detail in `http.ServeMux` and path patterns. Explain the difference.
3. **Break** — break Server timeouts and graceful shutdown on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `http.Handler` and `http.HandlerFunc` and `http.ServeMux` and path patterns in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using the default mux in a library. Explain the cause and the fix in your own words.
- **Mistake:** Omitting read, write, and idle timeouts. Explain the cause and the fix in your own words.
- **Mistake:** Writing headers after the body. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `http.Handler` and `http.HandlerFunc` and `http.ServeMux` and path patterns. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Server timeouts and graceful shutdown.
3. Document one trade-off and one failure mode of net/http Basics.

## Self-check

1. Explain `http.Handler` and `http.HandlerFunc` without opening this note.
2. Compare `http.ServeMux` and path patterns with an alternative. When is each appropriate?
3. What happens at the edge case of Server timeouts and graceful shutdown, and how do you detect it?
4. Where does net/http Basics appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package net/http](https://pkg.go.dev/net/http)
- [ ] [Go blog: writing web applications](https://go.dev/doc/articles/wiki/)
- [ ] [Go docs: HTTP servers](https://pkg.go.dev/net/http#Server)
- [ ] [Go blog: context](https://go.dev/blog/context)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
