# 08.1 Containers and Kubernetes

> **Level:** Expert · **Module:** 08 Kubernetes and Cloud · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Kubernetes schedules and restarts your process. The manifests and probes determine whether it does so correctly.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Pods, Deployments, Services, and Ingress; ConfigMaps and Secrets.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Pods, Deployments, Services, and Ingress

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pods, Deployments, Services, and Ingress"
```

### 2. ConfigMaps and Secrets

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "ConfigMaps and Secrets"
```

### 3. Liveness, readiness, and startup probes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Liveness, readiness, and startup probes"
```

### 4. Resource requests and limits

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Resource requests and limits"
```

### 5. Rollouts and rollbacks

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Rollouts and rollbacks"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Pods, Deployments, Services, and Ingress. Predict the output before running.
2. **Change** — modify one detail in ConfigMaps and Secrets. Explain the difference.
3. **Break** — break Rollouts and rollbacks on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Pods, Deployments, Services, and Ingress and ConfigMaps and Secrets in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using the same endpoint for liveness and readiness. Explain the cause and the fix in your own words.
- **Mistake:** Setting memory limits without measuring. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring termination grace periods. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Pods, Deployments, Services, and Ingress.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Rollouts and rollbacks.
3. Write a short ADR-style note explaining a trade-off in Containers and Kubernetes.

## Self-check

1. Explain Pods, Deployments, Services, and Ingress without opening this note.
2. Compare ConfigMaps and Secrets with an alternative. When is each appropriate?
3. What happens at the edge case of Rollouts and rollbacks, and how do you detect it?
4. Where does Containers and Kubernetes appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Kubernetes documentation](https://kubernetes.io/docs/home/)
- [ ] [Kubernetes: pod lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)
- [ ] [Google SRE workbook: SLOs](https://sre.google/workbook/implementing-slos/)
- [ ] [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
