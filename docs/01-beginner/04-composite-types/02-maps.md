# 04.2 Maps

> **Level:** Beginner · **Module:** 04 Composite Types · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Maps model lookup tables, counters, caches, and indexes. You need to understand missing keys, iteration order, and why maps are not safe for concurrent writes.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Map creation, the zero value, and `make`; The comma-ok idiom and missing keys.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Map creation, the zero value, and `make`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Map creation, the zero value, and `make`"
```

### 2. The comma-ok idiom and missing keys

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The comma-ok idiom and missing keys"
```

### 3. Iteration order and non-determinism

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Iteration order and non-determinism"
```

### 4. Maps are reference-like; assignment does not copy entries

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Maps are reference-like; assignment does not copy entries"
```

### 5. Concurrent map access and synchronization

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Concurrent map access and synchronization"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Map creation, the zero value, and `make`. Predict the output before running.
2. **Change** — modify one detail in The comma-ok idiom and missing keys. Explain the difference.
3. **Break** — break Concurrent map access and synchronization on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Map creation, the zero value, and `make` and The comma-ok idiom and missing keys in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Writing to a nil map. Explain the cause and the fix in your own words.
- **Mistake:** Relying on map iteration order. Explain the cause and the fix in your own words.
- **Mistake:** Reading and writing a map from multiple goroutines without a lock. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Map creation, the zero value, and `make` and The comma-ok idiom and missing keys. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Concurrent map access and synchronization.
3. Explain Maps aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Map creation, the zero value, and `make` without opening this note.
2. Compare The comma-ok idiom and missing keys with an alternative. When is each appropriate?
3. What happens at the edge case of Concurrent map access and synchronization, and how do you detect it?
4. Where does Maps appear in a production system?

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
