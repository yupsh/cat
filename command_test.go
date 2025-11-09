package command_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/cat"
)

// ==============================================================================
// Test Basic Functionality
// ==============================================================================

func TestCat_BasicPassThrough(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines("line1", "line2", "line3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1",
		"line2",
		"line3",
	})
}

func TestCat_EmptyInput(t *testing.T) {
	result := run.Quick(command.Cat())

	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

func TestCat_SingleLine(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines("single line").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"single line"})
}

func TestCat_EmptyLines(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines("", "", "").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"", "", ""})
}

// ==============================================================================
// Test NumberLines Flag
// ==============================================================================

func TestCat_NumberLines(t *testing.T) {
	result := run.Command(command.Cat(command.NumberLines)).
		WithStdinLines("line1", "line2", "line3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"     1\tline1",
		"     2\tline2",
		"     3\tline3",
	})
}

func TestCat_NumberLines_EmptyLines(t *testing.T) {
	result := run.Command(command.Cat(command.NumberLines)).
		WithStdinLines("line1", "", "line3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"     1\tline1",
		"     2\t",
		"     3\tline3",
	})
}

func TestCat_NumberLines_ManyLines(t *testing.T) {
	lines := make([]string, 100)
	for i := range lines {
		lines[i] = "line"
	}

	result := run.Command(command.Cat(command.NumberLines)).
		WithStdinLines(lines...).
		Run()

	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 100)
	// Check first and last
	assertion.Equal(t, result.Stdout[0], "     1\tline", "first line")
	assertion.Equal(t, result.Stdout[99], "   100\tline", "last line")
}

// ==============================================================================
// Test ShowEnds Flag
// ==============================================================================

func TestCat_ShowEnds(t *testing.T) {
	result := run.Command(command.Cat(command.ShowEnds)).
		WithStdinLines("line1", "line2", "line3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1$",
		"line2$",
		"line3$",
	})
}

func TestCat_ShowEnds_EmptyLines(t *testing.T) {
	result := run.Command(command.Cat(command.ShowEnds)).
		WithStdinLines("line1", "", "line3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1$",
		"$",
		"line3$",
	})
}

// ==============================================================================
// Test ShowTabs Flag
// ==============================================================================

func TestCat_ShowTabs(t *testing.T) {
	result := run.Command(command.Cat(command.ShowTabs)).
		WithStdinLines("col1\tcol2\tcol3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"col1^Icol2^Icol3"})
}

func TestCat_ShowTabs_NoTabs(t *testing.T) {
	result := run.Command(command.Cat(command.ShowTabs)).
		WithStdinLines("no tabs here").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"no tabs here"})
}

func TestCat_ShowTabs_MultipleTabs(t *testing.T) {
	result := run.Command(command.Cat(command.ShowTabs)).
		WithStdinLines("a\t\t\tb").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a^I^I^Ib"})
}

// ==============================================================================
// Test SqueezeBlank Flag
// ==============================================================================

func TestCat_SqueezeBlank(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("line1", "", "", "", "line2").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1",
		"",
		"line2",
	})
}

func TestCat_SqueezeBlank_SingleBlank(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("line1", "", "line2").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1",
		"",
		"line2",
	})
}

func TestCat_SqueezeBlank_NoBlankLines(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("line1", "line2", "line3").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1",
		"line2",
		"line3",
	})
}

func TestCat_SqueezeBlank_OnlyBlankLines(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("", "", "", "").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{""})
}

func TestCat_SqueezeBlank_MultipleGroups(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("a", "", "", "b", "", "", "", "c").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"a",
		"",
		"b",
		"",
		"c",
	})
}

// ==============================================================================
// Test TrimSpaces Flag
// ==============================================================================

func TestCat_TrimSpaces(t *testing.T) {
	result := run.Command(command.Cat(command.TrimSpaces)).
		WithStdinLines("line   ", "text\t\t", "end  \t  ").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line",
		"text",
		"end",
	})
}

func TestCat_TrimSpaces_NoTrailingSpaces(t *testing.T) {
	result := run.Command(command.Cat(command.TrimSpaces)).
		WithStdinLines("line1", "line2").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line1",
		"line2",
	})
}

func TestCat_TrimSpaces_LeadingSpacesPreserved(t *testing.T) {
	result := run.Command(command.Cat(command.TrimSpaces)).
		WithStdinLines("  leading", "\tleading").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"  leading",
		"\tleading",
	})
}

// ==============================================================================
// Test Flag Combinations
// ==============================================================================

func TestCat_NumberLines_ShowEnds(t *testing.T) {
	result := run.Command(command.Cat(command.NumberLines, command.ShowEnds)).
		WithStdinLines("line1", "line2").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"     1\tline1$",
		"     2\tline2$",
	})
}

func TestCat_NumberLines_ShowTabs(t *testing.T) {
	result := run.Command(command.Cat(command.NumberLines, command.ShowTabs)).
		WithStdinLines("a\tb").
		Run()

	assertion.NoError(t, result.Err)
	// ShowTabs is applied after NumberLines, so it replaces ALL tabs
	assertion.Lines(t, result.Stdout, []string{
		"     1^Ia^Ib",
	})
}

