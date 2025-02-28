package keymap

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
)

func Input(eventKey *tcell.EventKey) {
	switch editor.GOM.CrrMode {
	case "NORMAL":
		InputInNormalMode(eventKey)
	case "INSERT":
		InputInInsertMode(eventKey)
	case "COMMAND":
		InputInCommandMode(eventKey)
	}
}
