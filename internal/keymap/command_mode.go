package keymap

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/event"
)

func InputInCommandMode(eventKey *tcell.EventKey) {
	switch eventKey.Key() {
	default:
		if eventKey.Rune() != 0 {
			event.InsertInCommand(eventKey)
		}
	case tcell.KeyEnter:
		event.RunCommand()
		event.EnterInCommand()
	case tcell.KeyEscape:
		event.CommandToNormal()
	case tcell.KeyLeft:
		editor.GOM.CrrBuffer.MoveCursor(0, -1)
	case tcell.KeyRight:
		editor.GOM.CrrBuffer.MoveCursor(0, 1)

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		event.BackSpaceInCommand()
	}
}
