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
		line := editor.Editor.Buffer.Text[editor.Editor.Buffer.CurrentLine]
		if editor.Editor.Buffer.CurrentColumn < len(line) {
			editor.Editor.Buffer.Text[editor.Editor.Buffer.CurrentLine] = append(line[:editor.Editor.Buffer.CurrentColumn], line[editor.Editor.Buffer.CurrentColumn+1:]...)
		}

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		editor.Editor.Buffer.BackSpace()

	case tcell.KeyLeft:
		editor.Editor.Buffer.MoveCursor(0, -1)

	case tcell.KeyRight:
		editor.Editor.Buffer.MoveCursor(0, 1)

	case tcell.KeyUp:
		editor.Editor.Buffer.MoveCursor(-1, 0)

	case tcell.KeyDown:
		editor.Editor.Buffer.MoveCursor(1, 0)

	case tcell.KeyEnter:
		editor.Editor.Buffer.Enter()

	default:
		if eventKey.Rune() != 0 {
			editor.Editor.Buffer.Insert(eventKey.Rune())
		}
	}
}
