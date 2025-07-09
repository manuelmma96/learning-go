# Learning Go - Chapter Exercises

This project contains Go programming exercises organized by chapters using Go workspaces. Each chapter contains multiple exercises that demonstrate different Go concepts.

## Project Structure

```
learning_go/
├── go.work                    # Go workspace file
├── exercises/
│   ├── ch01/                  # Chapter 1
│   │   ├── go.mod            # Module: learning_go/ch01
│   │   ├── ex01/
│   │   │   └── main.go       # Hello World exercise
│   │   └── ex02/
│   │       └── main.go       # Variables exercise
│   └── ch02/                  # Chapter 2
│       ├── go.mod            # Module: learning_go/ch02
│       ├── ex01/
│       │   └── main.go       # Type conversion exercise
│       ├── ex02/
│       │   └── main.go       # String formatting exercise
│       └── ex03/
│           └── main.go       # Math operations exercise
├── Makefile                   # Build and run targets
├── go.mod                     # Root module
└── README.md                  # This file
```

## How to Run Exercises

### Using Make Commands

The project includes a Makefile with convenient targets for running exercises:

```bash
# Show all available commands
make help

# Run all exercises from a chapter
make run-ch01
make run-ch02

# Run all exercises from all chapters
make run-all

# Run individual exercises
make run-ex01    # Run exercise 1 from chapter 1
make run-ex02    # Run exercise 2 from chapter 1

# Code quality
make fmt         # Format all Go code
make vet         # Vet all Go code
make clean       # Clean build artifacts
```

### Using Go Commands Directly

You can also run exercises directly with Go commands:

```bash
# Run a specific exercise
go run ./exercises/ch01/ex01
go run ./exercises/ch01/ex02

# Build and run
go build -o ex01 ./exercises/ch02/ex01
./ex01
```

## Adding New Exercises

To add a new exercise:

1. Create a new directory in the appropriate chapter:
   ```bash
   mkdir -p exercises/ch02/ex04
   ```

2. Create a `main.go` file in the new directory:
   ```go
   package main
   
   import "fmt"
   
   func main() {
       // Your exercise code here
       fmt.Println("Exercise 4 from Chapter 2")
   }
   ```

3. Add a new target to the Makefile:
   ```makefile
   run-ex06:
       @echo "Running Exercise 4 from Chapter 2:"
       @go run ./exercises/ch02/ex04
   ```

## Current Exercises

### Chapter 1
- **ex01**: Hello World with multiple print statements

### Chapter 2
- **ex01**: Type conversion between byte and float64

## Benefits of Go Workspace Structure

1. **Each exercise is a proper module** - can have different dependencies
2. **Natural Go commands** - `go run ./exercises/ch02/ex01` works perfectly
3. **Scalable** - easy to add new chapters and exercises
4. **Testable** - each exercise can have its own tests
5. **Modern Go** - follows current Go conventions
6. **IDE friendly** - works great with VS Code, GoLand, etc.

## Prerequisites

- Go 1.24 or later
- Make (for using the Makefile targets)

## Getting Started

1. Clone or download this repository
2. Navigate to the project directory
3. Run `go work sync` to sync the workspace
4. Run `make help` to see available commands
5. Start with `make run-ch01` to run the first chapter exercises
