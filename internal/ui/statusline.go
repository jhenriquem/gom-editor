package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-nvim/config"
	"github.com/jhenriquem/go-nvim/internal/editor"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func StatusLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(config.ColorBgStatusLine).Foreground(config.ColorFgStatusLine)

	nameFile := editor.Editor.CrrBuffer.NameFile

	if editor.Editor.CrrBuffer.NameFile == "" {
		nameFile = "Empty"
	}

	status := fmt.Sprintf(" %s  %s  %d/%d", editor.Editor.Mode, nameFile, editor.Editor.CrrBuffer.CurrentLine+1, editor.Editor.CrrBuffer.CurrentColumn+1)

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
