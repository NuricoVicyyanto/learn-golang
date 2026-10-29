# 05.1 database/sql

> **Level:** Intermediate · **Module:** 05 Databases · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

`database/sql` is the standard abstraction over SQL drivers. Connection pooling and `Scan` behavior are common sources of production bugs.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `sql.DB` as a pool, not a connection; `Query`, `QueryRow`, and `Exec`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `sql.DB` as a pool, not a connection

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`sql.DB` as a pool, not a connection"
```

### 2. `Query`, `QueryRow`, and `Exec`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`Query`, `QueryRow`, and `Exec`"
```

### 3. Scanning rows and handling `sql.ErrNoRows`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Scanning rows and handling `sql.ErrNoRows`"
```

### 4. Prepared statements and parameter placeholders

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Prepared statements and parameter placeholders"
```

### 5. Connection pool configuration

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Connection pool configuration"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `sql.DB` as a pool, not a connection. Predict the output before running.
2. **Change** - modify one detail in `Query`, `QueryRow`, and `Exec`. Explain the difference.
3. **Break** - break Connection pool configuration on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `sql.DB` as a pool, not a connection and `Query`, `QueryRow`, and `Exec` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Forgetting to close `Rows`. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring `rows.Err()`. Explain the cause and the fix in your own words.
- **Mistake:** Building SQL with string concatenation. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies `sql.DB` as a pool, not a connection and `Query`, `QueryRow`, and `Exec`. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Connection pool configuration.
3. Document one trade-off and one failure mode of database/sql.

## Self-check

1. Explain `sql.DB` as a pool, not a connection without opening this note.
2. Compare `Query`, `QueryRow`, and `Exec` with an alternative. When is each appropriate?
3. What happens at the edge case of Connection pool configuration, and how do you detect it?
4. Where does database/sql appear in a production system?

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
