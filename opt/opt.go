package opt

// Boolean flag types with constants
type NumberLinesFlag bool

const (
	NumberLines   NumberLinesFlag = true
	NoNumberLines NumberLinesFlag = false
)

type ShowEndsFlag bool

const (
	ShowEnds   ShowEndsFlag = true
	NoShowEnds ShowEndsFlag = false
)

type ShowTabsFlag bool

const (
	ShowTabs   ShowTabsFlag = true
	NoShowTabs ShowTabsFlag = false
)

type SqueezeBlankFlag bool

const (
	SqueezeBlank   SqueezeBlankFlag = true
	NoSqueezeBlank SqueezeBlankFlag = false
)

// Flags represents the configuration options for the cat command
type Flags struct {
	NumberLines  NumberLinesFlag  // Number all output lines
	ShowEnds     ShowEndsFlag     // Display $ at end of each line
	ShowTabs     ShowTabsFlag     // Display TAB characters as ^I
	SqueezeBlank SqueezeBlankFlag // Suppress repeated empty output lines
}

// Flag configuration methods
func (f NumberLinesFlag) Configure(flags *Flags)  { flags.NumberLines = f }
func (f ShowEndsFlag) Configure(flags *Flags)     { flags.ShowEnds = f }
func (f ShowTabsFlag) Configure(flags *Flags)     { flags.ShowTabs = f }
func (f SqueezeBlankFlag) Configure(flags *Flags) { flags.SqueezeBlank = f }
