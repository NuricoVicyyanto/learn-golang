# 05.3 Kafka or NATS

> **Level:** Expert · **Module:** 05 gRPC and Event-Driven Systems · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Event brokers differ in ordering, retention, and delivery semantics. The choice shapes your failure handling.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Topics, partitions, and consumer groups; Delivery semantics and acknowledgements.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Topics, partitions, and consumer groups

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Topics, partitions, and consumer groups"
```

### 2. Delivery semantics and acknowledgements

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Delivery semantics and acknowledgements"
```

### 3. NATS subjects and JetStream

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "NATS subjects and JetStream"
```

### 4. Schema registry and event versioning

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Schema registry and event versioning"
```

### 5. Operational concerns: lag, retention, replay

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Operational concerns: lag, retention, replay"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Topics, partitions, and consumer groups. Predict the output before running.
2. **Change** — modify one detail in Delivery semantics and acknowledgements. Explain the difference.
3. **Break** — break Operational concerns: lag, retention, replay on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Topics, partitions, and consumer groups and Delivery semantics and acknowledgements in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Choosing a broker without understanding delivery guarantees. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring consumer lag until it becomes an incident. Explain the cause and the fix in your own words.
- **Mistake:** Publishing events without a schema or owner. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Topics, partitions, and consumer groups.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Operational concerns: lag, retention, replay.
3. Write a short ADR-style note explaining a trade-off in Kafka or NATS.

## Self-check

1. Explain Topics, partitions, and consumer groups without opening this note.
2. Compare Delivery semantics and acknowledgements with an alternative. When is each appropriate?
3. What happens at the edge case of Operational concerns: lag, retention, replay, and how do you detect it?
4. Where does Kafka or NATS appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [gRPC Go documentation](https://grpc.io/docs/languages/go/)
- [ ] [Protocol Buffers language guide](https://protobuf.dev/programming-guides/proto3/)
- [ ] [buf.build documentation](https://buf.build/docs/)
- [ ] [NATS documentation](https://docs.nats.io/)
- [ ] [Kafka documentation](https://kafka.apache.org/documentation/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
