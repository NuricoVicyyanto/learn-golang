# 08.3 Sort and Slices

> **Level:** Beginner · **Module:** 08 Standard Library Essentials · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Modern Go provides generic slice helpers, but you still need to understand ordering, stability, and search preconditions.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `sort.Interface` and `sort.Slice`; The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `sort.Interface` and `sort.Slice`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`sort.Interface` and `sort.Slice`"
```

### 2. The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index`"
```

### 3. Stable vs unstable sorting

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Stable vs unstable sorting"
```

### 4. Binary search and sorted preconditions

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Binary search and sorted preconditions"
```

### 5. Sorting structs by multiple fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Sorting structs by multiple fields"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `sort.Interface` and `sort.Slice`. Predict the output before running.
2. **Change** - modify one detail in The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index`. Explain the difference.
3. **Break** - break Sorting structs by multiple fields on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `sort.Interface` and `sort.Slice` and The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Binary-searching an unsorted slice. Explain the cause and the fix in your own words.
- **Mistake:** Comparing floats in a sort comparator without a total order. Explain the cause and the fix in your own words.
- **Mistake:** Mutating the slice while sorting. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `sort.Interface` and `sort.Slice` and The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Sorting structs by multiple fields.
3. Explain Sort and Slices aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `sort.Interface` and `sort.Slice` without opening this note.
2. Compare The `slices` package: `Sort`, `SortFunc`, `Contains`, `Index` with an alternative. When is each appropriate?
3. What happens at the edge case of Sorting structs by multiple fields, and how do you detect it?
4. Where does Sort and Slices appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package strings](https://pkg.go.dev/strings)
- [ ] [Package strconv](https://pkg.go.dev/strconv)
- [ ] [Package time](https://pkg.go.dev/time)
- [ ] [Package slices](https://pkg.go.dev/slices)
- [ ] [Package io](https://pkg.go.dev/io)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
