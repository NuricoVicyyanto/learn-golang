# 04.4 Pointers

> **Level:** Beginner · **Module:** 04 Composite Types · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Pointers let you share data and express optionality, but they also create aliasing and nil hazards. Learn the minimum you need instead of using pointers everywhere.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Address-of (`&`) and dereference (`*`); Pointers vs values and copy semantics.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Address-of (`&`) and dereference (`*`)

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Address-of (`&`) and dereference (`*`)"
```

### 2. Pointers vs values and copy semantics

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pointers vs values and copy semantics"
```

### 3. nil pointers and nil checks

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "nil pointers and nil checks"
```

### 4. When to use pointers: mutation, large structs, optionality

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When to use pointers: mutation, large structs, optionality"
```

### 5. Pointers, structs, and constructor functions

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pointers, structs, and constructor functions"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Address-of (`&`) and dereference (`*`). Predict the output before running.
2. **Change** — modify one detail in Pointers vs values and copy semantics. Explain the difference.
3. **Break** — break Pointers, structs, and constructor functions on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Address-of (`&`) and dereference (`*`) and Pointers vs values and copy semantics in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Dereferencing a nil pointer. Explain the cause and the fix in your own words.
- **Mistake:** Returning a pointer to a loop variable. Explain the cause and the fix in your own words.
- **Mistake:** Using pointers everywhere and making code harder to reason about. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Address-of (`&`) and dereference (`*`) and Pointers vs values and copy semantics. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Pointers, structs, and constructor functions.
3. Explain Pointers aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Address-of (`&`) and dereference (`*`) without opening this note.
2. Compare Pointers vs values and copy semantics with an alternative. When is each appropriate?
3. What happens at the edge case of Pointers, structs, and constructor functions, and how do you detect it?
4. Where does Pointers appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go slices: usage and internals](https://go.dev/blog/slices-intro)
- [ ] [Go maps in action](https://go.dev/blog/maps)
- [ ] [Go by Example: slices, maps, structs](https://gobyexample.com/)
- [ ] [The Go specification: composite types](https://go.dev/ref/spec#Composite_types)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
