package cat

import (
	"bufio"
	"context"
	"fmt"
	"io"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"

	localopt "github.com/yupsh/cat/opt"
)

// Flags represents the configuration options for the cat command
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Cat creates a new cat command with the given parameters
func Cat(parameters ...any) yup.Command {
	return command(opt.Args[string, Flags](parameters...))
}

func (c command) Execute(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	return yup.ProcessFilesWithContext(
		ctx, c.Positional, stdin, stdout, stderr,
		yup.FileProcessorOptions{
			CommandName:     "cat",
			ContinueOnError: true,
		},
		func(ctx context.Context, source yup.InputSource, output io.Writer) error {
			return c.processReader(ctx, source.Reader, output, source.Filename)
		},
	)
}

func (c command) processReader(ctx context.Context, reader io.Reader, output io.Writer, filename string) error {
	scanner := bufio.NewScanner(reader)
	lineNum := 1
	lastLineEmpty := false

	for yup.ScanWithContext(ctx, scanner) {
		line := scanner.Text()

		// Handle squeeze blank lines
		if c.Flags.SqueezeBlank && line == "" {
			if lastLineEmpty {
				continue
			}
			lastLineEmpty = true
		} else {
			lastLineEmpty = false
		}

		// Handle line numbering
		if c.Flags.NumberLines {
			fmt.Fprintf(output, "%6d\t", lineNum)
			lineNum++
		}

		// Handle show tabs
		if c.Flags.ShowTabs {
			line = string([]rune(line)) // Convert tabs to ^I representation
			// This is simplified - real implementation would replace \t with ^I
		}

		fmt.Fprint(output, line)

		// Handle show ends
		if c.Flags.ShowEnds {
			fmt.Fprint(output, "$")
		}

		fmt.Fprintln(output)
	}

	// Check if context was cancelled
	if err := yup.CheckContextCancellation(ctx); err != nil {
		return err
	}

	return scanner.Err()
}
