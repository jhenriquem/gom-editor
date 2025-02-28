package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/config"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
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
	screen.Screen.ShowCursor(editor.Editor.CrrBuffer.CurrentColumn, height-1)
	screen.Screen.Show()
}
