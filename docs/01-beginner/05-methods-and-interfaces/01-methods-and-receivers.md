# 05.1 Methods and Receivers

> **Level:** Beginner · **Module:** 05 Methods and Interfaces · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Methods attach behavior to named types. The receiver choice affects mutability, copying, and which interfaces the type satisfies.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Declaring methods on named types; Value vs pointer receivers.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Declaring methods on named types

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Declaring methods on named types"
```

### 2. Value vs pointer receivers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Value vs pointer receivers"
```

### 3. Method sets and interface satisfaction

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Method sets and interface satisfaction"
```

### 4. Choosing a receiver type consistently

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Choosing a receiver type consistently"
```

### 5. Methods on slices, maps, and other named types

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Methods on slices, maps, and other named types"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Declaring methods on named types. Predict the output before running.
2. **Change** — modify one detail in Value vs pointer receivers. Explain the difference.
3. **Break** — break Methods on slices, maps, and other named types on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Declaring methods on named types and Value vs pointer receivers in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Mixing value and pointer receivers for the same type. Explain the cause and the fix in your own words.
- **Mistake:** Mutating a value receiver and expecting the caller to see it. Explain the cause and the fix in your own words.
- **Mistake:** Calling a pointer method on a non-addressable value. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Declaring methods on named types and Value vs pointer receivers. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Methods on slices, maps, and other named types.
3. Explain Methods and Receivers aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Declaring methods on named types without opening this note.
2. Compare Value vs pointer receivers with an alternative. When is each appropriate?
3. What happens at the edge case of Methods on slices, maps, and other named types, and how do you detect it?
4. Where does Methods and Receivers appear in a production system?

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
