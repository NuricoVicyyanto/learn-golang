# 08.2 Time and Math

> **Level:** Beginner · **Module:** 08 Standard Library Essentials · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Time zones, durations, and monotonic clocks are easy to get wrong. Math edge cases such as NaN and infinity also deserve attention.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `time.Time`, `time.Duration`, and locations; Formatting and parsing with reference layouts.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `time.Time`, `time.Duration`, and locations

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`time.Time`, `time.Duration`, and locations"
```

### 2. Formatting and parsing with reference layouts

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Formatting and parsing with reference layouts"
```

### 3. Timers, tickers, and timeouts

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Timers, tickers, and timeouts"
```

### 4. Monotonic time and wall-clock changes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Monotonic time and wall-clock changes"
```

### 5. `math` constants, rounding, NaN, and infinity

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`math` constants, rounding, NaN, and infinity"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `time.Time`, `time.Duration`, and locations. Predict the output before running.
2. **Change** - modify one detail in Formatting and parsing with reference layouts. Explain the difference.
3. **Break** - break `math` constants, rounding, NaN, and infinity on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `time.Time`, `time.Duration`, and locations and Formatting and parsing with reference layouts in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using `==` on times instead of `Equal`. Explain the cause and the fix in your own words.
- **Mistake:** Formatting with the wrong reference layout. Explain the cause and the fix in your own words.
- **Mistake:** Assuming `time.Now()` is monotonic across serialization. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `time.Time`, `time.Duration`, and locations and Formatting and parsing with reference layouts. Keep it under 40 lines.
2. Add a table-driven test for the edge case around `math` constants, rounding, NaN, and infinity.
3. Explain Time and Math aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `time.Time`, `time.Duration`, and locations without opening this note.
2. Compare Formatting and parsing with reference layouts with an alternative. When is each appropriate?
3. What happens at the edge case of `math` constants, rounding, NaN, and infinity, and how do you detect it?
4. Where does Time and Math appear in a production system?

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
