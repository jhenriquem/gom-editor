package editor

import (
	"github.com/jhenriquem/go-nvim/config"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func (this *BufferStruct) Enter() {
	// fmt.Printf("CurrentLine: %d, CurrentColumn: %d, BufferSize: %d\n", this.CurrentLine, this.CurrentColumn, len(this.Text))
	newLineText := this.Text[this.CurrentLine][this.CurrentColumn:]

	this.Text[this.CurrentLine] = this.Text[this.CurrentLine][:this.CurrentColumn]

	var laterLines [][]rune = [][]rune{newLineText}

	laterLines = append(laterLines, this.Text[this.CurrentLine+1:]...)

	this.Text = append(this.Text[:this.CurrentLine+1], laterLines...)

	this.CurrentLine++
	this.CurrentColumn = 0

	_, screenHeight := screen.Screen.Size()
	if this.CurrentLine >= config.ScrollOffSet+screenHeight-config.ScrollOffNumber {
		config.ScrollOffSet++
	}
}

func (this *BufferStruct) Insert(char rune) {
	// if this.CurrentLine >= len(this.Text) {
	// this.Text = append(this.Text, []rune{})
	// }

	line := this.Text[this.CurrentLine]

	// Insere o caractere na posição correta
	this.Text[this.CurrentLine] = append(line[:this.CurrentColumn], append([]rune{char}, line[this.CurrentColumn:]...)...)

	this.CurrentColumn++
}

func (this *BufferStruct) BackSpace() {
	if this.CurrentColumn > 0 {

		this.CurrentColumn--

		this.Text[this.CurrentLine] = append(this.Text[this.CurrentLine][:this.CurrentColumn], this.Text[this.CurrentLine][this.CurrentColumn+1:]...)

	} else if this.CurrentLine > 0 {

		prevLine := this.Text[this.CurrentLine-1]
		this.CurrentColumn = len(prevLine)

		this.Text[this.CurrentLine-1] = append(prevLine, this.Text[this.CurrentLine]...)

		this.Text = append(this.Text[:this.CurrentLine], (this.Text)[this.CurrentLine+1:]...)
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
