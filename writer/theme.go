package writer

import "github.com/jedib0t/go-pretty/v6/text"

const (
	themeNone    = "none"
	themeBlue    = "blue"
	themeCyan    = "cyan"
	themeGreen   = "green"
	themeMagenta = "magenta"
	themeRed     = "red"
	themeYellow  = "yellow"
)

var (
	colorsBright = text.Colors{text.FgHiBlack, text.BgHiBlue}
	colorsDark   = text.Colors{text.FgHiBlue, text.BgHiBlack}
	colorsDiff   = text.Colors{text.FgHiYellow}
	gridColorMap = []int{ // 0==dark, 1==bright
		0, 0, 0, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0,
		1, 1, 1, 0, 0, 0, 1, 1, 1,
		1, 1, 1, 0, 0, 0, 1, 1, 1,
		1, 1, 1, 0, 0, 0, 1, 1, 1,
		0, 0, 0, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0,
		0, 0, 0, 1, 1, 1, 0, 0, 0,
	}
	themes = []string{
		themeNone,
		themeBlue,
		themeCyan,
		themeGreen,
		themeMagenta,
		themeRed,
		themeYellow,
	}
	themeSelected = themeBlue
)

// SetTheme overwrites the default theme used when rendering a Grid to a string.
func SetTheme(theme string) {
	themeSelected = theme
	switch theme {
	case "blue":
		colorsBright = text.Colors{text.FgHiBlack, text.BgHiBlue}
		colorsDark = text.Colors{text.FgHiBlue, text.BgHiBlack}
		colorsDiff = text.Colors{text.FgHiYellow}
	case "cyan":
		colorsBright = text.Colors{text.FgHiBlack, text.BgHiCyan}
		colorsDark = text.Colors{text.FgHiCyan, text.BgHiBlack}
		colorsDiff = text.Colors{text.FgGreen}
	case "green":
		colorsBright = text.Colors{text.FgHiBlack, text.BgHiGreen}
		colorsDark = text.Colors{text.FgHiGreen, text.BgHiBlack}
		colorsDiff = text.Colors{text.FgHiWhite}
	case "magenta":
		colorsBright = text.Colors{text.FgHiBlack, text.BgHiMagenta}
		colorsDark = text.Colors{text.FgHiMagenta, text.BgHiBlack}
		colorsDiff = text.Colors{text.FgHiWhite}
	case "red":
		colorsBright = text.Colors{text.FgHiBlack, text.BgHiRed}
		colorsDark = text.Colors{text.FgHiRed, text.BgHiBlack}
		colorsDiff = text.Colors{text.FgHiWhite}
	case "yellow":
		colorsBright = text.Colors{text.FgHiBlack, text.BgHiYellow}
		colorsDark = text.Colors{text.FgHiYellow, text.BgHiBlack}
		colorsDiff = text.Colors{text.FgHiBlue}
	default:
		colorsBright = text.Colors{}
		colorsDark = text.Colors{}
		colorsDiff = text.Colors{text.Italic, text.Faint}
	}
}
