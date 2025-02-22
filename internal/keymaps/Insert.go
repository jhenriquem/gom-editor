package keymaps

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-nvim/internal/editor"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func KeymapsInsert(eventKey *tcell.EventKey) {
	switch eventKey.Key() {

	case tcell.KeyEscape:
		editor.Editor.Mode = "NORMAL"
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)

	case tcell.KeyDelete:
		line := editor.Editor.CrrBuffer.Text[editor.Editor.CrrBuffer.CurrentLine]
		if editor.Editor.CrrBuffer.CurrentColumn < len(line) {
			editor.Editor.CrrBuffer.Text[editor.Editor.CrrBuffer.CurrentLine] = append(line[:editor.Editor.CrrBuffer.CurrentColumn], line[editor.Editor.CrrBuffer.CurrentColumn+1:]...)
		}

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		editor.Editor.CrrBuffer.BackSpace()

	case tcell.KeyLeft:
		editor.Editor.CrrBuffer.MoveCursor(0, -1)

	case tcell.KeyRight:
		editor.Editor.CrrBuffer.MoveCursor(0, 1)

	case tcell.KeyUp:
		editor.Editor.CrrBuffer.MoveCursor(-1, 0)

	case tcell.KeyDown:
		editor.Editor.CrrBuffer.MoveCursor(1, 0)

	case tcell.KeyEnter:
		editor.Editor.CrrBuffer.Enter()

	default:
		if eventKey.Rune() != 0 {
			editor.Editor.CrrBuffer.Insert(eventKey.Rune())
		}
	}
}
