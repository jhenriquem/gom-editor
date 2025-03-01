package config

import "github.com/gdamore/tcell/v2"

var (
	ScrollOffNumber int = 8
	ScrollOffSet    int = 0
)

// Color configuration
var (
	ColorBg            = tcell.ColorNone
	ColorFgText        = tcell.ColorWhite
	ColorBgText        = tcell.ColorNone
	ColorFgLineNumber  = tcell.ColorGray
	ColorBgLineNumber  = tcell.ColorNone
	ColorBgStatusLine  = tcell.ColorNone
	ColorFgStatusLine  = ColorFgText
	ColorBgCommandLine = tcell.ColorNone
	ColorFgCommandLine = ColorFgText

)
