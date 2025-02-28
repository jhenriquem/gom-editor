package renderer

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/config"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func StatusLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(config.ColorBgStatusLine).Foreground(config.ColorFgStatusLine)

	nameFile := editor.GOM.CrrBuffer.NameFile

	if editor.GOM.CrrBuffer.NameFile == "" {
		nameFile = "Empty"
	}

	section_a := fmt.Sprintf(" %s  %s  %d/%d", editor.GOM.CrrMode, nameFile, editor.GOM.CrrBuffer.CurrentLine+1, editor.GOM.CrrBuffer.CurrentColumn+1)

	section_b := fmt.Sprintf("[ %d / %d ]", editor.GOM.CrrBffIndex+1, len(editor.GOM.Buffers))

	padding := width - len(section_a)

	if padding < 0 {
		section_a = section_a[:width-1]
		padding = 0
	}

	for x := 0; x < width; x++ {
		char := ' '
		if x < len(section_a) {
			char = rune(section_a[x])
		}
		screen.Screen.SetContent(x, height-2, char, nil, bgStyle)
	}

	for x := width - len(section_b); x < width; x++ {
		char := ' '
		if x < width {
			char = rune(section_b[x+len(section_b)-width])
		}
		screen.Screen.SetContent(x, height-2, char, nil, bgStyle)
	}

	// Atualizar a tela
	screen.Screen.Show()
}
