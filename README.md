# advent2025

A Go project for learning and practicing Go programming with a focus on testing.

## Project Structure

```
.
├── calculator/           # Basic arithmetic operations
│   ├── calculator.go     # Calculator functions
│   └── calculator_test.go # Unit tests, table-driven tests, and benchmarks
├── stringutil/           # String manipulation utilities
│   ├── stringutil.go     # String utility functions
│   └── stringutil_test.go # Unit tests and benchmarks
└── go.mod                # Go module definition
```

## Getting Started

### Prerequisites

- Go 1.16 or higher installed on your system
- Basic understanding of Go syntax

### Running Tests

Run all tests in the project:
```bash
go test ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

Run tests for a specific package:
```bash
go test ./calculator
go test ./stringutil
```

Run tests with coverage:
```bash
go test -cover ./...
```

Generate detailed coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Running Benchmarks

Run all benchmarks:
```bash
go test -bench=. ./...
```

Run benchmarks with memory statistics:
```bash
go test -bench=. -benchmem ./...
```

Run specific benchmark:
```bash
go test -bench=BenchmarkAdd ./calculator
```

### Building the Code

Build all packages:
```bash
go build ./...
```

Check for compilation errors:
```bash
go vet ./...
```

Format code (automatically fix formatting):
```bash
go fmt ./...
```

## Learning Resources

### Testing Patterns Demonstrated

1. **Basic Unit Tests** - Simple test functions with assertions
   - See `TestAdd`, `TestSubtract` in `calculator_test.go`

2. **Table-Driven Tests** - Go's recommended pattern for testing multiple cases
   - See `TestMultiply`, `TestDivide` in `calculator_test.go`
   - See `TestReverse`, `TestIsPalindrome` in `stringutil_test.go`

3. **Error Testing** - Testing functions that return errors
   - See `TestDivide`, `TestFactorial` in `calculator_test.go`

4. **Subtests** - Using `t.Run()` to organize related tests
   - Used throughout both test files

5. **Benchmark Tests** - Performance testing
   - See `BenchmarkAdd`, `BenchmarkMultiply` in `calculator_test.go`
   - See `BenchmarkReverse` in `stringutil_test.go`

6. **Example Tests** - Tests that appear in documentation
   - See `ExampleAdd`, `ExampleDivide` in `calculator_test.go`

### Key Go Testing Concepts

- Test files must end with `_test.go`
- Test functions must start with `Test` and accept `*testing.T`
- Benchmark functions must start with `Benchmark` and accept `*testing.B`
- Example functions must start with `Example`
- Use `t.Run()` for subtests to organize and parallelize tests
- Table-driven tests reduce code duplication and improve maintainability

## Next Steps

1. **Explore the Code** - Read through the source files to understand the functions
2. **Run the Tests** - Execute the test suite and observe the output
3. **Modify Tests** - Try adding new test cases to existing tests
4. **Write New Functions** - Add your own functions and corresponding tests
5. **Experiment with Benchmarks** - Compare performance of different implementations
6. **Add New Packages** - Create additional packages following the existing structure

## Contributing

This is a learning project. Feel free to:
- Add new packages and functionality
- Improve existing tests
- Add more comprehensive examples
- Document interesting patterns you discover

Happy learning! 🚀