package buffer

import (
	"github.com/jhenriquem/Gom/config"
	"github.com/jhenriquem/Gom/internal/screen"
)

func (b *BufferStruct) Enter() {
	newLineText := b.Text[b.CurrentLine][b.CurrentColumn:]

	b.Text[b.CurrentLine] = b.Text[b.CurrentLine][:b.CurrentColumn]

	newText := make([][]rune, 0, len(b.Text)+1)

	newText = append(newText, b.Text[:b.CurrentLine+1]...)

	newText = append(newText, newLineText)
	newText = append(newText, b.Text[b.CurrentLine+1:]...)

	b.Text = newText

	b.CurrentLine++
	b.CurrentColumn = 0

	_, screenHeight := screen.Screen.Size()
	if b.CurrentLine >= config.ScrollOffSet+screenHeight-config.ScrollOffNumber {
		config.ScrollOffSet++
	}
}

func (b *BufferStruct) Insert(char rune) {
	line := b.Text[b.CurrentLine]

	newLine := make([]rune, len(line)+1)

	copy(newLine, line[:b.CurrentColumn])

	newLine[b.CurrentColumn] = char
	copy(newLine[b.CurrentColumn+1:], line[b.CurrentColumn:])

	b.Text[b.CurrentLine] = newLine
	b.CurrentColumn++
}

func (b *BufferStruct) BackSpace() {
	if b.CurrentColumn > 0 {

		b.CurrentColumn--
		line := b.Text[b.CurrentLine]
		newLine := make([]rune, len(line))

		copy(newLine, line[:b.CurrentColumn])
		copy(newLine[b.CurrentColumn:], line[b.CurrentColumn+1:])

		b.Text[b.CurrentLine] = newLine

	} else if b.CurrentLine > 0 {

		prevLine := b.Text[b.CurrentLine-1]
		b.CurrentColumn = len(prevLine)

		mergedLine := make([]rune, len(prevLine)+len(b.Text[b.CurrentLine]))
		copy(mergedLine, prevLine)
		copy(mergedLine[len(prevLine):], b.Text[b.CurrentLine])
		b.Text[b.CurrentLine-1] = mergedLine

		newText := make([][]rune, len(b.Text)-1)
		copy(newText, b.Text[:b.CurrentLine])
		copy(newText[b.CurrentLine:], b.Text[b.CurrentLine+1:])

		b.Text = newText
		b.CurrentLine--

		// aplicação do scroll
		if b.CurrentLine < config.ScrollOffSet+config.ScrollOffNumber && config.ScrollOffSet >= 1 {
			config.ScrollOffSet--
		}

	}
}

func (b *BufferStruct) MoveCursor(rowDelta, colDelta int) {
	newLine := b.CurrentLine + rowDelta
	newColumn := b.CurrentColumn + colDelta

	// Garantir que o cursor não ultrapasse os limites do buffer
	if newLine < 0 {
		newLine = 0
	} else if newLine >= len(b.Text) {
		newLine = len(b.Text) - 1
	}

	if newColumn < 0 {
		if newLine > 0 {
			newLine--
			newColumn = len(b.Text[newLine])
		} else {
			newColumn = 0
		}
	} else if newColumn > len(b.Text[newLine]) {
		if newLine < len(b.Text)-1 {
			newLine++
			newColumn = 0
		} else {
			newColumn = len(b.Text[newLine])
		}
	}

	b.CurrentLine = newLine
	b.CurrentColumn = newColumn

	// Atualizar ScrollOffSet
	_, screenHeight := screen.Screen.Size()
	visibleHeight := screenHeight - 3

	if b.CurrentLine < config.ScrollOffSet {
		config.ScrollOffSet = b.CurrentLine
	} else if b.CurrentLine >= config.ScrollOffSet+visibleHeight {
		config.ScrollOffSet = b.CurrentLine - visibleHeight + 1
	}
}
