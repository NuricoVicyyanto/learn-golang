# Exercises - Errors and Panics

> **Level:** Beginner · **Rule:** attempt without notes first. Record failures honestly in [solutions.md](solutions.md).

Work in `labs/`. Keep each exercise small enough to finish in one sitting. Write tests where behavior is non-trivial.

## 1. Build a configuration loader that returns wrapped errors with the file name, line number, and field name.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 2. Define `ValidationError` with field and reason, then use `errors.As` to collect all validation failures.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 3. Write a function that recovers from a panic at an HTTP handler boundary, logs the stack, and returns 500 without leaking internals.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 4. Rebuild from memory

Pick the topic you understood least. Close the notes and rebuild its main example from scratch. Compare with your note and record the gaps.

## 5. Test and explain

Add a table-driven test for one exercise above. Then write three sentences explaining why the test cases are sufficient - and one case you deliberately left out.

## Review

| Exercise | Attempts | Passed | Time | What I learned |
| --- | --- | --- | --- | --- |
| 1 |  |  |  |  |
| 2 |  |  |  |  |
| 3 |  |  |  |  |
| 4 |  |  |  |  |
| 5 |  |  |  |  |
