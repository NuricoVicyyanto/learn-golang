# Debug Log Template

> Copy this into your study log or keep a running `debug-log.md` in a lab folder. The goal is to turn bugs into reusable knowledge.

# Debug Log

## Bug - _short title_

- **Date:**
- **Module / project:**
- **Severity:** blocks work / wrong result / cosmetic
- **Time spent:**

### Symptom

<!-- Exact error message and what you observed. -->

```text
paste the error here
```

### Minimal reproduction

```go
// The smallest program that still fails.
```

### Expected vs actual

- **Expected:**
- **Actual:**

### Hypotheses

| # | Hypothesis | Test | Result |
| --- | --- | --- | --- |
| 1 |  |  |  |
| 2 |  |  |  |

### Root cause

<!-- One or two sentences. Be precise. -->

### Fix

```diff
- old code
+ new code
```

### Verification

- [ ] Regression test added
- [ ] `go test ./...` passes
- [ ] `go vet ./...` passes
- [ ] Race detector passes (if concurrency)
- [ ] Related code checked for the same bug

### Lesson

<!-- What rule will you follow next time? -->
