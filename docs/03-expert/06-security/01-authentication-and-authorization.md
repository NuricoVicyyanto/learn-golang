# 06.1 Authentication and Authorization

> **Level:** Expert · **Module:** 06 Security · **Status:** ⬜ Not started
> **Estimated time:** 90–120 min

## Why it matters

Authentication proves identity; authorization decides access. Confusing the two creates broken access control, a top API risk.

## Learning objectives

By the end of this note you can:

- [ ] Explain these concepts in your own words: Sessions vs tokens vs JWT; OAuth2 and OIDC flows.
- [ ] Apply this topic in a small program or experiment.
- [ ] Recognize the common failure modes covered in this note.
- [ ] Connect this topic to the module project.

## Core concepts

Work through each idea. Replace every prompt with your own explanation and a runnable example.

### 1. Sessions vs tokens vs JWT

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Sessions vs tokens vs JWT"
```

### 2. OAuth2 and OIDC flows

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "OAuth2 and OIDC flows"
```

### 3. Password hashing with Argon2id or bcrypt

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Password hashing with Argon2id or bcrypt"
```

### 4. RBAC and ABAC

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "RBAC and ABAC"
```

### 5. Token validation, expiry, and revocation

<!-- What is it? Why does it exist? When does it bite?
     Write 2–4 sentences in your own words, then add a minimal example. -->

**My explanation:**

**Minimal example:**

```go
// TODO: smallest example for "Token validation, expiry, and revocation"
```



## Code experiments

1. **Baseline** - reproduce the smallest working example for Sessions vs tokens vs JWT. Predict the output before running.
2. **Change** - modify one detail in OAuth2 and OIDC flows. Explain the difference.
3. **Break** - break Token validation, expiry, and revocation on purpose. Capture the compiler or runtime error and explain it.
4. **Combine** - use Sessions vs tokens vs JWT and OAuth2 and OIDC flows in one small program.

```go
// TODO: experiment space
```

## Common mistakes

- **Mistake:** Rolling your own crypto or token format. Explain the cause and the fix in your own words.
- **Mistake:** Checking authentication but not object ownership. Explain the cause and the fix in your own words.
- **Mistake:** Storing tokens in logs or URLs. Explain the cause and the fix in your own words.

## Practice

1. Build a minimal repro or benchmark that isolates Sessions vs tokens vs JWT.
2. Measure it with a profile, trace, or metrics. Record before/after numbers for Token validation, expiry, and revocation.
3. Write a short ADR-style note explaining a trade-off in Authentication and Authorization.

## Self-check

1. Explain Sessions vs tokens vs JWT without opening this note.
2. Compare OAuth2 and OIDC flows with an alternative. When is each appropriate?
3. What happens at the edge case of Token validation, expiry, and revocation, and how do you detect it?
4. Where does Authentication and Authorization appear in a production system?

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
