package keymaps

import (
	"github.com/gdamore/tcell/v2"

	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func KeymapsNormal(eventKey *tcell.EventKey) {
	switch eventKey.Rune() {
	case 'i':
		editor.Editor.Mode = "INSERT"
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)

	case 'h':
		editor.Editor.CrrBuffer.MoveCursor(0, -1)

	case 'l':
		editor.Editor.CrrBuffer.MoveCursor(0, 1)

	case 'k':
		editor.Editor.CrrBuffer.MoveCursor(-1, 0)

	case 'j':
		editor.Editor.CrrBuffer.MoveCursor(1, 0)

	case ':':
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)
		editor.Editor.CrrBuffer.CurrentColumn = 1
		editor.Editor.Mode = "COMMAND"
	}
}
