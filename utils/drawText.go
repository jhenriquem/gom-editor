package utils

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func DrawText(screen tcell.Screen, lines [][]rune, crrCol, crrLine *int) {
	screen.Clear()
	stText := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	for y, line := range lines {
		lineNumber := fmt.Sprintf("%4d ", y+1)

		stLineNumber := tcell.StyleDefault.Foreground(tcell.ColorYellow)

		for x, char := range lineNumber {
			screen.SetContent(x, y, char, nil, stLineNumber)
		}

		for x, char := range line {
			screen.SetContent(x+5, y, char, nil, stText)
		}
	}

	// Exibir cursor
	screen.ShowCursor(*crrCol+5, *crrLine)
	screen.Show()
}
