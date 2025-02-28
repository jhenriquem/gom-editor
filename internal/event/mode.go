package event

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func CommandToNormal() {
	editor.GOM.CrrMode = "NORMAL"
	editor.GOM.CrrCommand = []rune{':'}
	editor.GOM.CrrBuffer.CurrentLine = 0
	screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	editor.GOM.CrrBuffer.CurrentColumn = 0
}

func NormalToCommand() {
	screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)
	editor.GOM.CrrBuffer.CurrentColumn = 1
	editor.GOM.CrrMode = "COMMAND"
}
