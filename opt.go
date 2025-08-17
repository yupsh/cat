package command

type trimTrailingSpacesFlag bool

const (
	TrimSpaces   trimTrailingSpacesFlag = true
	NoTrimSpaces trimTrailingSpacesFlag = false
)

type numberLinesFlag bool

const (
	NumberLines   numberLinesFlag = true
	NoNumberLines numberLinesFlag = false
)

type showEndsFlag bool

const (
	ShowEnds   showEndsFlag = true
	NoShowEnds showEndsFlag = false
)

type showTabsFlag bool

const (
	ShowTabs   showTabsFlag = true
	NoShowTabs showTabsFlag = false
)

type squeezeBlankFlag bool

const (
	SqueezeBlank   squeezeBlankFlag = true
	NoSqueezeBlank squeezeBlankFlag = false
)

type flags struct {
	trimTrailingSpaces trimTrailingSpacesFlag
	numberLines        numberLinesFlag
	showEnds           showEndsFlag
	showTabs           showTabsFlag
	squeezeBlank       squeezeBlankFlag
}

func (f trimTrailingSpacesFlag) Configure(flags *flags) { flags.trimTrailingSpaces = f }
func (f numberLinesFlag) Configure(flags *flags)        { flags.numberLines = f }
func (f showEndsFlag) Configure(flags *flags)           { flags.showEnds = f }
func (f showTabsFlag) Configure(flags *flags)           { flags.showTabs = f }
func (f squeezeBlankFlag) Configure(flags *flags)       { flags.squeezeBlank = f }
