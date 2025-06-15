package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/core"
	"github.com/jhenriquem/gom-editor/internal/ui"
)

func main() {
	s := ui.InitScreen()

	defer s.Fini()

	core.Init()

	for core.Running {

		lines := core.Gom.Buffers[core.Gom.IndexBuffer].Lines
		coordinates := []int{core.Gom.Buffers[core.Gom.IndexBuffer].CursorX, core.Gom.Buffers[core.Gom.IndexBuffer].CursorY}

		ui.Load(lines, coordinates)

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEsc {
				core.Running = false
			}
			// 		keymap.Input(ev)
		}
	}
}
