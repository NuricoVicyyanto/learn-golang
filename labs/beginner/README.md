# Beginner Labs

Suggested folders for the beginner modules. Create them as you start each module.

```text
beginner/
├── 01-hello/
├── 02-tooling/
├── 03-language-basics/
├── 04-composite-types/
├── 05-interfaces/
├── 06-errors/
├── 07-packages/
├── 08-stdlib/
├── 09-files-json/
├── 10-testing/
└── capstone-cli/
```

## First exercise

```bash
cd labs
mkdir -p beginner/01-hello
```

Create `beginner/01-hello/main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run ./beginner/01-hello
```

Then change the program to print your name and the current year. Run it again.

## Rules

- Type every example; do not paste.
- Keep each exercise small enough to finish in one sitting.
- Add a test when behavior is non-trivial.
- Commit after each working exercise.
