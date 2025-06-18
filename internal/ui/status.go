package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func Status(filename string, ismodificed bool, cod_X, cod_Y int) {
	screenWidth, _ := screen.Size()

	stText := tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite)

	modificedSign := " "
	if ismodificed {
		modificedSign = " + "
	}
	status := fmt.Sprintf("%s %s %d/%d ", filename, modificedSign, cod_X, cod_Y)

	for x, char := range status {
		screen.SetContent(screenWidth-len(status)+x, 0, char, nil, stText)
	}
}
