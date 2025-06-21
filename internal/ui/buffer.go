package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/settings"
)

func Buffer(lines []string, cod_Y int) {
	_, screenHeight := screen.Size()

	if cod_Y < settings.ScrollOffSet {
		settings.ScrollOffSet = cod_Y
	} else if cod_Y >= settings.ScrollOffSet+screenHeight-1 {
		settings.ScrollOffSet = cod_Y - (screenHeight - 1)
	}

	stText := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)
	stLineNumber := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorBlue)

	visibleEnd := settings.ScrollOffSet + screenHeight - 2

	for i := 0; i < screenHeight; i++ {
		lineIndex := settings.ScrollOffSet + i

		if lineIndex >= len(lines) {
			break // Evita desenhar fora do buffer
		}

		lineNumber := fmt.Sprintf("%4d ", lineIndex+1)
		for x, char := range lineNumber {
			screen.SetContent(x, i, char, nil, stLineNumber)
		}

		for x, char := range lines[lineIndex] {
			screen.SetContent(x+7, i, char, nil, stText)
		}
	}

	if cod_Y < settings.ScrollOffSet {
		settings.ScrollOffSet = cod_Y
	} else if cod_Y >= visibleEnd {
		settings.ScrollOffSet = cod_Y - screenHeight - 2
	}
}
