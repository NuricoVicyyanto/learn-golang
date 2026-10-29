# 05.2 Streaming RPCs

> **Level:** Expert · **Module:** 05 gRPC and Event-Driven Systems · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Streaming is powerful for large or continuous data, but flow control and cancellation become part of the design.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Server, client, and bidirectional streaming; Flow control and message sizing.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Server, client, and bidirectional streaming

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Server, client, and bidirectional streaming"
```

### 2. Flow control and message sizing

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Flow control and message sizing"
```

### 3. Deadlines and cancellation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Deadlines and cancellation"
```

### 4. Backpressure on streams

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Backpressure on streams"
```

### 5. Testing streaming handlers

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Testing streaming handlers"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Server, client, and bidirectional streaming. Predict the output before running.
2. **Change** — modify one detail in Flow control and message sizing. Explain the difference.
3. **Break** — break Testing streaming handlers on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Server, client, and bidirectional streaming and Flow control and message sizing in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Sending unbounded messages without flow control. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring `io.EOF` semantics. Explain the cause and the fix in your own words.
- **Mistake:** Leaking a stream when the client disconnects. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Server, client, and bidirectional streaming.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Testing streaming handlers.
3. Write a short ADR-style note explaining a trade-off in Streaming RPCs.

## Self-check

1. Explain Server, client, and bidirectional streaming without opening this note.
2. Compare Flow control and message sizing with an alternative. When is each appropriate?
3. What happens at the edge case of Testing streaming handlers, and how do you detect it?
4. Where does Streaming RPCs appear in a production system?

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
