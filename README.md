# Learn Go - From Beginner to Expert

A structured, documentation-first Go curriculum. Every module tells you **what to learn**, **what to build**, and **how to prove you are done**. Most topic notes are intentionally skeletons: you complete them in your own words, because writing and explaining are part of learning.

> **Status:** framework ready. Fill the notes level by level and tick the boxes in [`PROGRESS.md`](PROGRESS.md).

---

## Who this is for

- Developers who want a complete, ordered path through Go.
- Self-taught learners who need structure instead of a pile of tutorials.
- Engineers moving from another language who want to avoid "tourist Go".
- Anyone preparing for backend, platform, or cloud-native work with Go.

## How the repository is organized

```text
learn-golang/
├── README.md                  You are here
├── ROADMAP.md                 Level map, schedule, and exit criteria
├── PROGRESS.md                Single place to track every module
├── docs/
│   ├── 00-guides/             How to study, templates, conventions
│   ├── 01-beginner/          Language fundamentals and first CLI
│   ├── 02-intermediate/      Concurrency, HTTP, databases, deployment
│   ├── 03-expert/            Runtime, performance, distributed systems
│   └── 04-reference/         Glossary, cheatsheets, resources
└── labs/                      Go module for all hands-on code
```

Each module folder follows the same shape:

```text
NN-module-name/
├── README.md            Module map, outcomes, checklist, mini project
├── 01-topic.md          One focused topic note (fill this in)
├── 02-topic.md
├── ...
├── exercises.md         Practice tasks
└── solutions.md         Your own solutions and reflections
```

Capstone folders add `requirements.md`, `milestones.md`, `rubric.md`, and `retrospective.md`. Modules 01–02 are reading modules; their exercises live in the module README.

## What is already written vs. what you complete

| Already provided | Yours to complete |
| --- | --- |
| Roadmap, schedules, and level gates | Your own schedule in `PROGRESS.md` |
| Study guides and templates | Your own study log |
| Modules 01–02 (full lessons) | Exercises and solution logs |
| Module guides, topic maps, and prompts | Topic notes in your own words |
| Runnable lab skeletons for every topic | Implement each TODO and its test |
| Capstone requirements and rubrics | Capstone implementation and retrospective |
| Glossary, cheatsheet, stdlib map, resources | Corrections and additions based on experience |

The skeleton is intentional: 95 topic notes give you a place to think, not a transcript to copy.

## The three levels

| Level | Focus | Typical duration | You can... |
| --- | --- | --- | --- |
| [Beginner](docs/01-beginner/README.md) | Syntax, types, functions, interfaces, errors, packages, testing, files | 8–10 weeks · 80–100 h | Write idiomatic small programs and a tested CLI tool |
| [Intermediate](docs/02-intermediate/README.md) | Concurrency, context, generics, HTTP APIs, SQL, architecture, observability, Docker | 10–14 weeks · 120–160 h | Design, test, ship, and operate a production-style REST API |
| [Expert](docs/03-expert/README.md) | Runtime internals, performance, advanced concurrency, distributed systems, gRPC, security, Kubernetes | 12–16 weeks · 160–220 h | Diagnose performance issues and design reliable distributed services |

See [`ROADMAP.md`](ROADMAP.md) for the full sequence and gate checks.

## Quick start

1. Read [`ROADMAP.md`](ROADMAP.md) and pick a schedule (10 h or 20 h per week).
2. Read [`docs/00-guides/how-to-study.md`](docs/00-guides/how-to-study.md).
3. Open [`PROGRESS.md`](PROGRESS.md) and mark module 1 as in progress.
4. Run the starter lab for the module. The code and file names already exist:

   ```bash
   cd labs
   go run ./beginner/01-introduction-to-go
   go test ./beginner/01-introduction-to-go -v
   ```

5. Study the module README, fill each topic note, implement the matching `NN_<topic>_lab.go`, and do the exercises.
6. Only move on when the module checklist and the level gate are complete.

## How to use the notes

Every topic note has the same sections:

- **Why it matters** - connect the concept to real code.
- **Learning objectives** - observable things you must be able to do.
- **Core concepts** - explain each idea in your own words, with a small example.
- **Code experiments** - run, break, and fix the code.
- **Common mistakes** - record mistakes you actually made.
- **Practice** - small tasks that force recall.
- **Self-check** - questions you must answer without notes.
- **Summary** - five bullets maximum.
- **References** - official docs first.

A note is done when you can explain it to another developer and your code compiles and passes tests.

## Study rules

1. **Type it, do not copy it.** Copying code creates familiarity, not skill.
2. **One concept, one experiment.** Change something and observe what breaks.
3. **Tests are part of the exercise.** If it is worth writing, it is worth testing.
4. **Keep a debug log.** Record the bug, the cause, and the fix.
5. **Build something each level.** A capstone turns isolated knowledge into engineering judgment.
6. **Review before advancing.** Re-explain the previous module from memory.
7. **Do not skip the beginner level.** Advanced Go is mostly fundamentals under pressure.

## Status legend

| Symbol | Meaning |
| --- | --- |
| ⬜ | Not started |
| 🟨 | In progress |
| ✅ | Done and reviewed |
| 🔁 | Needs review |
| ⏸️ | Paused / blocked |

## Progress tracking

Use [`PROGRESS.md`](PROGRESS.md) as the source of truth. The per-module checkboxes are duplicated in each module README for convenience, but the root file is the dashboard.

## Conventions

- Documentation language: **English**.
- File names: lowercase kebab-case, two-digit numeric prefixes.
- Code lives in [`labs/`](labs/README.md), never inside the docs tree.
- Use fenced code blocks with a language tag (`go`, `bash`, `sql`, `yaml`).
- Link to official documentation whenever possible.
- Do not commit secrets, tokens, databases, or build artifacts.

See [`docs/00-guides/note-conventions.md`](docs/00-guides/note-conventions.md) for details.

## Definition of done

A module is complete when all of the following are true:

- [ ] Every topic note is filled in your own words.
- [ ] Every matching `labs/.../NN_<topic>_lab.go` is implemented.
- [ ] All exercises are attempted and recorded.
- [ ] The mini project compiles, is formatted, and has tests.
- [ ] You passed the self-assessment without notes.
- [ ] The module checklist and [`PROGRESS.md`](PROGRESS.md) are updated.

Full rubric: [`docs/00-guides/definition-of-done.md`](docs/00-guides/definition-of-done.md).

## Reference

- [Glossary](docs/04-reference/glossary.md)
- [Go command cheatsheet](docs/04-reference/go-commands-cheatsheet.md)
- [Standard library map](docs/04-reference/standard-library-map.md)
- [Learning resources](docs/04-reference/learning-resources.md)
- [FAQ](docs/04-reference/faq.md)

## License and use

This is a personal learning repository. Use the structure freely, but write your own notes and solutions.
