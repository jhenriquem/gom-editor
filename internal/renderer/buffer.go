package renderer

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/config"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func Buffer() {
	_, screenHeight := screen.Screen.Size()

	if editor.GOM.CrrBuffer.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = editor.GOM.CrrBuffer.CurrentLine
	}

	if editor.GOM.CrrBuffer.CurrentLine >= config.ScrollOffSet+screenHeight-1 {
		config.ScrollOffSet = editor.GOM.CrrBuffer.CurrentLine - (screenHeight - 1)
	}

	screen.Screen.Clear()

	stText := tcell.StyleDefault.Background(config.ColorBgText).Foreground(config.ColorFgText)
	stLineNumber := tcell.StyleDefault.Background(config.ColorBgLineNumber).Foreground(config.ColorFgLineNumber)

	visibleEnd := config.ScrollOffSet + screenHeight - 3

	for i := 0; i < screenHeight-3; i++ {
		lineIndex := config.ScrollOffSet + i

		if lineIndex >= len(editor.GOM.CrrBuffer.Text) {
			break // Evita desenhar fora do buffer
		}

		lineNumber := fmt.Sprintf("%4d ", lineIndex+1)
		for x, char := range lineNumber {
			screen.Screen.SetContent(x, i, char, nil, stLineNumber)
		}

		for x, char := range editor.GOM.CrrBuffer.Text[lineIndex] {
			screen.Screen.SetContent(x+7, i, char, nil, stText)
		}
	}

	if editor.GOM.CrrBuffer.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = editor.GOM.CrrBuffer.CurrentLine
	} else if editor.GOM.CrrBuffer.CurrentLine >= visibleEnd {
		config.ScrollOffSet = editor.GOM.CrrBuffer.CurrentLine - (screenHeight - 3) + 1
	}

	cursorScreenRow := editor.GOM.CrrBuffer.CurrentLine - config.ScrollOffSet
	screen.Screen.ShowCursor(editor.GOM.CrrBuffer.CurrentColumn+7, cursorScreenRow)
	screen.Screen.Show()
}
