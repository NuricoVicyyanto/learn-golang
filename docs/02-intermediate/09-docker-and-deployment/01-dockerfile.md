# 09.1 Dockerfile for Go

> **Level:** Intermediate · **Module:** 09 Docker and Deployment · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Go produces static binaries, which enables tiny images. Multi-stage builds, non-root users, and cache-friendly layers are the baseline.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Multi-stage builds; Static binaries and `CGO_ENABLED`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Multi-stage builds

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Multi-stage builds"
```

### 2. Static binaries and `CGO_ENABLED`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Static binaries and `CGO_ENABLED`"
```

### 3. Non-root users and read-only filesystems

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Non-root users and read-only filesystems"
```

### 4. Layer caching and `.dockerignore`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Layer caching and `.dockerignore`"
```

### 5. Image size and vulnerability scanning

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Image size and vulnerability scanning"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Multi-stage builds. Predict the output before running.
2. **Change** — modify one detail in Static binaries and `CGO_ENABLED`. Explain the difference.
3. **Break** — break Image size and vulnerability scanning on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Multi-stage builds and Static binaries and `CGO_ENABLED` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Copying the entire repository into the build context. Explain the cause and the fix in your own words.
- **Mistake:** Running as root. Explain the cause and the fix in your own words.
- **Mistake:** Building without pinning the Go version. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Multi-stage builds and Static binaries and `CGO_ENABLED`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Image size and vulnerability scanning.
3. Document one trade-off and one failure mode of Dockerfile for Go.

## Self-check

1. Explain Multi-stage builds without opening this note.
2. Compare Static binaries and `CGO_ENABLED` with an alternative. When is each appropriate?
3. What happens at the edge case of Image size and vulnerability scanning, and how do you detect it?
4. Where does Dockerfile for Go appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Docker: multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
- [ ] [Docker Compose documentation](https://docs.docker.com/compose/)
- [ ] [Go docs: containerize an app](https://docs.docker.com/language/golang/)
- [ ] [GitHub Actions documentation](https://docs.github.com/actions)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
