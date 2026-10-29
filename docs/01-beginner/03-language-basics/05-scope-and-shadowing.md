# 03.5 Scope and Shadowing

> **Level:** Beginner · **Module:** 03 Language Basics · **Status:** ⬜ Not started
> **Estimated time:** 60 min

## Why it matters

Scope determines what a name means. Shadowing is one of the most common sources of confusing bugs in Go, especially around `err` and loop variables.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Block scope and declaration order; Package-level scope and initialization order.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Block scope and declaration order

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Block scope and declaration order"
```

### 2. Package-level scope and initialization order

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Package-level scope and initialization order"
```

### 3. Shadowing and the `:=` trap

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Shadowing and the `:=` trap"
```

### 4. Lifetime vs visibility

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Lifetime vs visibility"
```

### 5. Naming and package-level state

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Naming and package-level state"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Block scope and declaration order. Predict the output before running.
2. **Change** — modify one detail in Package-level scope and initialization order. Explain the difference.
3. **Break** — break Naming and package-level state on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Block scope and declaration order and Package-level scope and initialization order in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Shadowing `err` in an inner block and losing the original error. Explain the cause and the fix in your own words.
- **Mistake:** Assuming package initialization order across files. Explain the cause and the fix in your own words.
- **Mistake:** Using package globals as hidden dependencies. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Block scope and declaration order and Package-level scope and initialization order. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Naming and package-level state.
3. Explain Scope and Shadowing aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Block scope and declaration order without opening this note.
2. Compare Package-level scope and initialization order with an alternative. When is each appropriate?
3. What happens at the edge case of Naming and package-level state, and how do you detect it?
4. Where does Scope and Shadowing appear in a production system?

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
