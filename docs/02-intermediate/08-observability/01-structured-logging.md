# 08.1 Structured Logging

> **Level:** Intermediate · **Module:** 08 Observability · **Status:** ⬜ Not started
> **Estimated time:** 60–75 min

## Why it matters

Structured logs are queryable. The difference between a log line and a field is the difference between grep and an investigation.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `slog` handlers and output formats; Context fields and request IDs.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `slog` handlers and output formats

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`slog` handlers and output formats"
```

### 2. Context fields and request IDs

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Context fields and request IDs"
```

### 3. Levels and sampling

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Levels and sampling"
```

### 4. Redaction and PII

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Redaction and PII"
```

### 5. Correlating logs with traces

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Correlating logs with traces"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `slog` handlers and output formats. Predict the output before running.
2. **Change** — modify one detail in Context fields and request IDs. Explain the difference.
3. **Break** — break Correlating logs with traces on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `slog` handlers and output formats and Context fields and request IDs in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Logging high-cardinality values as labels. Explain the cause and the fix in your own words.
- **Mistake:** Emitting secrets or tokens. Explain the cause and the fix in your own words.
- **Mistake:** Logging at debug level in production without sampling. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `slog` handlers and output formats and Context fields and request IDs. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Correlating logs with traces.
3. Document one trade-off and one failure mode of Structured Logging.

## Self-check

1. Explain `slog` handlers and output formats without opening this note.
2. Compare Context fields and request IDs with an alternative. When is each appropriate?
3. What happens at the edge case of Correlating logs with traces, and how do you detect it?
4. Where does Structured Logging appear in a production system?

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
