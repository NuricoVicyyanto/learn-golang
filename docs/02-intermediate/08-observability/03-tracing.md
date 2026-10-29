# 08.3 Tracing

> **Level:** Intermediate · **Module:** 08 Observability · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Traces show where time went across services. They are essential once a request touches more than one process.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Spans, traces, and context propagation; OpenTelemetry SDK and exporters.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Spans, traces, and context propagation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Spans, traces, and context propagation"
```

### 2. OpenTelemetry SDK and exporters

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "OpenTelemetry SDK and exporters"
```

### 3. Instrumenting HTTP, SQL, and custom operations

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Instrumenting HTTP, SQL, and custom operations"
```

### 4. Sampling strategies

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Sampling strategies"
```

### 5. Correlating logs and traces

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Correlating logs and traces"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Spans, traces, and context propagation. Predict the output before running.
2. **Change** - modify one detail in OpenTelemetry SDK and exporters. Explain the difference.
3. **Break** - break Correlating logs and traces on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Spans, traces, and context propagation and OpenTelemetry SDK and exporters in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Creating spans without useful attributes. Explain the cause and the fix in your own words.
- **Mistake:** Sampling everything in high-traffic systems. Explain the cause and the fix in your own words.
- **Mistake:** Breaking trace context across goroutines or queues. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Spans, traces, and context propagation and OpenTelemetry SDK and exporters. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Correlating logs and traces.
3. Document one trade-off and one failure mode of Tracing.

## Self-check

1. Explain Spans, traces, and context propagation without opening this note.
2. Compare OpenTelemetry SDK and exporters with an alternative. When is each appropriate?
3. What happens at the edge case of Correlating logs and traces, and how do you detect it?
4. Where does Tracing appear in a production system?

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
