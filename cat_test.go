package cat_test

import (
	"context"
	"os"
	"strings"

	"github.com/yupsh/cat"
	"github.com/yupsh/cat/opt"
)

func ExampleCat() {
	ctx := context.Background()
	input := strings.NewReader("Hello World\nSecond Line\n")

	cmd := cat.Cat()
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// Hello World
	// Second Line
}

func ExampleCat_withNumbering() {
	ctx := context.Background()
	input := strings.NewReader("First line\nSecond line\nThird line\n")

	cmd := cat.Cat(opt.NumberLines)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	//      1	First line
	//      2	Second line
	//      3	Third line
}

func ExampleCat_withShowEnds() {
	ctx := context.Background()
	input := strings.NewReader("Line 1\nLine 2\n")

	cmd := cat.Cat(opt.ShowEnds)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// Line 1$
	// Line 2$
}

func ExampleCat_withSqueezeBlank() {
	ctx := context.Background()
	input := strings.NewReader("Line 1\n\n\nLine 2\n")

	cmd := cat.Cat(opt.SqueezeBlank)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// Line 1
	//
	// Line 2
}

func ExampleCat_multipleFlags() {
	ctx := context.Background()
	input := strings.NewReader("Hello\n\n\nWorld\n")

	cmd := cat.Cat(opt.NumberLines, opt.ShowEnds, opt.SqueezeBlank)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	//      1	Hello$
	//      2	$
	//      3	World$
}
