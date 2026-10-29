# 02.1 Context Basics

> **Level:** Intermediate · **Module:** 02 Context and Cancellation · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Context is a tree of cancellation signals and request-scoped values. Misusing it creates hidden coupling or leaks.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The context tree and parent-child relationships; `context.Background` vs `context.TODO`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The context tree and parent-child relationships

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The context tree and parent-child relationships"
```

### 2. `context.Background` vs `context.TODO`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`context.Background` vs `context.TODO`"
```

### 3. Request-scoped values and why to use them sparingly

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Request-scoped values and why to use them sparingly"
```

### 4. Passing context as the first parameter

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Passing context as the first parameter"
```

### 5. Never storing context in a struct

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Never storing context in a struct"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for The context tree and parent-child relationships. Predict the output before running.
2. **Change** — modify one detail in `context.Background` vs `context.TODO`. Explain the difference.
3. **Break** — break Never storing context in a struct on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use The context tree and parent-child relationships and `context.Background` vs `context.TODO` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using context values for required dependencies. Explain the cause and the fix in your own words.
- **Mistake:** Storing a context in a struct field. Explain the cause and the fix in your own words.
- **Mistake:** Passing `nil` context. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies The context tree and parent-child relationships and `context.Background` vs `context.TODO`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Never storing context in a struct.
3. Document one trade-off and one failure mode of Context Basics.

## Self-check

1. Explain The context tree and parent-child relationships without opening this note.
2. Compare `context.Background` vs `context.TODO` with an alternative. When is each appropriate?
3. What happens at the edge case of Never storing context in a struct, and how do you detect it?
4. Where does Context Basics appear in a production system?

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
