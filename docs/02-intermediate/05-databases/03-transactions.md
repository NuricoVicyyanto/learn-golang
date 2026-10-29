# 05.3 Transactions

> **Level:** Intermediate · **Module:** 05 Databases · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Transactions protect invariants across multiple statements. Isolation levels and deadlocks determine whether your logic is actually safe.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `Begin`, `Commit`, and `Rollback`; Isolation levels and anomalies.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `Begin`, `Commit`, and `Rollback`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`Begin`, `Commit`, and `Rollback`"
```

### 2. Isolation levels and anomalies

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Isolation levels and anomalies"
```

### 3. Deadlocks and retry strategy

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Deadlocks and retry strategy"
```

### 4. Context and timeouts in transactions

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Context and timeouts in transactions"
```

### 5. Repository-level transaction boundaries

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Repository-level transaction boundaries"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `Begin`, `Commit`, and `Rollback`. Predict the output before running.
2. **Change** — modify one detail in Isolation levels and anomalies. Explain the difference.
3. **Break** — break Repository-level transaction boundaries on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `Begin`, `Commit`, and `Rollback` and Isolation levels and anomalies in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Ignoring the error from `Rollback` after `Commit`. Explain the cause and the fix in your own words.
- **Mistake:** Holding a transaction open during network calls. Explain the cause and the fix in your own words.
- **Mistake:** Assuming the default isolation level prevents every anomaly. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `Begin`, `Commit`, and `Rollback` and Isolation levels and anomalies. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Repository-level transaction boundaries.
3. Document one trade-off and one failure mode of Transactions.

## Self-check

1. Explain `Begin`, `Commit`, and `Rollback` without opening this note.
2. Compare Isolation levels and anomalies with an alternative. When is each appropriate?
3. What happens at the edge case of Repository-level transaction boundaries, and how do you detect it?
4. Where does Transactions appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package database/sql](https://pkg.go.dev/database/sql)
- [ ] [Go docs: accessing databases](https://go.dev/doc/database/)
- [ ] [Go blog: database/sql](https://go.dev/blog/database-sql)
- [ ] [PostgreSQL documentation](https://www.postgresql.org/docs/)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
