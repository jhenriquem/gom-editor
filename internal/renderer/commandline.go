package renderer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/config"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func CommandLine() {
	if editor.GOM.CrrMode != "COMMAND" {
		return
	}

	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(config.ColorBgCommandLine).Foreground(config.ColorFgCommandLine)

	for x := 0; x < width; x++ {
		char := ' '

		if x < len(editor.GOM.CrrCommand) {
			char = editor.GOM.CrrCommand[x]
		}
		screen.Screen.SetContent(x, height-1, char, nil, bgStyle)
	}
	// Atualizar a tela
	screen.Screen.ShowCursor(editor.GOM.CrrBuffer.CurrentColumn, height-1)
	screen.Screen.Show()
}
