# 07.1 Test Doubles

> **Level:** Intermediate · **Module:** 07 Advanced Testing · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

A test double should make a test faster and clearer. Over-mocking produces tests that pass while production breaks.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Fakes vs stubs vs mocks; Consumer-defined interfaces.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Fakes vs stubs vs mocks

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Fakes vs stubs vs mocks"
```

### 2. Consumer-defined interfaces

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Consumer-defined interfaces"
```

### 3. Fixtures and builders

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Fixtures and builders"
```

### 4. Assertions and `testify`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Assertions and `testify`"
```

### 5. When to use the real dependency

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "When to use the real dependency"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Fakes vs stubs vs mocks. Predict the output before running.
2. **Change** — modify one detail in Consumer-defined interfaces. Explain the difference.
3. **Break** — break When to use the real dependency on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Fakes vs stubs vs mocks and Consumer-defined interfaces in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Mocking every dependency and testing nothing real. Explain the cause and the fix in your own words.
- **Mistake:** Asserting on call counts instead of behavior. Explain the cause and the fix in your own words.
- **Mistake:** Sharing mutable fixtures between tests. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Fakes vs stubs vs mocks and Consumer-defined interfaces. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises When to use the real dependency.
3. Document one trade-off and one failure mode of Test Doubles.

## Self-check

1. Explain Fakes vs stubs vs mocks without opening this note.
2. Compare Consumer-defined interfaces with an alternative. When is each appropriate?
3. What happens at the edge case of When to use the real dependency, and how do you detect it?
4. Where does Test Doubles appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package httptest](https://pkg.go.dev/net/http/httptest)
- [ ] [Go blog: fuzzing](https://go.dev/blog/fuzz-beta)
- [ ] [Go docs: fuzzing](https://go.dev/doc/security/fuzz/)
- [ ] [testcontainers-go](https://golang.testcontainers.org/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
