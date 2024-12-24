package render

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/screen"
)

func RenderCommandLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)

	for x := 0; x < width; x++ {
		char := ' '

		if x < len(global.Command) {
			char = global.Command[x]
		}
		screen.Screen.SetContent(x, height-1, char, nil, bgStyle)
	}
	// Atualizar a tela
	screen.Screen.ShowCursor(global.CurrentColumn, height-1)
	screen.Screen.Show()
}
