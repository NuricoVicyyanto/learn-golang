# Exercises — Runtime Internals

> **Level:** Expert · **Rule:** attempt without notes first. Record failures honestly in [solutions.md](solutions.md).

Work in `labs/`. Keep each exercise small enough to finish in one sitting. Write tests where behavior is non-trivial.

## 1. Write a program that shows a scheduler delay using `runtime/trace`, then reduce it by changing `GOMAXPROCS` or removing a blocking call.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 2. Demonstrate a data race that the race detector catches, fix it with a channel and with a mutex, and compare the trade-offs.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 3. Measure allocations before and after removing an interface boxing or closure capture; report `-benchmem` results.

**Goal:** _state what success looks like._

**Constraints:** _what you may not use, if anything._

**Stretch:** _one extra edge case or optimization._

## 4. Rebuild from memory

Pick the topic you understood least. Close the notes and rebuild its main example from scratch. Compare with your note and record the gaps.

## 5. Test and explain

Add a table-driven test for one exercise above. Then write three sentences explaining why the test cases are sufficient — and one case you deliberately left out.

## Review

| Exercise | Attempts | Passed | Time | What I learned |
| --- | --- | --- | --- | --- |
| 1 |  |  |  |  |
| 2 |  |  |  |  |
| 3 |  |  |  |  |
| 4 |  |  |  |  |
| 5 |  |  |  |  |
