# 04.1 Arrays and Slices

> **Level:** Beginner · **Module:** 04 Composite Types · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Slices are used in almost every Go program. Their header, capacity, and aliasing behavior are the source of both their power and their most common bugs.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Arrays have a fixed length that is part of their type; The slice header: pointer, length, and capacity.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Arrays have a fixed length that is part of their type

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Arrays have a fixed length that is part of their type"
```

### 2. The slice header: pointer, length, and capacity

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The slice header: pointer, length, and capacity"
```

### 3. `append`, reallocation, and copy semantics

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`append`, reallocation, and copy semantics"
```

### 4. Slicing, aliasing, and the full slice expression

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Slicing, aliasing, and the full slice expression"
```

### 5. `copy`, `slices.Clone`, and removing elements

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`copy`, `slices.Clone`, and removing elements"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Arrays have a fixed length that is part of their type. Predict the output before running.
2. **Change** - modify one detail in The slice header: pointer, length, and capacity. Explain the difference.
3. **Break** - break `copy`, `slices.Clone`, and removing elements on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Arrays have a fixed length that is part of their type and The slice header: pointer, length, and capacity in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Assuming `append` always copies the underlying array. Explain the cause and the fix in your own words.
- **Mistake:** Holding a small subslice that keeps a large array alive. Explain the cause and the fix in your own words.
- **Mistake:** Removing elements while ranging over the same slice. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Arrays have a fixed length that is part of their type and The slice header: pointer, length, and capacity. Keep it under 40 lines.
2. Add a table-driven test for the edge case around `copy`, `slices.Clone`, and removing elements.
3. Explain Arrays and Slices aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Arrays have a fixed length that is part of their type without opening this note.
2. Compare The slice header: pointer, length, and capacity with an alternative. When is each appropriate?
3. What happens at the edge case of `copy`, `slices.Clone`, and removing elements, and how do you detect it?
4. Where does Arrays and Slices appear in a production system?

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
