package keymaps

import (
	"github.com/gdamore/tcell/v2"

	"github.com/jhenriquem/go-nvim/internal/editor"
)

func KeymapsNormal(eventKey *tcell.EventKey) {
	switch eventKey.Rune() {

	case 'i':
		editor.Editor.Mode = "INSERT"

	case 'h':
		editor.Editor.Buffer.MoveCursor(0, -1)

	case 'l':
		editor.Editor.Buffer.MoveCursor(0, 1)

	case 'k':
		editor.Editor.Buffer.MoveCursor(-1, 0)

	case 'j':
		editor.Editor.Buffer.MoveCursor(1, 0)

	case ':':
		editor.Editor.Buffer.CurrentColumn = 1
		editor.Editor.Mode = "COMMAND"
	}
}
