# 05.2 Interfaces

> **Level:** Beginner · **Module:** 05 Methods and Interfaces · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Interfaces let you depend on behavior instead of concrete types. They are most useful when small and defined by the package that consumes them.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Implicit implementation: no `implements` keyword; An interface value holds a dynamic type and value.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Implicit implementation: no `implements` keyword

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Implicit implementation: no `implements` keyword"
```

### 2. An interface value holds a dynamic type and value

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "An interface value holds a dynamic type and value"
```

### 3. The empty interface and `any`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The empty interface and `any`"
```

### 4. Small interfaces and consumer-side design

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Small interfaces and consumer-side design"
```

### 5. nil interface values vs interfaces holding nil pointers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "nil interface values vs interfaces holding nil pointers"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Implicit implementation: no `implements` keyword. Predict the output before running.
2. **Change** - modify one detail in An interface value holds a dynamic type and value. Explain the difference.
3. **Break** - break nil interface values vs interfaces holding nil pointers on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Implicit implementation: no `implements` keyword and An interface value holds a dynamic type and value in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Defining huge interfaces that are hard to fake. Explain the cause and the fix in your own words.
- **Mistake:** Assuming an interface holding a nil pointer equals nil. Explain the cause and the fix in your own words.
- **Mistake:** Putting interfaces in the producer package by default. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Implicit implementation: no `implements` keyword and An interface value holds a dynamic type and value. Keep it under 40 lines.
2. Add a table-driven test for the edge case around nil interface values vs interfaces holding nil pointers.
3. Explain Interfaces aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Implicit implementation: no `implements` keyword without opening this note.
2. Compare An interface value holds a dynamic type and value with an alternative. When is each appropriate?
3. What happens at the edge case of nil interface values vs interfaces holding nil pointers, and how do you detect it?
4. Where does Interfaces appear in a production system?

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
