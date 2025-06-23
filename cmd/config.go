package cmd

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/core"
	"github.com/jhenriquem/gom-editor/internal/event"
	"github.com/jhenriquem/gom-editor/internal/settings"
	"github.com/jhenriquem/gom-editor/internal/ui"
)

func Execute() {
	s := ui.InitScreen()

	defer s.Fini()

	core.Init()

	if len(os.Args) > 1 {
		filename := os.Args[1]
		core.Gom.OpenFile(filename)
	}

	for core.Running {

		buffer := core.Gom.GetCurrentBuffer()
		ui.Load(buffer)

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			_, settings.ScreenHeight = s.Size()
		case *tcell.EventKey:
			event.KeyInput(ev)
		}
	}
}
