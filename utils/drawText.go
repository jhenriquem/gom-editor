package utils

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
	"github.com/jhenriquem/go-neovim/screen"
)

func DrawText() {
	screen.Screen.Clear()
	stText := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	for y, line := range global.Lines {
		lineNumber := fmt.Sprintf("%4d ", y+1)

		stLineNumber := tcell.StyleDefault.Foreground(tcell.ColorYellow)

		for x, char := range lineNumber {
			screen.Screen.SetContent(x, y, char, nil, stLineNumber)
		}

		for x, char := range line {
			screen.Screen.SetContent(x+5, y, char, nil, stText)
		}
	}

	// Exibir cursor
	screen.Screen.ShowCursor(global.CurrentColumn+5, global.CurrentLine)
	screen.Screen.Show()
}
