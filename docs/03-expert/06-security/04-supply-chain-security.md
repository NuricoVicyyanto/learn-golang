# 06.4 Supply Chain Security

> **Level:** Expert · **Module:** 06 Security · **Status:** ⬜ Not started
> **Estimated time:** 75–90 min

## Why it matters

Your binary includes every dependency. Knowing what is inside and what has vulnerabilities is part of shipping safely.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: `govulncheck` and dependency scanning; SBOMs and provenance.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. `govulncheck` and dependency scanning

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "`govulncheck` and dependency scanning"
```

### 2. SBOMs and provenance

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "SBOMs and provenance"
```

### 3. Signed releases and checksums

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Signed releases and checksums"
```

### 4. Reproducible builds

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Reproducible builds"
```

### 5. Dependency update policy

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Dependency update policy"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for `govulncheck` and dependency scanning. Predict the output before running.
2. **Change** - modify one detail in SBOMs and provenance. Explain the difference.
3. **Break** - break Dependency update policy on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use `govulncheck` and dependency scanning and SBOMs and provenance in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Pinning dependencies and never updating. Explain the cause and the fix in your own words.
- **Mistake:** Publishing unsigned binaries. Explain the cause and the fix in your own words.
- **Mistake:** Ignoring indirect dependency vulnerabilities. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates `govulncheck` and dependency scanning.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Dependency update policy.
3. Write a short ADR-style note explaining a trade-off in Supply Chain Security.

## Self-check

1. Explain `govulncheck` and dependency scanning without opening this note.
2. Compare SBOMs and provenance with an alternative. When is each appropriate?
3. What happens at the edge case of Dependency update policy, and how do you detect it?
4. Where does Supply Chain Security appear in a production system?

## Summary

- 
- 
- 
- 
- 

## References

- [ ] [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/)
- [ ] [Go docs: security](https://go.dev/doc/security/)
- [ ] [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
- [ ] [OWASP secrets management cheat sheet](https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html)

## Review

| Date | Recall score (1–5) | What I forgot | Next review |
| --- | --- | --- | --- |
|  |  |  |  |
