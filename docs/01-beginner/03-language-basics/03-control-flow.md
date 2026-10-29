# 03.3 Control Flow

> **Level:** Beginner · **Module:** 03 Language Basics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Go has fewer control-flow constructs than most languages, but they compose well. You need to know the initializer form, expressionless switch, and the three loop forms.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `if`, `else if`, `else`, and the initializer statement; `switch`, expressionless switch, and the no-fallthrough default.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `if`, `else if`, `else`, and the initializer statement

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`if`, `else if`, `else`, and the initializer statement"
```

### 2. `switch`, expressionless switch, and the no-fallthrough default

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`switch`, expressionless switch, and the no-fallthrough default"
```

### 3. The three forms of `for`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The three forms of `for`"
```

### 4. `range` over slices, maps, strings, and channels

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`range` over slices, maps, strings, and channels"
```

### 5. `break`, `continue`, labels, and when to avoid `goto`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`break`, `continue`, labels, and when to avoid `goto`"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `if`, `else if`, `else`, and the initializer statement. Predict the output before running.
2. **Change** - modify one detail in `switch`, expressionless switch, and the no-fallthrough default. Explain the difference.
3. **Break** - break `break`, `continue`, labels, and when to avoid `goto` on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `if`, `else if`, `else`, and the initializer statement and `switch`, expressionless switch, and the no-fallthrough default in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Loop-variable capture in goroutines or closures before Go 1.22. Explain the cause and the fix in your own words.
- **Mistake:** Off-by-one errors in the classic `for` loop. Explain the cause and the fix in your own words.
- **Mistake:** Mutating a slice while ranging over it. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `if`, `else if`, `else`, and the initializer statement and `switch`, expressionless switch, and the no-fallthrough default. Keep it under 40 lines.
2. Add a table-driven test for the edge case around `break`, `continue`, labels, and when to avoid `goto`.
3. Explain Control Flow aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `if`, `else if`, `else`, and the initializer statement without opening this note.
2. Compare `switch`, expressionless switch, and the no-fallthrough default with an alternative. When is each appropriate?
3. What happens at the edge case of `break`, `continue`, labels, and when to avoid `goto`, and how do you detect it?
4. Where does Control Flow appear in a production system?

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
