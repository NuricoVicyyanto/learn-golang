# 08.1 Strings and strconv

> **Level:** Beginner · **Module:** 08 Standard Library Essentials · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Text processing is everywhere. Go strings are immutable byte sequences, and Unicode adds rune-level complexity that must be handled deliberately.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `strings` package: split, join, trim, replace, prefix/suffix; `strings.Builder` and efficient concatenation.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `strings` package: split, join, trim, replace, prefix/suffix

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `strings` package: split, join, trim, replace, prefix/suffix"
```

### 2. `strings.Builder` and efficient concatenation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`strings.Builder` and efficient concatenation"
```

### 3. `strconv` for numbers, booleans, and quoting

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`strconv` for numbers, booleans, and quoting"
```

### 4. UTF-8, runes, and byte vs character indexing

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "UTF-8, runes, and byte vs character indexing"
```

### 5. Case conversion and normalization caveats

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Case conversion and normalization caveats"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for The `strings` package: split, join, trim, replace, prefix/suffix. Predict the output before running.
2. **Change** - modify one detail in `strings.Builder` and efficient concatenation. Explain the difference.
3. **Break** - break Case conversion and normalization caveats on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use The `strings` package: split, join, trim, replace, prefix/suffix and `strings.Builder` and efficient concatenation in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Indexing a string and assuming one byte equals one character. Explain the cause and the fix in your own words.
- **Mistake:** Concatenating in a loop and creating quadratic work. Explain the cause and the fix in your own words.
- **Mistake:** Using `strconv.Atoi` without handling the error. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses The `strings` package: split, join, trim, replace, prefix/suffix and `strings.Builder` and efficient concatenation. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Case conversion and normalization caveats.
3. Explain Strings and strconv aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain The `strings` package: split, join, trim, replace, prefix/suffix without opening this note.
2. Compare `strings.Builder` and efficient concatenation with an alternative. When is each appropriate?
3. What happens at the edge case of Case conversion and normalization caveats, and how do you detect it?
4. Where does Strings and strconv appear in a production system?

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
