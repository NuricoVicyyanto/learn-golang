# Note Conventions

Consistency makes the repository searchable and easy to review. Follow these rules unless a module says otherwise.

---

## Language

- All documentation, comments, commit messages, and notes are in **English**.
- Code identifiers use standard Go naming (`MixedCaps`, not `snake_case`).
- Indonesian translations, if any, belong outside this repository.

## Folder and file names

- Lowercase kebab-case: `language-basics`, `table-driven-tests.md`.
- Two-digit prefixes for ordering: `01-variables.md`, `02-functions.md`.
- Module folders use two digits: `03-language-basics`.
- Capstone files are fixed: `README.md`, `requirements.md`, `milestones.md`, `rubric.md`, `retrospective.md`.
- No spaces in file names. No dates in topic names; dates belong in logs.

## Topic note structure

Every topic note starts with:

```markdown
# Topic Title

> **Level:** Beginner · **Module:** 03 Language Basics · **Status:** ⬜ Not started
> **Estimated time:** 60–90 min

## Why it matters
## Learning objectives
## Core concepts
## Code experiments
## Common mistakes
## Practice
## Self-check
## Summary
## References
```

Do not delete sections. If a section does not apply, write why.

## Status markers

Use the same markers as `PROGRESS.md`:

| Marker | Meaning |
| --- | --- |
| ⬜ | Not started |
| 🟨 | In progress |
| ✅ | Done and reviewed |
| 🔁 | Needs review |
| ⏸️ | Paused |

Change the status line when you start and finish a note.

## Code blocks

Always add a language tag:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

- Use `bash` for shell commands.
- Use `sql`, `yaml`, `json`, `proto`, `dockerfile` where relevant.
- Keep examples short enough to run.
- Mark incomplete examples with `// TODO:`.
- Do not paste huge generated files into notes; link to the lab path instead.

## Links

- Prefer relative links inside the repository.
- Link to official documentation for external references.
- Check links before marking a module done.
- When linking to a lab, use a relative path from the note: `../../../labs/...`.

## Where code lives

| Kind | Location |
| --- | --- |
| Experiments and exercises | `labs/<level>/<module>/<exercise>/` |
| Capstone | `labs/<level>/capstone-<name>/` |
| Snippets in notes | Short, illustrative only |

Each lab program should be runnable from the `labs` module:

```bash
cd labs
go run ./beginner/03-language-basics
go test ./beginner/03-language-basics -v
```

If an exercise genuinely needs its own dependencies, create a nested module with `go mod init` inside that exercise folder and document the command in its README.

## Commits

Use short imperative messages:

```text
docs(beginner): complete functions note
feat(labs): add word-count exercise
fix(beginner): correct slice aliasing example
```

One logical change per commit. Commit notes and code separately when practical.

## Markdown style

- One `#` H1 per file.
- Use `##` for main sections and `###` for subtopics.
- Keep lines readable; do not hard-wrap prose aggressively.
- Use tables for comparisons, not for long explanations.
- Use `- [ ]` checklists for actionable items.
- Use blockquotes for metadata and warnings.

## Definition of a filled note

A note is filled when:

- [ ] Every prompt has an answer in your own words.
- [ ] At least one code experiment ran.
- [ ] At least one mistake or edge case is recorded.
- [ ] The summary is five bullets or fewer.
- [ ] The self-check can be answered without opening the note.

## What not to commit

- Secrets, tokens, `.env` files with real values.
- Database files, binaries, coverage output.
- Vendor directories unless the module explicitly asks for it.
- Generated code without the command that produced it.
