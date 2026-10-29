# 09.2 Encoding JSON

> **Level:** Beginner · **Module:** 09 Files and JSON · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

JSON is the default API format. Struct tags, `omitempty`, number handling, and unknown fields determine whether your contract is stable.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `json.Marshal` and `json.Unmarshal`; Struct tags, field names, and `omitempty`.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `json.Marshal` and `json.Unmarshal`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`json.Marshal` and `json.Unmarshal`"
```

### 2. Struct tags, field names, and `omitempty`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Struct tags, field names, and `omitempty`"
```

### 3. `json.Decoder` for streams and strict decoding

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`json.Decoder` for streams and strict decoding"
```

### 4. Custom `MarshalJSON` and `UnmarshalJSON`

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Custom `MarshalJSON` and `UnmarshalJSON`"
```

### 5. Numbers, `json.Number`, and unknown fields

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Numbers, `json.Number`, and unknown fields"
```



## Code experiments

1. **Baseline** — reproduce the smallest working example for `json.Marshal` and `json.Unmarshal`. Predict the output before running.
2. **Change** — modify one detail in Struct tags, field names, and `omitempty`. Explain the difference.
3. **Break** — break Numbers, `json.Number`, and unknown fields on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** — use `json.Marshal` and `json.Unmarshal` and Struct tags, field names, and `omitempty` in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Using `omitempty` on a field where zero is meaningful. Explain the cause and the fix in your own words.
- **Mistake:** Unmarshalling into a map and losing type safety. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring trailing data in a JSON stream. Explain the cause and the fix in your own words.

## Practice

1. Write a program that uses `json.Marshal` and `json.Unmarshal` and Struct tags, field names, and `omitempty`. Keep it under 40 lines.
2. Add a table-driven test for the edge case around Numbers, `json.Number`, and unknown fields.
3. Explain Encoding JSON aloud in three minutes, then write the summary from memory.

## Self-check

1. Explain `json.Marshal` and `json.Unmarshal` without opening this note.
2. Compare Struct tags, field names, and `omitempty` with an alternative. When is each appropriate?
3. What happens at the edge case of Numbers, `json.Number`, and unknown fields, and how do you detect it?
4. Where does Encoding JSON appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [Package os](https://pkg.go.dev/os)
- [ ] [Package encoding/json](https://pkg.go.dev/encoding/json)
- [ ] [Go blog: JSON and Go](https://go.dev/blog/json)
- [ ] [Package bufio](https://pkg.go.dev/bufio)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
