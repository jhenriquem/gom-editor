package editor

import (
	"github.com/jhenriquem/go-nvim/config"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func (this *BufferStruct) Enter() {
	newLineText := this.Text[this.CurrentLine][this.CurrentColumn:]

	this.Text[this.CurrentLine] = this.Text[this.CurrentLine][:this.CurrentColumn]

	newText := make([][]rune, 0, len(this.Text)+1)

	newText = append(newText, this.Text[:this.CurrentLine+1]...)

	newText = append(newText, newLineText)
	newText = append(newText, this.Text[this.CurrentLine+1:]...)

	this.Text = newText

	this.CurrentLine++
	this.CurrentColumn = 0

	_, screenHeight := screen.Screen.Size()
	if this.CurrentLine >= config.ScrollOffSet+screenHeight-config.ScrollOffNumber {
		config.ScrollOffSet++
	}
}

func (this *BufferStruct) Insert(char rune) {
	line := this.Text[this.CurrentLine]

	newLine := make([]rune, len(line)+1)

	copy(newLine, line[:this.CurrentColumn])

	newLine[this.CurrentColumn] = char
	copy(newLine[this.CurrentColumn+1:], line[this.CurrentColumn:])

	this.Text[this.CurrentLine] = newLine
	this.CurrentColumn++
}

func (this *BufferStruct) BackSpace() {
	if this.CurrentColumn > 0 {

		this.CurrentColumn--
		line := this.Text[this.CurrentLine]
		newLine := make([]rune, len(line))

		copy(newLine, line[:this.CurrentColumn])
		copy(newLine[this.CurrentColumn:], line[this.CurrentColumn+1:])

		this.Text[this.CurrentLine] = newLine

	} else if this.CurrentLine > 0 {

		prevLine := this.Text[this.CurrentLine-1]
		this.CurrentColumn = len(prevLine)

		mergedLine := make([]rune, len(prevLine)+len(this.Text[this.CurrentLine]))
		copy(mergedLine, prevLine)
		copy(mergedLine[len(prevLine):], this.Text[this.CurrentLine])
		this.Text[this.CurrentLine-1] = mergedLine

		newText := make([][]rune, len(this.Text)-1)
		copy(newText, this.Text[:this.CurrentLine])
		copy(newText[this.CurrentLine:], this.Text[this.CurrentLine+1:])

		this.Text = newText
		this.CurrentLine--

		// aplicação do scroll
		if this.CurrentLine < config.ScrollOffSet+config.ScrollOffNumber && config.ScrollOffSet >= 1 {
			config.ScrollOffSet--
		}

	}
}

func (this *BufferStruct) MoveCursor(rowDelta, colDelta int) {
	newLine := this.CurrentLine + rowDelta
	newColumn := this.CurrentColumn + colDelta

	// Garantir que o cursor não ultrapasse os limites do buffer
	if newLine < 0 {
		newLine = 0
	} else if newLine >= len(this.Text) {
		newLine = len(this.Text) - 1
	}

	if newColumn < 0 {
		if newLine > 0 {
			newLine--
			newColumn = len(this.Text[newLine])
		} else {
			newColumn = 0
		}
	} else if newColumn > len(this.Text[newLine]) {
		if newLine < len(this.Text)-1 {
			newLine++
			newColumn = 0
		} else {
			newColumn = len(this.Text[newLine])
		}
	}

	this.CurrentLine = newLine
	this.CurrentColumn = newColumn

	// Atualizar ScrollOffSet
	_, screenHeight := screen.Screen.Size()
	visibleHeight := screenHeight - 3

	if this.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = this.CurrentLine
	} else if this.CurrentLine >= config.ScrollOffSet+visibleHeight {
		config.ScrollOffSet = this.CurrentLine - visibleHeight + 1
	}
}
