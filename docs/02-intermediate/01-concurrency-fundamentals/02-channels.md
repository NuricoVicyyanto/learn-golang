# 01.2 Channels

> **Level:** Intermediate · **Module:** 01 Concurrency Fundamentals · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Channels transfer ownership and synchronize execution. Their close semantics and buffering rules are small but unforgiving.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Unbuffered vs buffered channels; Send, receive, and blocking semantics.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Unbuffered vs buffered channels

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Unbuffered vs buffered channels"
```

### 2. Send, receive, and blocking semantics

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Send, receive, and blocking semantics"
```

### 3. Closing channels and the receiver's role

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Closing channels and the receiver's role"
```

### 4. `range` over channels and `ok` on receive

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`range` over channels and `ok` on receive"
```

### 5. Directional channels and nil channels

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Directional channels and nil channels"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Unbuffered vs buffered channels. Predict the output before running.
2. **Change** - modify one detail in Send, receive, and blocking semantics. Explain the difference.
3. **Break** - break Directional channels and nil channels on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Unbuffered vs buffered channels and Send, receive, and blocking semantics in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Sending on a closed channel. Explain the cause and the fix in your own words.
- **Mistake:** Closing a channel from the receiver side. Explain the cause and the fix in your own words.
- **Mistake:** Using a nil channel and blocking forever. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Unbuffered vs buffered channels and Send, receive, and blocking semantics. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Directional channels and nil channels.
3. Document one trade-off and one failure mode of Channels.

## Self-check

1. Explain Unbuffered vs buffered channels without opening this note.
2. Compare Send, receive, and blocking semantics with an alternative. When is each appropriate?
3. What happens at the edge case of Directional channels and nil channels, and how do you detect it?
4. Where does Channels appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go blog: share memory by communicating](https://go.dev/blog/codelab-share)
- [ ] [The Go memory model](https://go.dev/ref/mem)
- [ ] [Package sync](https://pkg.go.dev/sync)
- [ ] [Go blog: advanced concurrency patterns](https://go.dev/blog/advanced-go-concurrency-patterns)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
