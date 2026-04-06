# Developer Guidelines for Platoon Go

This document contains project-specific information for developers and AI agents working on the Platoon Go rewrite.

## Build and Configuration

### Requirements
- **Go 1.26.1+**: The project uses modern Go features and idioms.
- **Dependencies**: Managed via `go.mod`. Run `go mod download` to fetch them.

### Building
To build the application binary:
```shell
go build -o bin/platoon main.go
```
The binary will be placed in the `bin/` directory.

### Configuration
Platoon Go uses a `platoon.yml` file in the project root. You can generate a template using:
```shell
./bin/platoon init
```

## Testing

### Running Tests
Run all tests in the project:
```shell
go test ./...
```

Run tests for a specific package (e.g., `internal/output`):
```shell
go test ./internal/output/...
```

### Adding New Tests
- Place test files in the same directory as the code being tested, with a `_test.go` suffix.
- Use the standard `testing` package.
- **Go 1.26 Idiom**: Always use `t.Context()` when a test function needs a context.

### Example Test Case
To verify the output package, you can create a test like this:
```go
package output

import (
	"testing"
	"github.com/fatih/color"
)

func TestHighlight(t *testing.T) {
	color.NoColor = false // Force color output for testing
	text := "hello"
	expected := color.New(color.FgGreen).Sprint(text)
	result := Highlight(text)
	if result != expected {
		t.Errorf("Highlight(%q) = %q; want %q", text, result, expected)
	}
}
```

## Development Guidelines

### Modern Go Idioms (Go 1.26)
This project strictly follows modern Go practices. Key idioms to use:
- **`any`** instead of `interface{}`.
- **`omitzero`** instead of `omitempty` for JSON tags (especially for `time.Time`, `time.Duration`, and structs).
- **`new(val)`** for pointers to literals (e.g., `cfg.Timeout = new(30)`).
- **`slices` and `maps` packages** for common operations (e.g., `slices.Contains`, `maps.Keys`).
- **`for i := range n`** for simple loops from 0 to n-1.
- **`errors.AsType[T](err)`** for checking error types.
- **`wg.Go(fn)`** when using `sync.WaitGroup` to spawn goroutines.

### Project Structure
- `cmd/`: CLI command definitions using Cobra.
- `internal/`: Private library code.
    - `config/`: Configuration loading and validation.
    - `deploy/`: Core deployment orchestration.
    - `ssh/`: SSH client implementation.
    - `output/`: Terminal output styling and logging.

### Code Style
- Follow standard Go formatting (`gofmt`).
- Use descriptive variable names, but keep them concise in short scopes.
- Handle all errors explicitly. Use `errors.Is` and `errors.AsType` for error checking.
