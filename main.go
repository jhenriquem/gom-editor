package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/utils"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Erro ao iniciar a tela: %v", err)
	}

	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalf("Erro ao inicializar a tela: %v", err)
	}

	stScreen := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	screen.SetStyle(stScreen)

	screen.Clear()

	var lines [][]rune // Armazena o texto por linha

	lines = append(lines, []rune{})

	var currentColumn, currentLine int = 0, 0

	var currentMODE string = "NORMAL"

	screen.Show()

	for {
		utils.DrawText(screen, lines, &currentColumn, &currentLine)

		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync() // Redesenhar em caso de redimensionamento

		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return
			}
			if currentMODE == "INSERT" {
				utils.KeymapsEvents(ev, &lines, &currentColumn, &currentLine)
			}
		}
	}
}
