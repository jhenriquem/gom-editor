package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/modes"
	"github.com/jhenriquem/go-neovim/screen"
	"github.com/jhenriquem/go-neovim/utils"
)

func main() {
	log.Println("Iniciando o programa...")
	screen.ScreenInitializer()

	defer screen.Screen.Fini()

	var lines [][]rune // Armazena o texto por linha

	lines = append(lines, []rune{})

	var currentColumn, currentLine int = 0, 0

	for {
		utils.DrawText(screen.Screen, lines, &currentColumn, &currentLine)

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Screen.Sync() // Redesenhar em caso de redimensionamento

		case *tcell.EventKey:

			if modes.CurrentMODE == "NORMAL" {

				if ev.Rune() == 113 {
					return
				}
				modes.KeymapsEventsForNormalMode(ev)
			} else if modes.CurrentMODE == "INSERT" {
				modes.KeymapsEventsForInsertMode(ev, &lines, &currentColumn, &currentLine)
			}
		}
	}
}
