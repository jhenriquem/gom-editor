package event

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/core"
)

func KeyInput(eventKey *tcell.EventKey) {
	buffer := core.Gom.GetCurrentBuffer()

	switch eventKey.Key() {
	case tcell.KeyEscape:
		core.Running = !CloseEditor()
	case tcell.KeyLeft:
		buffer.MoveCursor(0, -1)
	case tcell.KeyRight:
		buffer.MoveCursor(0, 1)
	case tcell.KeyUp:
		buffer.MoveCursor(-1, 0)
	case tcell.KeyDown:
		buffer.MoveCursor(1, 0)
	case tcell.KeyEnter:
		buffer.EnterLine()
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		buffer.BackSpace()
	case tcell.KeyDelete:
		buffer.DeleteKey()
	default:
		if eventKey.Rune() != 0 {
			buffer.Insert(eventKey.Rune())
		}

	}

	// Verifica se o arquivo foi modificado
	ModificedFile()
}
