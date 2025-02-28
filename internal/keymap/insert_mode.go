package keymap

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/event"
	"github.com/jhenriquem/Gom/internal/screen"
)

func InputInInsertMode(eventKey *tcell.EventKey) {
	switch eventKey.Key() {
	case tcell.KeyEscape:
		editor.GOM.CrrMode = "NORMAL"
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	case tcell.KeyLeft:
		editor.GOM.CrrBuffer.MoveCursor(0, -1)
	case tcell.KeyRight:
		editor.GOM.CrrBuffer.MoveCursor(0, 1)
	case tcell.KeyUp:
		editor.GOM.CrrBuffer.MoveCursor(-1, 0)
	case tcell.KeyDown:
		editor.GOM.CrrBuffer.MoveCursor(1, 0)
	case tcell.KeyEnter:
		editor.GOM.CrrBuffer.Enter()
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		editor.GOM.CrrBuffer.BackSpace()
	case tcell.KeyDelete:
		event.DeleteInInsertMode()
	default:
		if eventKey.Rune() != 0 {
			editor.GOM.CrrBuffer.Insert(eventKey.Rune())
		}

	}
}
