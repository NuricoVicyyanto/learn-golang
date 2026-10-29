# 04.3 Structs

> **Level:** Beginner · **Module:** 04 Composite Types · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Structs are Go's main way to model domain concepts. Literals, embedding, tags, and copy behavior determine how clean your APIs feel.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Defining structs and exported vs unexported fields; Composite literals, positional vs keyed fields.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Defining structs and exported vs unexported fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Defining structs and exported vs unexported fields"
```

### 2. Composite literals, positional vs keyed fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Composite literals, positional vs keyed fields"
```

### 3. Zero values and usable defaults

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Zero values and usable defaults"
```

### 4. Struct assignment copies all fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Struct assignment copies all fields"
```

### 5. Struct tags and their role in JSON, validation, and databases

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Struct tags and their role in JSON, validation, and databases"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for Defining structs and exported vs unexported fields. Predict the output before running.
2. **Change** — modify one detail in Composite literals, positional vs keyed fields. Explain the difference.
3. **Break** — break Struct tags and their role in JSON, validation, and databases on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use Defining structs and exported vs unexported fields and Composite literals, positional vs keyed fields in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using positional literals in exported APIs. Explain the cause and the fix in your own words.
- **Mistake:** Forgetting that struct assignment is a copy. Explain the cause and the fix in your own words.
- **Mistake:** Adding tags that no code actually reads. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses Defining structs and exported vs unexported fields and Composite literals, positional vs keyed fields. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Struct tags and their role in JSON, validation, and databases.
3. Explain Structs aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain Defining structs and exported vs unexported fields without opening this note.
2. Compare Composite literals, positional vs keyed fields with an alternative. When is each appropriate?
3. What happens at the edge case of Struct tags and their role in JSON, validation, and databases, and how do you detect it?
4. Where does Structs appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Go slices: usage and internals](https://go.dev/blog/slices-intro)
- [ ] [Go maps in action](https://go.dev/blog/maps)
- [ ] [Go by Example: slices, maps, structs](https://gobyexample.com/)
- [ ] [The Go specification: composite types](https://go.dev/ref/spec#Composite_types)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
