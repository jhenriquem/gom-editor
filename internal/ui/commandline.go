package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-nvim/config"
	"github.com/jhenriquem/go-nvim/internal/editor"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func CommandLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(config.ColorBgCommandLine).Foreground(config.ColorFgCommandLine)

	for x := 0; x < width; x++ {
		char := ' '

		if x < len(editor.Editor.CurrentCommand) {
			char = editor.Editor.CurrentCommand[x]
		}
		screen.Screen.SetContent(x, height-1, char, nil, bgStyle)
	}
	// Atualizar a tela
	screen.Screen.ShowCursor(editor.Editor.Buffer.CurrentColumn, height-1)
	screen.Screen.Show()
}
