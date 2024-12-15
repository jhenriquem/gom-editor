package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
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

	crrLine := 0
	crrCol := 0

	screen.Show()

	// Função responsavel por escrevar as linhas na tela
	// Posteriormente vou remover ela daqui e criar um file só para ela (organização)
	drawScreen := func() {
		screen.Clear()
		stText := tcell.StyleDefault.Foreground(tcell.ColorWhite)

		for y, line := range lines {
			lineNumber := fmt.Sprintf("%4d ", y+1)

			stLineNumber := tcell.StyleDefault.Foreground(tcell.ColorYellow)

			for x, char := range lineNumber {
				screen.SetContent(x, y, char, nil, stLineNumber)
			}

			for x, char := range line {
				screen.SetContent(x+5, y, char, nil, stText)
			}
		}

		// Exibir cursor
		screen.ShowCursor(crrCol+5, crrLine)
		screen.Show()
	}

	for {
		drawScreen()

		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			screen.Sync() // Redesenhar em caso de redimensionamento

		case *tcell.EventKey:

			// NOTA : Vou mudar isso, o click de um tecla para sair pode ser um pouco arriscado
			// então tenho que adicionar uma logica de modos como no neovim
			if ev.Key() == tcell.KeyEscape {
				return // Sair ao pressionar ESC
			}

			// NOTA : Ter inumeras linhas relacionadas ao tratamento das teclas me incomoda e tornar
			// condigo ilegivel e extenso, então vou criar um file só para isso, organizar todos
			// esses eventos e criar ponteros para crrLine, crrCol, lines, ev.Key. Pretendo organi-
			// zar todo o projeto em files, e deixar o main só para incializar a tela e o programa

			switch ev.Key() {
			case tcell.KeyEnter:

				// Move o texto presente após o cursor para a nova linha
				currLineText := lines[crrLine][crrCol:]
				lines[crrLine] = lines[crrLine][:crrCol]

				lines = append(lines[:crrLine+1], append([][]rune{currLineText[0:]}, lines[crrLine+1:]...)...)

				crrLine++
				crrCol = 0

			case tcell.KeyBackspace, tcell.KeyBackspace2:

				if crrCol > 0 {
					crrCol--
					lines[crrLine] = append(lines[crrLine][:crrCol], lines[crrLine][crrCol+1:]...)

				} else if crrLine > 0 {

					prevLine := lines[crrLine-1]
					crrCol = len(prevLine)

					lines[crrLine-1] = append(prevLine, lines[crrLine]...)

					lines = append(lines[:crrLine], lines[crrLine+1:]...)
					crrLine--

				}

			case tcell.KeyLeft:
				// Mover cursor para a esquerda
				if crrCol > 0 {
					crrCol--
				} else if crrLine > 0 {
					crrLine--
					crrCol = len(lines[crrLine])
				}

			case tcell.KeyRight:
				// Mover cursor para a direita
				if crrCol < len(lines[crrLine]) {
					crrCol++
				} else if crrLine < len(lines)-1 {
					crrLine++
					crrCol = 0
				}

			case tcell.KeyUp:
				// Mover cursor para cima
				if crrLine > 0 {
					crrLine--
					if crrCol > len(lines[crrLine]) {
						crrCol = len(lines[crrLine])
					}
				}

			case tcell.KeyDown:
				// Mover cursor para baixo
				if crrLine < len(lines)-1 {
					crrLine++
					if crrCol > len(lines[crrLine]) {
						crrCol = len(lines[crrLine])
					}
				}

			default:
				// Adiciona as caracteres à linha
				if ev.Rune() != 0 {
					line := lines[crrLine]
					lines[crrLine] = append(line[:crrCol], append([]rune{ev.Rune()}, line[crrCol:]...)...)
					crrCol++
				}
			}
		}
	}
}
