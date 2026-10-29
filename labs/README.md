# Labs

Every module in the curriculum has a **runnable Go starter package**. You do not need to create folders or guess file names: open the module lab, read the TODO, implement it, and run the tests.

All labs live inside one Go module (`learn-golang/labs`), so commands are simple.

## Quick start

```bash
cd labs

# list the exercises in a module
go run ./beginner/03-language-basics

# run one exercise
go run ./beginner/03-language-basics variables-constants-types

# run the tests
go test ./beginner/03-language-basics -v
```

An unimplemented exercise prints `TODO` and exits with status 1. That is expected until you fill it in.

## How a module lab works

```text
labs/beginner/03-language-basics/
├── README.md                        lab map, workflow, checklist
├── main.go                          exercise runner
├── topics.go                        registry of exercises
├── 01_variables_constants_types_lab.go
├── 02_operators_and_expressions_lab.go
├── ...
└── exercises_test.go                one skipped test per exercise
```

- Each `NN_<topic>_lab.go` file contains one function with a TODO.
- `main.go` runs a single exercise by slug, or lists them.
- `exercises_test.go` starts every test with `t.Skip`. Remove the skip and add assertions when you implement the exercise.
- The module README links to every topic note in `docs/`.

## Level map

### Beginner

| Module | Command |
| --- | --- |
| 01 Introduction to Go | `go run ./beginner/01-introduction-to-go` |
| 02 Installation and Setup | `go run ./beginner/02-installation-and-setup` |
| 03 Language Basics | `go run ./beginner/03-language-basics` |
| 04 Composite Types | `go run ./beginner/04-composite-types` |
| 05 Methods and Interfaces | `go run ./beginner/05-methods-and-interfaces` |
| 06 Errors and Panics | `go run ./beginner/06-errors-and-panics` |
| 07 Packages and Modules | `go run ./beginner/07-packages-and-modules` |
| 08 Standard Library Essentials | `go run ./beginner/08-standard-library` |
| 09 Files and JSON | `go run ./beginner/09-files-and-json` |
| 10 Testing Basics | `go run ./beginner/10-testing-basics` |
| Capstone: CLI | `go run ./beginner/capstone-cli` |

### Intermediate

| Module | Command |
| --- | --- |
| 01 Concurrency Fundamentals | `go run ./intermediate/01-concurrency-fundamentals` |
| 02 Context and Cancellation | `go run ./intermediate/02-context` |
| 03 Generics | `go run ./intermediate/03-generics` |
| 04 HTTP Clients and Servers | `go run ./intermediate/04-http-clients-and-servers` |
| 05 Databases | `go run ./intermediate/05-databases` |
| 06 Project Architecture | `go run ./intermediate/06-project-architecture` |
| 07 Advanced Testing | `go run ./intermediate/07-testing-advanced` |
| 08 Observability | `go run ./intermediate/08-observability` |
| 09 Docker and Deployment | `go run ./intermediate/09-docker-and-deployment` |
| Capstone: REST API | `go run ./intermediate/capstone-rest-api/cmd/api` |

### Expert

| Module | Command |
| --- | --- |
| 01 Runtime Internals | `go run ./expert/01-runtime-internals` |
| 02 Performance Engineering | `go run ./expert/02-performance-engineering` |
| 03 Advanced Concurrency | `go run ./expert/03-advanced-concurrency` |
| 04 Distributed Systems | `go run ./expert/04-distributed-systems` |
| 05 gRPC and Event-Driven | `go run ./expert/05-grpc-and-event-driven` |
| 06 Security | `go run ./expert/06-security` |
| 07 Data at Scale | `go run ./expert/07-data-at-scale` |
| 08 Kubernetes and Cloud | `go run ./expert/08-kubernetes-and-cloud` |
| Capstone: Production Platform | `go run ./expert/capstone-platform/cmd/api` |

## Workflow for each topic

1. Open the module lab README and the linked topic note in `docs/`.
2. Open the matching `NN_<topic>_lab.go`.
3. Replace the placeholder with a small runnable experiment that covers the concepts in the file header.
4. Run the exercise and read the output.
5. Open `exercises_test.go`, remove the `t.Skip`, and add real assertions.
6. Run:

   ```bash
   gofmt -w .
   go vet ./...
   go test ./... -v
   ```

7. Mark the topic and module in `PROGRESS.md`.

## Rules

- Keep one exercise small enough for one sitting.
- Print labels so the output is understandable.
- Do not import a dependency unless the topic needs it.
- Do not silently delete TODOs; implement them or mark them blocked.
- If a module needs PostgreSQL, Redis, or a broker, start it with Docker Compose and document the command in the module README.
- Do not commit secrets, `.env` files, or generated databases.

## Capstones

The three capstone labs are full project skeletons, not single exercises:

- `beginner/capstone-cli` - commands, JSON store, tests
- `intermediate/capstone-rest-api` - HTTP API, config, store, migrations, Docker
- `expert/capstone-platform` - API and worker binaries, delivery logic, Kubernetes manifests

Read the capstone README in `docs/` first, then follow the TODO order in the lab README.
