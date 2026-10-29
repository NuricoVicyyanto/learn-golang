# 03.1 Type Parameters

> **Level:** Intermediate · **Module:** 03 Generics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Type parameters let one implementation work across types while keeping compile-time safety. The syntax is small; the design judgment is the hard part.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Generic function syntax and instantiation; Type inference at call sites.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Generic function syntax and instantiation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Generic function syntax and instantiation"
```

### 2. Type inference at call sites

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Type inference at call sites"
```

### 3. Zero values of a type parameter

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Zero values of a type parameter"
```

### 4. Multiple type parameters

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Multiple type parameters"
```

### 5. When generics improve or hurt readability

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When generics improve or hurt readability"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Generic function syntax and instantiation. Predict the output before running.
2. **Change** — modify one detail in Type inference at call sites. Explain the difference.
3. **Break** — break When generics improve or hurt readability on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Generic function syntax and instantiation and Type inference at call sites in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using generics where a concrete type or interface is simpler. Explain the cause and the fix in your own words.
- **Mistake:** Assuming you can compare any type parameter with `==`. Explain the cause and the fix in your own words.
- **Mistake:** Over-constraining a function until it is unusable. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Generic function syntax and instantiation and Type inference at call sites. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises When generics improve or hurt readability.
3. Document one trade-off and one failure mode of Type Parameters.

## Self-check

1. Explain Generic function syntax and instantiation without opening this note.
2. Compare Type inference at call sites with an alternative. When is each appropriate?
3. What happens at the edge case of When generics improve or hurt readability, and how do you detect it?
4. Where does Type Parameters appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go blog: generics](https://go.dev/blog/intro-generics)
- [ ] [Go spec: type parameters](https://go.dev/ref/spec#Type_parameter_declarations)
- [ ] [When to use generics](https://go.dev/blog/when-generics)
- [ ] [Package slices](https://pkg.go.dev/slices)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
