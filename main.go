package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/modes"
	"github.com/jhenriquem/go-neovim/render"
	"github.com/jhenriquem/go-neovim/screen"
)

func main() {
	log.Println("Iniciando o programa...")
	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	// Inicialização de lines
	global.Lines = append(global.Lines, []rune{})

	for global.RunningApp {
		render.RenderLines()
		render.RenderStatusLine()
		render.RenderCommandLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Screen.Sync() // Redesenhar em caso de redimensionamento

		case *tcell.EventKey:
			if modes.CurrentMODE == "NORMAL" {
				modes.KeymapsEventsForNormalMode(ev)
			} else if modes.CurrentMODE == "INSERT" {
				modes.KeymapsEventsForInsertMode(ev)
			}
		}
	}
}
