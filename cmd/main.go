package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/keymap"
	"github.com/jhenriquem/Gom/internal/renderer"
	"github.com/jhenriquem/Gom/internal/screen"
)

func main() {
	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	editor.Inicializer()

	for editor.GOM.Running {

		renderer.Buffer()
		renderer.CommandLine()
		renderer.StatusLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Screen.Sync()
		case *tcell.EventKey:
			keymap.Input(ev)
		}
	}
}
