# yup.cat Examples

This directory contains runnable Go examples that demonstrate how to use the `yup.cat` command.

## Running Examples

The examples are implemented as Go example functions in the test files. You can run them using:

```bash
# Run all examples
go test -v

# Run specific examples
go test -run ExampleCat
go test -run ExampleCat_withNumbering

# View examples in documentation
go doc -all
```

## Available Examples

See `../cat_test.go` for the actual runnable examples:

- `ExampleCat()` - Basic file concatenation
- `ExampleCat_withNumbering()` - Add line numbers
- `ExampleCat_withShowEnds()` - Show line endings with $
- `ExampleCat_withSqueezeBlank()` - Compress multiple blank lines
- `ExampleCat_multipleFlags()` - Combine multiple formatting options

## Usage Patterns

### Standalone Usage
```go
import "github.com/nicerobot/yup.cat"
import "github.com/nicerobot/yup.cat/flags"

cmd := cat.Cat("file.txt", flags.NumberLines, flags.ShowEnds)
err := cmd.Execute(ctx, input, output, stderr)
```

### Pipeline Usage
```go
import "github.com/nicerobot/yup.sh"

pipeline := yup.Pipe(
    cat.Cat("data.txt", flags.NumberLines),
    // ... other commands
)
err := pipeline.Execute(ctx, input, output, stderr)
```

## Testing

All examples are automatically tested when you run `go test` to ensure they produce the expected output.