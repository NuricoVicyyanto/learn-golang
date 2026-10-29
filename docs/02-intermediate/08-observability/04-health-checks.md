# 08.4 Health Checks

> **Level:** Intermediate · **Module:** 08 Observability · **Status:** ⬜ Not started
> **Estimated time:** 60 min

## Why it matters

Kubernetes and load balancers rely on health endpoints. A wrong health check can restart a healthy service or keep routing to a broken one.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Liveness vs readiness; Dependency checks and timeouts.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Liveness vs readiness

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Liveness vs readiness"
```

### 2. Dependency checks and timeouts

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Dependency checks and timeouts"
```

### 3. Graceful degradation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Graceful degradation"
```

### 4. Startup probes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Startup probes"
```

### 5. What not to check

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "What not to check"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Liveness vs readiness. Predict the output before running.
2. **Change** - modify one detail in Dependency checks and timeouts. Explain the difference.
3. **Break** - break What not to check on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Liveness vs readiness and Dependency checks and timeouts in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Checking a dependency in a liveness probe. Explain the cause and the fix in your own words.
- **Mistake:** Blocking health checks on slow downstreams. Explain the cause and the fix in your own words.
- **Mistake:** Returning 200 while the service cannot serve traffic. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Liveness vs readiness and Dependency checks and timeouts. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises What not to check.
3. Document one trade-off and one failure mode of Health Checks.

## Self-check

1. Explain Liveness vs readiness without opening this note.
2. Compare Dependency checks and timeouts with an alternative. When is each appropriate?
3. What happens at the edge case of What not to check, and how do you detect it?
4. Where does Health Checks appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package log/slog](https://pkg.go.dev/log/slog)
- [ ] [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)
- [ ] [Prometheus client_golang](https://pkg.go.dev/github.com/prometheus/client_golang)
- [ ] [Google SRE book: monitoring](https://sre.google/sre-book/monitoring-distributed-systems/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
