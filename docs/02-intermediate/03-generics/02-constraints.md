# 03.2 Constraints

> **Level:** Intermediate · **Module:** 03 Generics · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Constraints define what operations a type parameter supports. They are the contract that makes generic code compile.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: The `comparable` constraint; Union constraints and `~` approximation.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. The `comparable` constraint

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "The `comparable` constraint"
```

### 2. Union constraints and `~` approximation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Union constraints and `~` approximation"
```

### 3. Custom constraint interfaces

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Custom constraint interfaces"
```

### 4. `any` vs `interface{}`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`any` vs `interface{}`"
```

### 5. Constraint design and documentation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Constraint design and documentation"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for The `comparable` constraint. Predict the output before running.
2. **Change** — modify one detail in Union constraints and `~` approximation. Explain the difference.
3. **Break** — break Constraint design and documentation on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use The `comparable` constraint and Union constraints and `~` approximation in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Forgetting `~` for named types. Explain the cause and the fix in your own words.
- **Mistake:** Putting methods in a constraint and calling them incorrectly. Explain the cause and the fix in your own words.
- **Mistake:** Using `any` and losing all type safety. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies The `comparable` constraint and Union constraints and `~` approximation. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Constraint design and documentation.
3. Document one trade-off and one failure mode of Constraints.

## Self-check

1. Explain The `comparable` constraint without opening this note.
2. Compare Union constraints and `~` approximation with an alternative. When is each appropriate?
3. What happens at the edge case of Constraint design and documentation, and how do you detect it?
4. Where does Constraints appear in a production system?

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
