# 05.2 Migrations

> **Level:** Intermediate · **Module:** 05 Databases · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Schema changes are code deployments with state. Versioned, reversible, and reviewed migrations prevent outages.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Schema versioning and migration files; Up and down migrations.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Schema versioning and migration files

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Schema versioning and migration files"
```

### 2. Up and down migrations

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Up and down migrations"
```

### 3. Zero-downtime changes

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Zero-downtime changes"
```

### 4. Seeding and test data

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Seeding and test data"
```

### 5. Migration tooling and CI checks

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Migration tooling and CI checks"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Schema versioning and migration files. Predict the output before running.
2. **Change** - modify one detail in Up and down migrations. Explain the difference.
3. **Break** - break Migration tooling and CI checks on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Schema versioning and migration files and Up and down migrations in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Editing an already-applied migration. Explain the cause and the fix in your own words.
- **Mistake:** Dropping a column before deploying code that stopped using it. Explain the cause and the fix in your own words.
- **Mistake:** Running migrations without a backup or lock. Explain the cause and the fix in your own words.

## Practice

1. Implement a small package that applies Schema versioning and migration files and Up and down migrations. Expose a clear API and test it.
2. Write an integration-style test or benchmark that exercises Migration tooling and CI checks.
3. Document one trade-off and one failure mode of Migrations.

## Self-check

1. Explain Schema versioning and migration files without opening this note.
2. Compare Up and down migrations with an alternative. When is each appropriate?
3. What happens at the edge case of Migration tooling and CI checks, and how do you detect it?
4. Where does Migrations appear in a production system?

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
