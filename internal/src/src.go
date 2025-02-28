package src

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/keymaps"
	"github.com/jhenriquem/Gom/internal/screen"
	"github.com/jhenriquem/Gom/internal/ui"
)

func RunEditor() {
	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	for editor.Editor.Running {

		if editor.Editor.Mode != "COMMAND" {
			ui.Editor()
		} else {
			ui.CommandLine()
		}
		ui.StatusLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:

			screen.Screen.Sync()

		case *tcell.EventKey:

			keymaps.KeymapsLogicModes(ev)
		}
	}
}
