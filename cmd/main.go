package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/ui"
)

func main() {
	s := ui.InitScreen()

	defer s.Fini()

	for {

		// for editor.GOM.Running {
		//
		// 	renderer.Buffer()
		// 	renderer.CommandLine()
		// 	renderer.StatusLine()
		//
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEsc {
				return
			}
			// 		keymap.Input(ev)
		}
	}
}
