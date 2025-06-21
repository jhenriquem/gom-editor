package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func Status(filename string, ismodificed bool, cod_X, cod_Y int) {
	screenWidth, _ := screen.Size()

	stText := tcell.StyleDefault.Background(tcell.ColorDarkGray).Foreground(tcell.ColorBlack)

	modificedSign := " "
	if ismodificed {
		modificedSign = " + "
	}

	if filename == "" {
		filename = "Empty"
	}

	status := fmt.Sprintf("  %s %s %d/%d  ", filename, modificedSign, cod_X, cod_Y)

	for x, char := range status {
		screen.SetContent(screenWidth-len(status)+x, 0, char, nil, stText)
	}

	screen.Show()
}