func TestCat_ShowEnds_ShowTabs(t *testing.T) {
	result := run.Command(command.Cat(command.ShowEnds, command.ShowTabs)).
		WithStdinLines("a\tb").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"a^Ib$",
	})
}

func TestCat_AllFlags(t *testing.T) {
	result := run.Command(command.Cat(
		command.NumberLines,
		command.ShowEnds,
		command.ShowTabs,
		command.SqueezeBlank,
		command.TrimSpaces,
	)).WithStdinLines(
		"line1\t  ",
		"",
		"",
		"line2\t  ",
	).Run()

	assertion.NoError(t, result.Err)
	// Order of operations:
	// 1. TrimSpaces removes trailing spaces/tabs -> "line1", "", "", "line2"
	// 2. SqueezeBlank: first "" emitted, second "" skipped
	// 3. NumberLines uses INPUT line number (not output), adds number + tab
	// 4. ShowTabs replaces ALL tabs (including the one from NumberLines)
	// 5. ShowEnds adds $
	assertion.Lines(t, result.Stdout, []string{
		"     1^Iline1$",     // Input line 1, TrimSpaces removed trailing tab
		"     2^I$",          // Input line 2, first blank line
		"     4^Iline2$",     // Input line 4, input line 3 was squeezed out
	})
}

// ==============================================================================
// Test Special Characters and Edge Cases
// ==============================================================================

func TestCat_Unicode(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines("日本語", "中文", "한국어").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"日本語",
		"中文",
		"한국어",
	})
}

func TestCat_LongLines(t *testing.T) {
	longLine := strings.Repeat("a", 10000)
	result := run.Command(command.Cat()).
		WithStdinLines(longLine).
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{longLine})
}

func TestCat_ManyLines(t *testing.T) {
	lines := make([]string, 1000)
	for i := range lines {
		lines[i] = "line"
	}

	result := run.Command(command.Cat()).
		WithStdinLines(lines...).
		Run()

	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1000)
}

func TestCat_SpecialCharacters(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines("!@#$%^&*()").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"!@#$%^&*()"})
}

func TestCat_MixedContent(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines(
			"normal line",
			"",
			"line with\ttabs",
			"日本語",
			"trailing spaces   ",
		).Run()

	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 5)
}

// ==============================================================================
// Test Error Handling
// ==============================================================================

func TestCat_InputError(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinError(errors.New("read failed")).
		Run()

	assertion.ErrorContains(t, result.Err, "read failed")
}

func TestCat_OutputError(t *testing.T) {
	result := run.Command(command.Cat()).
		WithStdinLines("test").
		WithStdoutError(errors.New("write failed")).
		Run()

	assertion.ErrorContains(t, result.Err, "write failed")
}

// ==============================================================================
// Table-Driven Tests
// ==============================================================================

func TestCat_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		flags    []any
		input    []string
		expected []string
	}{
		{
			name:     "basic pass-through",
			flags:    []any{},
			input:    []string{"a", "b"},
			expected: []string{"a", "b"},
		},
		{
			name:     "number lines",
			flags:    []any{command.NumberLines},
			input:    []string{"a", "b"},
			expected: []string{"     1\ta", "     2\tb"},
		},
		{
			name:     "show ends",
			flags:    []any{command.ShowEnds},
			input:    []string{"a", "b"},
			expected: []string{"a$", "b$"},
		},
		{
			name:     "show tabs",
			flags:    []any{command.ShowTabs},
			input:    []string{"a\tb"},
			expected: []string{"a^Ib"},
		},
		{
			name:     "squeeze blank",
			flags:    []any{command.SqueezeBlank},
			input:    []string{"a", "", "", "b"},
			expected: []string{"a", "", "b"},
		},
		{
			name:     "trim spaces",
			flags:    []any{command.TrimSpaces},
			input:    []string{"text  ", "more\t\t"},
			expected: []string{"text", "more"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := run.Command(command.Cat(tt.flags...)).
				WithStdinLines(tt.input...).
				Run()

			assertion.NoError(t, result.Err)
			assertion.Lines(t, result.Stdout, tt.expected)
		})
	}
}

// ==============================================================================
// Edge Cases for SqueezeBlank
// ==============================================================================

func TestCat_SqueezeBlank_StartWithBlanks(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("", "", "line").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"",
		"line",
	})
}

func TestCat_SqueezeBlank_EndWithBlanks(t *testing.T) {
	result := run.Command(command.Cat(command.SqueezeBlank)).
		WithStdinLines("line", "", "").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{
		"line",
		"",
	})
}

// ==============================================================================
// Edge Cases for NumberLines
// ==============================================================================

func TestCat_NumberLines_Format(t *testing.T) {
	// Verify the exact format: 6-digit right-aligned number + tab
	result := run.Command(command.Cat(command.NumberLines)).
		WithStdinLines("test").
		Run()

	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
	// Should be exactly "     1\ttest" (5 spaces, 1, tab, text)
	assertion.Equal(t, result.Stdout[0], "     1\ttest", "line format")
}

