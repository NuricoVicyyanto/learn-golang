# Requirements - Capstone: CLI Application

> Fill in acceptance criteria before coding. A requirement without a test is a wish.

## Scope

**In scope:**

- Three or more subcommands
- Local persistence (JSON file or SQLite)
- Input validation and actionable errors
- Human-readable output plus a `--json` mode
- Unit and command-level tests

**Out of scope:**

- Network service or multi-user sync
- Web UI
- Authentication
- Plugin system

## Functional requirements

### FR-01 - Create, list, update, and delete records from the command line

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-02 - Support filtering or search with at least two criteria

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-03 - Persist data between runs without corruption

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-04 - Provide a `--json` output mode for scripting

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-05 - Return non-zero exit codes and useful messages on failure

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

### FR-06 - Include a `--help` page for every command

- **Acceptance criteria:** _write 2–4 testable statements._
- **Priority:** must / should / could
- **Notes:**

## Non-functional requirements

| Area | Target | How I will verify it |
| --- | --- | --- |
| Performance | List 10,000 records in under 200 ms on a laptop | _test, benchmark, or manual check_ |
| Reliability | No data loss on interrupted writes; corrupt input is reported, not overwritten | _test, benchmark, or manual check_ |
| Usability | A new user can complete the main flow from `--help` alone | _test, benchmark, or manual check_ |
| Portability | Builds and runs on Windows, macOS, and Linux | _test, benchmark, or manual check_ |
| Testability | Core logic and command behavior are covered by automated tests | _test, benchmark, or manual check_ |

## Data model

| Entity | Key fields | Constraints |
| --- | --- | --- |
| Record | ID, title, status, tags, created_at, updated_at | ID is unique; status is one of a fixed set |
| Store file | schema_version, records | Must survive concurrent-safe writes within one process |

## Interface

```text
app add <title> [--tag ...]
app list [--status ...] [--tag ...] [--json]
app update <id> [--title ...] [--status ...]
app delete <id>
app stats
```

## Assumptions and open questions

1.
2.
3.
