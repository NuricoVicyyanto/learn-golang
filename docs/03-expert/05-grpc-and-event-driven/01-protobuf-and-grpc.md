# 05.1 Protobuf and gRPC

> **Level:** Expert · **Module:** 05 gRPC and Event-Driven Systems · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

gRPC gives you a typed contract, efficient transport, and generated clients. The schema is the API, so compatibility matters.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Proto3 schema design and field numbering; Code generation with `protoc` or `buf`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Proto3 schema design and field numbering

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Proto3 schema design and field numbering"
```

### 2. Code generation with `protoc` or `buf`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Code generation with `protoc` or `buf`"
```

### 3. Unary RPCs and status codes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Unary RPCs and status codes"
```

### 4. Interceptors for logging, auth, and metrics

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Interceptors for logging, auth, and metrics"
```

### 5. Schema evolution and backward compatibility

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Schema evolution and backward compatibility"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Proto3 schema design and field numbering. Predict the output before running.
2. **Change** - modify one detail in Code generation with `protoc` or `buf`. Explain the difference.
3. **Break** - break Schema evolution and backward compatibility on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Proto3 schema design and field numbering and Code generation with `protoc` or `buf` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Reusing or renumbering proto field numbers. Explain the cause and the fix in your own words.
- **Mistake:** Returning internal errors as `Unknown`. Explain the cause and the fix in your own words.
- **Mistake:** Forgetting deadlines on the client. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Proto3 schema design and field numbering.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Schema evolution and backward compatibility.
3. Write a short ADR-style note explaining a trade-off in Protobuf and gRPC.

## Self-check

1. Explain Proto3 schema design and field numbering without opening this note.
2. Compare Code generation with `protoc` or `buf` with an alternative. When is each appropriate?
3. What happens at the edge case of Schema evolution and backward compatibility, and how do you detect it?
4. Where does Protobuf and gRPC appear in a production system?

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
