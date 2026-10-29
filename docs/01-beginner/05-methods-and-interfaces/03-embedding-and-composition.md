# 05.3 Embedding and Composition

> **Level:** Beginner · **Module:** 05 Methods and Interfaces · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Go favors composition over inheritance. Embedding promotes fields and methods, but it is not subclassing; confusing the two leads to fragile designs.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Embedding structs and promoted fields; Embedding interfaces and method promotion.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Embedding structs and promoted fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Embedding structs and promoted fields"
```

### 2. Embedding interfaces and method promotion

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Embedding interfaces and method promotion"
```

### 3. Composition with explicit fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Composition with explicit fields"
```

### 4. Method resolution and name conflicts

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Method resolution and name conflicts"
```

### 5. When embedding improves an API and when it leaks internals

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When embedding improves an API and when it leaks internals"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Embedding structs and promoted fields. Predict the output before running.
2. **Change** — modify one detail in Embedding interfaces and method promotion. Explain the difference.
3. **Break** — break When embedding improves an API and when it leaks internals on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Embedding structs and promoted fields and Embedding interfaces and method promotion in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Treating embedding as inheritance or polymorphism. Explain the cause and the fix in your own words.
- **Mistake:** Exposing unwanted methods through embedding. Explain the cause and the fix in your own words.
- **Mistake:** Ambiguous selectors after embedding two types. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Embedding structs and promoted fields and Embedding interfaces and method promotion. Keep it under 40 lines.
2. Add a table-driven test for the edge case around When embedding improves an API and when it leaks internals.
3. Explain Embedding and Composition aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Embedding structs and promoted fields without opening this note.
2. Compare Embedding interfaces and method promotion with an alternative. When is each appropriate?
3. What happens at the edge case of When embedding improves an API and when it leaks internals, and how do you detect it?
4. Where does Embedding and Composition appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Effective Go: interfaces](https://go.dev/doc/effective_go#interfaces)
- [ ] [Go blog: interfaces](https://go.dev/blog/)
- [ ] [Go by Example: interfaces](https://gobyexample.com/interfaces)
- [ ] [The Go specification: method sets](https://go.dev/ref/spec#Method_sets)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
