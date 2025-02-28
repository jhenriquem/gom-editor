package keymap

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/event"
	"github.com/jhenriquem/Gom/internal/screen"
)

func InputInNormalMode(eventKey *tcell.EventKey) {
	switch eventKey.Rune() {
	case 'i':
		editor.GOM.CrrMode = "INSERT"
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)

	case 'h':
		editor.GOM.CrrBuffer.MoveCursor(0, -1)

	case 'l':
		editor.GOM.CrrBuffer.MoveCursor(0, 1)

	case 'k':
		editor.GOM.CrrBuffer.MoveCursor(-1, 0)

	case 'j':
		editor.GOM.CrrBuffer.MoveCursor(1, 0)

	case ':':
		event.NormalToCommand()

	}
}
