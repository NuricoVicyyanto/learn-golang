# Beginner Level

> **Status:** ⬜ Not started · **Estimated time:** 8–10 weeks · 80–100 hours

## Who this level is for

Developers who are new to Go or want to close gaps in the fundamentals. You can program in another language, but you do not need prior Go experience.

## Prerequisites

- Basic programming: variables, loops, functions, and arrays.
- A computer where you can install Go and run a terminal.
- Willingness to type every example instead of copying.

## Learning outcomes

- [ ] Read and write idiomatic Go for small and medium programs.
- [ ] Model data with structs, slices, maps, and interfaces.
- [ ] Handle errors deliberately and test behavior with the standard library.
- [ ] Organize code into packages and modules.
- [ ] Ship a tested command-line application.

## Modules

| # | Module | Est. |
| --- | --- | --- |
| 01 | [Introduction to Go](01-introduction-to-go/README.md) | 3 hours |
| 02 | [Installation and Setup](02-installation-and-setup/README.md) | 3 hours |
| 03 | [Language Basics](03-language-basics/README.md) | 12 hours |
| 04 | [Composite Types](04-composite-types/README.md) | 12 hours |
| 05 | [Methods and Interfaces](05-methods-and-interfaces/README.md) | 10 hours |
| 06 | [Errors and Panics](06-errors-and-panics/README.md) | 8 hours |
| 07 | [Packages and Modules](07-packages-and-modules/README.md) | 6 hours |
| 08 | [Standard Library Essentials](08-standard-library/README.md) | 10 hours |
| 09 | [Files and JSON](09-files-and-json/README.md) | 8 hours |
| 10 | [Testing Basics](10-testing-basics/README.md) | 8 hours |
| 11 | **[Capstone: CLI Application](11-capstone-cli/README.md)** | 15 hours |

## Capstone

Build a tested command-line application with persistent storage, validation, `--json` output, and a real test suite. See [11 Capstone: CLI application](11-capstone-cli/README.md).

## How to work through this level

1. Read this page and the module README before starting.
2. Follow the weekly rhythm in [`ROADMAP.md`](../../ROADMAP.md).
3. Fill every topic note in your own words.
4. Complete the exercises before the mini project.
5. Build the capstone over multiple sessions; do not wait until the end.
6. Run the exit gate before moving to the next level.

## Common pitfalls at this level

- **Copying examples without changing them.** If you never break the code, you never learn the boundary.
- **Reaching for frameworks too early.** Learn `net/http`, `database/sql`, and the standard library first.
- **Ignoring errors to make the happy path compile.** Error paths are the program.
- **Over-abstracting.** Interfaces and packages should appear when duplication or a real boundary demands them, not preemptively.
- **Skipping tests until the end.** Tests are how you know the refactor did not change behavior.
- **Comparing Go to your old language constantly.** Learn the idioms on their own terms first.

## Exit gate

Complete the capstone, pass its rubric with at least 14/18, and answer the beginner self-assessment without notes. Then run the gate review in `ROADMAP.md`.

## References

- [ ] [A Tour of Go](https://go.dev/tour/)
- [ ] [Effective Go](https://go.dev/doc/effective_go)
- [ ] [Go by Example](https://gobyexample.com/)
- [ ] [The Go specification](https://go.dev/ref/spec)
- [ ] [Go standard library](https://pkg.go.dev/std)
