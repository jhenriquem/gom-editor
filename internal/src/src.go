package src

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/internal/editor"
	"github.com/jhenriquem/go-neovim/internal/keymaps"
	"github.com/jhenriquem/go-neovim/internal/screen"
)

func RunEditor() {
	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	for editor.Editor.Running {

		if editor.Editor.Mode != "COMMAND" {
			editor.Editor.Buffer.RenderBuffer()
		} else {
			editor.Editor.CommandLine()
		}
		editor.Editor.StatusLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Screen.Sync()

		case *tcell.EventKey:
			keymaps.KeymapsLogicModes(ev)
		}
	}
}
