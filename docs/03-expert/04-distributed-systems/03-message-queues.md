# 04.3 Message Queues

> **Level:** Expert · **Module:** 04 Distributed Systems · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Queues decouple producers and consumers, but delivery is rarely exactly-once. You must design for duplicates, ordering, and poison messages.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: At-most-once, at-least-once, and exactly-once claims; Consumer groups and partition assignment.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. At-most-once, at-least-once, and exactly-once claims

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "At-most-once, at-least-once, and exactly-once claims"
```

### 2. Consumer groups and partition assignment

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Consumer groups and partition assignment"
```

### 3. Ordering guarantees and keys

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Ordering guarantees and keys"
```

### 4. Dead-letter queues and poison messages

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Dead-letter queues and poison messages"
```

### 5. Visibility timeouts and acknowledgements

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Visibility timeouts and acknowledgements"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for At-most-once, at-least-once, and exactly-once claims. Predict the output before running.
2. **Change** — modify one detail in Consumer groups and partition assignment. Explain the difference.
3. **Break** — break Visibility timeouts and acknowledgements on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use At-most-once, at-least-once, and exactly-once claims and Consumer groups and partition assignment in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Assuming global ordering. Explain the cause and the fix in your own words.
- **Mistake:** Retrying a poison message forever. Explain the cause and the fix in your own words.
- **Mistake:** Committing offsets before processing succeeds. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates At-most-once, at-least-once, and exactly-once claims.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Visibility timeouts and acknowledgements.
3. Write a short ADR-style note explaining a trade-off in Message Queues.

## Self-check

1. Explain At-most-once, at-least-once, and exactly-once claims without opening this note.
2. Compare Consumer groups and partition assignment with an alternative. When is each appropriate?
3. What happens at the edge case of Visibility timeouts and acknowledgements, and how do you detect it?
4. Where does Message Queues appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Microservices.io: transactional outbox](https://microservices.io/patterns/data/transactional-outbox.html)
- [ ] [Google SRE book: handling overload](https://sre.google/sre-book/handling-overload/)
- [ ] [AWS: exponential backoff and jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/)
- [ ] [Go blog: context](https://go.dev/blog/context)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
