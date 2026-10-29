# 07.3 Sharding and Replication

> **Level:** Expert · **Module:** 07 Data at Scale · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Single-node databases eventually hit limits. Replication and sharding buy capacity at the cost of consistency and operational complexity.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Read replicas and replication lag; Partition keys and hotspots.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Read replicas and replication lag

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Read replicas and replication lag"
```

### 2. Partition keys and hotspots

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Partition keys and hotspots"
```

### 3. Consistent hashing and resharding

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Consistent hashing and resharding"
```

### 4. Consistency trade-offs and read-your-writes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Consistency trade-offs and read-your-writes"
```

### 5. Operational tooling and failure modes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Operational tooling and failure modes"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Read replicas and replication lag. Predict the output before running.
2. **Change** - modify one detail in Partition keys and hotspots. Explain the difference.
3. **Break** - break Operational tooling and failure modes on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Read replicas and replication lag and Partition keys and hotspots in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Choosing a shard key with hot partitions. Explain the cause and the fix in your own words.
- **Mistake:** Reading from a replica and expecting fresh writes. Explain the cause and the fix in your own words.
- **Mistake:** Sharding before exhausting simpler options. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Read replicas and replication lag.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Operational tooling and failure modes.
3. Write a short ADR-style note explaining a trade-off in Sharding and Replication.

## Self-check

1. Explain Read replicas and replication lag without opening this note.
2. Compare Partition keys and hotspots with an alternative. When is each appropriate?
3. What happens at the edge case of Operational tooling and failure modes, and how do you detect it?
4. Where does Sharding and Replication appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Use The Index, Luke](https://use-the-index-luke.com/)
- [ ] [Redis documentation](https://redis.io/docs/latest/)
- [ ] [PostgreSQL: indexes](https://www.postgresql.org/docs/current/indexes.html)
- [ ] [Google SRE book: data integrity](https://sre.google/sre-book/data-integrity/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
