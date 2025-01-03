package keymaps

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/internal/editor"
)

func KeymapsLogicModes(eventKey *tcell.EventKey) {
	switch editor.Editor.Mode {
	case "NORMAL":
		KeymapsNormal(eventKey)
	case "INSERT":
		KeymapsInsert(eventKey)
	case "COMMAND":
		KeymapsCommand(eventKey)
	}
}
