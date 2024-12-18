package render

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/modes"
	"github.com/jhenriquem/go-neovim/screen"
)

func RenderStatusLine() {
	width, height := screen.Screen.Size()

	// Estilo da barra de status
	bgStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorYellow)

	// Formatar texto da barra de status
	status := fmt.Sprintf(" %s | %d/%d ", modes.CurrentMODE, global.CurrentLine+1, global.CurrentColumn+1)
	padding := width - len(status)

	if padding < 0 {
		status = status[:width-1]
		padding = 0
	}

	for x := 0; x < width; x++ {
		char := ' '
		if x < len(status) {
			char = rune(status[x])
		}
		screen.Screen.SetContent(x, height-2, char, nil, bgStyle)
	}
	// Atualizar a tela
	screen.Screen.Show()
}
