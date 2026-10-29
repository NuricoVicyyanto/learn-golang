# 09.3 CI/CD

> **Level:** Intermediate · **Module:** 09 Docker and Deployment · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Automated checks are how a team keeps quality while moving fast. The pipeline should be fast, deterministic, and hard to bypass.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Pipeline stages: lint, test, build, scan, deploy; Caching modules and build layers.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Pipeline stages: lint, test, build, scan, deploy

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pipeline stages: lint, test, build, scan, deploy"
```

### 2. Caching modules and build layers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Caching modules and build layers"
```

### 3. Test matrix and race detection

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Test matrix and race detection"
```

### 4. Versioning and release artifacts

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Versioning and release artifacts"
```

### 5. Secrets and environment promotion

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Secrets and environment promotion"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Pipeline stages: lint, test, build, scan, deploy. Predict the output before running.
2. **Change** - modify one detail in Caching modules and build layers. Explain the difference.
3. **Break** - break Secrets and environment promotion on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Pipeline stages: lint, test, build, scan, deploy and Caching modules and build layers in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Deploying without running tests. Explain the cause and the fix in your own words.
- **Mistake:** Storing secrets in workflow files. Explain the cause and the fix in your own words.
- **Mistake:** Publishing mutable `latest`-only images. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Pipeline stages: lint, test, build, scan, deploy and Caching modules and build layers. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Secrets and environment promotion.
3. Document one trade-off and one failure mode of CI/CD.

## Self-check

1. Explain Pipeline stages: lint, test, build, scan, deploy without opening this note.
2. Compare Caching modules and build layers with an alternative. When is each appropriate?
3. What happens at the edge case of Secrets and environment promotion, and how do you detect it?
4. Where does CI/CD appear in a production system?

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
