# 10.2 Table-Driven Tests

> **Level:** Beginner · **Module:** 10 Testing Basics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Table-driven tests scale with cases and make failures readable. Subtests and parallel execution keep large suites fast.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Table structs and test cases; Subtests with `t.Run`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Table structs and test cases

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Table structs and test cases"
```

### 2. Subtests with `t.Run`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Subtests with `t.Run`"
```

### 3. `t.Parallel` and loop-variable capture

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`t.Parallel` and loop-variable capture"
```

### 4. Test helpers with `t.Helper`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Test helpers with `t.Helper`"
```

### 5. Golden files for large outputs

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Golden files for large outputs"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Table structs and test cases. Predict the output before running.
2. **Change** — modify one detail in Subtests with `t.Run`. Explain the difference.
3. **Break** — break Golden files for large outputs on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Table structs and test cases and Subtests with `t.Run` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Sharing mutable state between parallel subtests. Explain the cause and the fix in your own words.
- **Mistake:** Omitting the case name from failure output. Explain the cause and the fix in your own words.
- **Mistake:** Treating a table as a substitute for thinking about edge cases. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Table structs and test cases and Subtests with `t.Run`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Golden files for large outputs.
3. Explain Table-Driven Tests aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Table structs and test cases without opening this note.
2. Compare Subtests with `t.Run` with an alternative. When is each appropriate?
3. What happens at the edge case of Golden files for large outputs, and how do you detect it?
4. Where does Table-Driven Tests appear in a production system?

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
