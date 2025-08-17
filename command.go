package command

import (
	"fmt"
	"strings"

	yup `github.com/gloo-foo/framework`
)

type command yup.Inputs[yup.File, flags]

func Cat(parameters ...any) yup.Command {
	// Initialize automatically opens files when T is yup.File
	return command(yup.Initialize[yup.File, flags](parameters...))
}

func (p command) Executor() yup.CommandExecutor {
	var lastWasBlank bool

	// Wrap the helper so framework routes stdin vs files automatically
	return yup.Inputs[yup.File, flags](p).Wrap(
		yup.StatefulLineTransform(func(lineNum int64, line string) (string, bool) {
		if p.Flags.trimTrailingSpaces {
			line = strings.TrimRight(line, " \t")
		}

		if p.Flags.squeezeBlank && line == "" {
			if lastWasBlank {
				return "", false
			}
			lastWasBlank = true
		} else {
			lastWasBlank = false
		}

		if p.Flags.numberLines {
			line = fmt.Sprintf("%6d\t%s", int(lineNum), line)
		}

		if p.Flags.showTabs {
			line = strings.ReplaceAll(line, "\t", "^I")
		}

		if p.Flags.showEnds {
			line += "$"
		}

		return line, true
	}).Executor(),
	)
}
