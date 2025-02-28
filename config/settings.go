package config

import "github.com/gdamore/tcell/v2"

var (
	ScrollOffNumber int = 8
	ScrollOffSet    int = 0
)

// Color configuration
var (
	ColorBg            = tcell.ColorNone
	ColorFgText        = tcell.Color250
	ColorBgText        = tcell.ColorNone
	ColorFgLineNumber  = tcell.Color240
	ColorBgLineNumber  = tcell.ColorNone
	ColorBgStatusLine  = tcell.Color234
	ColorFgStatusLine  = tcell.Color249
	ColorBgCommandLine = tcell.ColorNone
	ColorFgCommandLine = tcell.Color249
)
