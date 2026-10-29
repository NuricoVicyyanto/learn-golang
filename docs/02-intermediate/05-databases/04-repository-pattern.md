# 05.4 Repository Pattern

> **Level:** Intermediate · **Module:** 05 Databases · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

A repository hides query details behind a domain-facing interface, making business logic testable and storage replaceable.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Interfaces at the consumer boundary; Mapping rows to domain types.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Interfaces at the consumer boundary

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Interfaces at the consumer boundary"
```

### 2. Mapping rows to domain types

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Mapping rows to domain types"
```

### 3. Pagination and filtering

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Pagination and filtering"
```

### 4. Transactions across repositories

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Transactions across repositories"
```

### 5. Testing with fakes and real databases

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Testing with fakes and real databases"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Interfaces at the consumer boundary. Predict the output before running.
2. **Change** — modify one detail in Mapping rows to domain types. Explain the difference.
3. **Break** — break Testing with fakes and real databases on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Interfaces at the consumer boundary and Mapping rows to domain types in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Creating one generic repository for every entity. Explain the cause and the fix in your own words.
- **Mistake:** Leaking `sql.Rows` or driver types into the domain. Explain the cause and the fix in your own words.
- **Mistake:** Over-mocking the database and never testing SQL. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Interfaces at the consumer boundary and Mapping rows to domain types. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Testing with fakes and real databases.
3. Document one trade-off and one failure mode of Repository Pattern.

## Self-check

1. Explain Interfaces at the consumer boundary without opening this note.
2. Compare Mapping rows to domain types with an alternative. When is each appropriate?
3. What happens at the edge case of Testing with fakes and real databases, and how do you detect it?
4. Where does Repository Pattern appear in a production system?

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
