# yup.cat

A pure Go implementation of the `cat` command that adheres to the yup.sh Command interface.

## Features

- Concatenate and display files
- Number output lines
- Show line endings
- Show tab characters
- Squeeze multiple blank lines
- Strongly typed flag system

## Usage

```go
import "github.com/yupsh/yup.cat"

// Basic usage
cmd := cat.Cat("file1.txt", "file2.txt")

// With flags
cmd := cat.Cat("file.txt", cat.NumberLines{}, cat.ShowEnds{})

// Execute
err := cmd.Execute(ctx, input, output, stderr)
```

## Flags

- `NumberLines{}` - Number all output lines (-n)
- `ShowEnds{}` - Display $ at end of each line (-E)
- `ShowTabs{}` - Display TAB characters as ^I (-T)
- `SqueezeBlank{}` - Suppress repeated empty output lines (-s)
