# 03.1 Variables, Constants, and Types

> **Level:** Beginner · **Module:** 03 Language Basics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Every Go program is built from typed values. Understanding zero values, inference, constants, and conversion prevents a large class of compile-time and runtime bugs.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `var` declarations vs short variable declarations (`:=`); Zero values and type inference.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `var` declarations vs short variable declarations (`:=`)

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`var` declarations vs short variable declarations (`:=`)"
```

### 2. Zero values and type inference

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Zero values and type inference"
```

### 3. Constants, `iota`, and typed vs untyped constants

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Constants, `iota`, and typed vs untyped constants"
```

### 4. Basic types: bool, integers, floats, strings, runes, bytes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Basic types: bool, integers, floats, strings, runes, bytes"
```

### 5. Type conversion, overflow, and `strconv`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Type conversion, overflow, and `strconv`"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `var` declarations vs short variable declarations (`:=`). Predict the output before running.
2. **Change** - modify one detail in Zero values and type inference. Explain the difference.
3. **Break** - break Type conversion, overflow, and `strconv` on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `var` declarations vs short variable declarations (`:=`) and Zero values and type inference in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using `:=` outside a function or accidentally redeclaring a variable. Explain the cause and the fix in your own words.
- **Mistake:** Mixing numeric types without an explicit conversion. Explain the cause and the fix in your own words.
- **Mistake:** Assuming a string index returns a character rather than a byte. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `var` declarations vs short variable declarations (`:=`) and Zero values and type inference. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Type conversion, overflow, and `strconv`.
3. Explain Variables, Constants, and Types aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `var` declarations vs short variable declarations (`:=`) without opening this note.
2. Compare Zero values and type inference with an alternative. When is each appropriate?
3. What happens at the edge case of Type conversion, overflow, and `strconv`, and how do you detect it?
4. Where does Variables, Constants, and Types appear in a production system?

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
