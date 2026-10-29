# 05.4 Type Assertions and Switches

> **Level:** Beginner · **Module:** 05 Methods and Interfaces · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

When you must recover a concrete type from an interface, assertions and type switches are the safe tools. Panics and reflection are for exceptional cases.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Type assertions and the comma-ok form; Type switches.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Type assertions and the comma-ok form

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Type assertions and the comma-ok form"
```

### 2. Type switches

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Type switches"
```

### 3. When an assertion can panic

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When an assertion can panic"
```

### 4. Interfaces and dynamic dispatch

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Interfaces and dynamic dispatch"
```

### 5. Reflection as a last resort

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Reflection as a last resort"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Type assertions and the comma-ok form. Predict the output before running.
2. **Change** - modify one detail in Type switches. Explain the difference.
3. **Break** - break Reflection as a last resort on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Type assertions and the comma-ok form and Type switches in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using a single-value assertion on data that may not match. Explain the cause and the fix in your own words.
- **Mistake:** Switching on concrete types when an interface method would do. Explain the cause and the fix in your own words.
- **Mistake:** Using reflection for a problem generics can solve. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Type assertions and the comma-ok form and Type switches. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Reflection as a last resort.
3. Explain Type Assertions and Switches aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Type assertions and the comma-ok form without opening this note.
2. Compare Type switches with an alternative. When is each appropriate?
3. What happens at the edge case of Reflection as a last resort, and how do you detect it?
4. Where does Type Assertions and Switches appear in a production system?

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
