# 10.1 Your First Test

> **Level:** Beginner · **Module:** 10 Testing Basics · **Status:** ⬜ Not started
> **Estimated time:** 60 min

## Why it matters

A test is a small program that checks behavior. Starting with the mechanics makes the rest of testing feel routine.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `testing` package and `TestXxx` naming; `t.Errorf` vs `t.Fatalf`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `testing` package and `TestXxx` naming

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `testing` package and `TestXxx` naming"
```

### 2. `t.Errorf` vs `t.Fatalf`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`t.Errorf` vs `t.Fatalf`"
```

### 3. Running tests: package, pattern, verbose

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Running tests: package, pattern, verbose"
```

### 4. Test file placement and package choice

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Test file placement and package choice"
```

### 5. Test output and failure messages

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Test output and failure messages"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for The `testing` package and `TestXxx` naming. Predict the output before running.
2. **Change** - modify one detail in `t.Errorf` vs `t.Fatalf`. Explain the difference.
3. **Break** - break Test output and failure messages on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use The `testing` package and `TestXxx` naming and `t.Errorf` vs `t.Fatalf` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Testing implementation details instead of behavior. Explain the cause and the fix in your own words.
- **Mistake:** Calling `t.Fatal` from a goroutine. Explain the cause and the fix in your own words.
- **Mistake:** Writing tests that pass for the wrong reason. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses The `testing` package and `TestXxx` naming and `t.Errorf` vs `t.Fatalf`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Test output and failure messages.
3. Explain Your First Test aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain The `testing` package and `TestXxx` naming without opening this note.
2. Compare `t.Errorf` vs `t.Fatalf` with an alternative. When is each appropriate?
3. What happens at the edge case of Test output and failure messages, and how do you detect it?
4. Where does Your First Test appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package testing](https://pkg.go.dev/testing)
- [ ] [Go blog: table-driven tests](https://go.dev/blog/subtests)
- [ ] [Go docs: coverage](https://go.dev/doc/build-cover)
- [ ] [Go by Example: testing](https://gobyexample.com/testing)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
