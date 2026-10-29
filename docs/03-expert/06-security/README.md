# 06 · Security

> **Level:** Expert · **Status:** ⬜ Not started · **Estimated time:** 14 hours

## Why this module

Security is a property of the whole system, not a middleware you add at the end. You will threat-model, authenticate, authorize, and harden a Go service.

## Prerequisites

- Intermediate HTTP and database modules completed.

## Learning outcomes

- [ ] Implement authentication and authorization with clear boundaries.
- [ ] Recognize and mitigate common API vulnerabilities.
- [ ] Manage secrets and TLS configuration.
- [ ] Harden the dependency and build supply chain.

## Topics

| # | Topic | Est. | Status |
| --- | --- | --- | --- |
| 06.1 | [Authentication and Authorization](01-authentication-and-authorization.md) | 90–120 min | ⬜ |
| 06.2 | [OWASP for APIs](02-owasp-for-apis.md) | 90 min | ⬜ |
| 06.3 | [Secrets and TLS](03-secrets-and-tls.md) | 75–90 min | ⬜ |
| 06.4 | [Supply Chain Security](04-supply-chain-security.md) | 75–90 min | ⬜ |

## Module checklist

- [ ] Read every topic note and explain it without notes.
- [ ] Implement every exercise in the module lab.
- [ ] Complete the [exercises](exercises.md).
- [ ] Record your solutions in [solutions.md](solutions.md).
- [ ] Build the mini project.
- [ ] Pass the self-assessment.
- [ ] Run `gofmt`, `go vet`, and `go test`.
- [ ] Update `PROGRESS.md`.

## Exercises

Do the exercises without notes first. Then compare with `solutions.md` and write down what you missed.

- [Exercises](exercises.md)
- [Solutions](solutions.md)

## Lab starter

Runnable starter code for every topic in this module:

[`labs/expert/06-security`](../../../labs/expert/06-security/README.md)

```bash
cd labs
go run ./expert/06-security
go test ./expert/06-security -v
```

## Mini project — Security Hardening Review

Take the intermediate capstone API and perform a security review: threat model, authn/authz tests, input limits, secret handling, TLS, and dependency scanning. Fix the top three findings with tests.

**Deliverables:**

- [ ] Write a threat model with trust boundaries.
- [ ] Add negative authorization tests for every resource.
- [ ] Harden headers, limits, and error responses.
- [ ] Document residual risks and accepted trade-offs.

**Definition of done:**

- [ ] Compiles and runs from a clean checkout.
- [ ] Formatted with `gofmt`; `go vet ./...` passes.
- [ ] Tests cover the core behavior and at least two edge cases.
- [ ] README explains how to run it and one design decision.
- [ ] You can explain every line without notes.

## Self-assessment

1. What is the difference between authentication and authorization?
2. Why is object-level authorization easy to miss?
3. How should a service obtain secrets in production?
4. What does `govulncheck` actually check?

Score yourself honestly. If you miss more than one question, review the module before moving on.

## References

- [ ] [OWASP API Security Top 10](https://owasp.org/API-Security/editions/2023/en/0x11-t10/)
- [ ] [Go docs: security](https://go.dev/doc/security/)
- [ ] [govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
- [ ] [OWASP secrets management cheat sheet](https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html)
