package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/config"
	"github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/internal/modes"
	"github.com/jhenriquem/go-neovim/internal/render"
	"github.com/jhenriquem/go-neovim/internal/screen"
)

func main() {
	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	// Logica para renderização das linhas perante o tamanho da janela
	_, ScreenHeight := screen.Screen.Size()

	if global.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = global.CurrentLine
	}

	if global.CurrentLine >= config.ScrollOffSet+ScreenHeight-1 {
		config.ScrollOffSet = global.CurrentLine - (ScreenHeight - 1)
	}

	// Inicialização de lines
	global.Lines = make([][]rune, 1)

	for global.RunningApp {

		if modes.CurrentMODE != "COMMAND" {
			render.RenderLines(ScreenHeight)
		} else {
			render.RenderCommandLine()
		}
		render.RenderStatusLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Screen.Sync() // Redesenhar em caso de redimensionamento

		case *tcell.EventKey:
			modes.KeymapsLogicModes(ev)
		}
	}
}
