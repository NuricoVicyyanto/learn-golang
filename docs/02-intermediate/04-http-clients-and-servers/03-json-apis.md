# 04.3 JSON APIs

> **Level:** Intermediate · **Module:** 04 HTTP Clients and Servers · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

An API contract is more than JSON tags. Validation, error shape, pagination, and versioning determine whether clients can use it reliably.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Decoding requests with size limits; Encoding responses and setting content type.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Decoding requests with size limits

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Decoding requests with size limits"
```

### 2. Encoding responses and setting content type

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Encoding responses and setting content type"
```

### 3. Validation and field-level errors

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Validation and field-level errors"
```

### 4. Consistent error envelopes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Consistent error envelopes"
```

### 5. Pagination, filtering, and versioning

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pagination, filtering, and versioning"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Decoding requests with size limits. Predict the output before running.
2. **Change** — modify one detail in Encoding responses and setting content type. Explain the difference.
3. **Break** — break Pagination, filtering, and versioning on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Decoding requests with size limits and Encoding responses and setting content type in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Decoding unbounded request bodies. Explain the cause and the fix in your own words.
- **Mistake:** Leaking internal error messages. Explain the cause and the fix in your own words.
- **Mistake:** Changing response shapes without versioning. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Decoding requests with size limits and Encoding responses and setting content type. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Pagination, filtering, and versioning.
3. Document one trade-off and one failure mode of JSON APIs.

## Self-check

1. Explain Decoding requests with size limits without opening this note.
2. Compare Encoding responses and setting content type with an alternative. When is each appropriate?
3. What happens at the edge case of Pagination, filtering, and versioning, and how do you detect it?
4. Where does JSON APIs appear in a production system?

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
