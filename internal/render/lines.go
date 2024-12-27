package render

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/config"
	"github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/internal/screen"
)

func RenderLines(screenHeight int) {
	screen.Screen.Clear()
	stText := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	stLineNumber := tcell.StyleDefault.Foreground(tcell.ColorYellow)

	visibleEnd := config.ScrollOffSet + screenHeight - 3

	for i := 0; i < screenHeight-3; i++ {
		lineIndex := config.ScrollOffSet + i

		if lineIndex >= len(global.Lines) {
			break // Evita desenhar fora do buffer
		}

		lineNumber := fmt.Sprintf("%4d ", lineIndex+1)
		for x, char := range lineNumber {
			screen.Screen.SetContent(x, i, char, nil, stLineNumber)
		}

		for x, char := range global.Lines[lineIndex] {
			screen.Screen.SetContent(x+5, i, char, nil, stText)
		}
	}

	// Ajustar o cursor dentro da área visível
	if global.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = global.CurrentLine
	} else if global.CurrentLine >= visibleEnd {
		config.ScrollOffSet = global.CurrentLine - (screenHeight - 3) + 1
	}

	cursorScreenRow := global.CurrentLine - config.ScrollOffSet
	screen.Screen.ShowCursor(global.CurrentColumn+5, cursorScreenRow)
	screen.Screen.Show()
}
