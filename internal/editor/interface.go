package editor

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-nvim/config"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func (this *BufferStruct) RenderBuffer() {
	_, screenHeight := screen.Screen.Size()

	if this.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = this.CurrentLine
	}

	if this.CurrentLine >= config.ScrollOffSet+screenHeight-1 {
		config.ScrollOffSet = this.CurrentLine - (screenHeight - 1)
	}

	screen.Screen.Clear()

	stText := tcell.StyleDefault.Background(tcell.Color233).Foreground(tcell.Color250)
	stLineNumber := tcell.StyleDefault.Background(tcell.Color235).Foreground(tcell.Color240)

	visibleEnd := config.ScrollOffSet + screenHeight - 3

	for i := 0; i < screenHeight-3; i++ {
		lineIndex := config.ScrollOffSet + i

		if lineIndex >= len(this.Text) {
			break // Evita desenhar fora do buffer
		}

		lineNumber := fmt.Sprintf("%4d ", lineIndex+1)
		for x, char := range lineNumber {
			screen.Screen.SetContent(x, i, char, nil, stLineNumber)
		}

		for x, char := range this.Text[lineIndex] {
			screen.Screen.SetContent(x+7, i, char, nil, stText)
		}
	}

	if this.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = this.CurrentLine
	} else if this.CurrentLine >= visibleEnd {
		config.ScrollOffSet = this.CurrentLine - (screenHeight - 3) + 1
	}

	cursorScreenRow := this.CurrentLine - config.ScrollOffSet
	screen.Screen.ShowCursor(this.CurrentColumn+7, cursorScreenRow)
	screen.Screen.Show()
}

func (this *EditorStruct) CommandLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(tcell.Color233).Foreground(tcell.Color249)

	for x := 0; x < width; x++ {
		char := ' '

		if x < len(this.CurrentCommand) {
			char = this.CurrentCommand[x]
		}
		screen.Screen.SetContent(x, height-1, char, nil, bgStyle)
	}
	// Atualizar a tela
	screen.Screen.ShowCursor(this.Buffer.CurrentColumn, height-1)
	screen.Screen.Show()
}

func (this *EditorStruct) StatusLine() {
	width, height := screen.Screen.Size()

	bgStyle := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.Color249)

	nameFile := this.Currentfile

	if this.Currentfile == "" {
		nameFile = "Empty"
	}

	status := fmt.Sprintf(" %s  %s  %d/%d", this.Mode, nameFile, this.Buffer.CurrentLine+1, this.Buffer.CurrentColumn+1)

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

	indexMsg := 0

	bgStyleMsg := tcell.StyleDefault.Background(tcell.Color234).Foreground(tcell.Color208)
	for x := (width / 2) - 8; x < (width/2)+8; x++ {
		char := ' '
		if indexMsg < len(this.Message) {
			char = rune(this.Message[indexMsg])
			indexMsg++
		}
		screen.Screen.SetContent(x, height-2, char, nil, bgStyleMsg)
	}

	// Atualizar a tela
	screen.Screen.Show()
}
