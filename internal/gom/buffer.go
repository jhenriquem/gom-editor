package gom

import (
	"github.com/jhenriquem/gom-editor/internal/settings"
)

type Buffer struct {
	Lines      []string
	CursorX    int
	CursorY    int
	Filename   string
	IsModified bool
}

func (b *Buffer) Insert(char rune) {
	prevText := b.Lines[b.CursorY][:b.CursorX]
	nextText := b.Lines[b.CursorY][b.CursorX:]

	prevText += string(char)

	b.Lines[b.CursorY] = prevText + nextText

	b.CursorX++
}

func (b *Buffer) EnterLine() {
	newLine := b.Lines[b.CursorY][b.CursorX:]
	originalLine := b.Lines[b.CursorY][:b.CursorX]

	nextLines := make([]string, len(b.Lines[b.CursorY+1:]))
	copy(nextLines, b.Lines[b.CursorY+1:])

	b.Lines[b.CursorY] = originalLine
	b.Lines = append(b.Lines[:b.CursorY+1], append([]string{newLine}, nextLines...)...)

	b.CursorY++
	b.CursorX = 0

	screenHeight := settings.ScreenHeight
	if b.CursorY >= settings.ScrollOffSet+screenHeight-settings.ScrollOffNumber {
		settings.ScrollOffSet++
	}
}

func (b *Buffer) DeleteKey() {
	line := b.Lines[b.CursorY]
	if b.CursorX < len(line) {
		b.Lines[b.CursorY] = line[:b.CursorX]
		b.Lines[b.CursorY] += line[b.CursorX+1:]

	} else if b.CursorX == len(line) && len(b.Lines) > 1 && b.CursorY < len(b.Lines)-1 {

		prevLine := b.Lines[:b.CursorY]
		nextLines := b.Lines[b.CursorY+1:]

		b.Lines = prevLine
		b.Lines = append(b.Lines, nextLines...)
	}
}

func (b *Buffer) BackSpace() {
	if b.CursorX > 0 {

		b.CursorX--
		line := b.Lines[b.CursorY]
		var newLine string = ""

		newLine += line[:b.CursorX]
		newLine += line[b.CursorX+1:]

		b.Lines[b.CursorY] = newLine

	} else if b.CursorY > 0 {

		prevLine := b.Lines[b.CursorY-1]
		b.CursorX = len(prevLine)

		var mergedLine string = ""
		mergedLine += prevLine
		mergedLine += b.Lines[b.CursorY]
		b.Lines[b.CursorY-1] = mergedLine

		newText := make([]string, len(b.Lines)-1)
		copy(newText, b.Lines[:b.CursorY])
		copy(newText[b.CursorY:], b.Lines[b.CursorY+1:])

		b.Lines = newText
		b.CursorY--

		// aplicação do scroll
		if b.CursorY < settings.ScrollOffSet+settings.ScrollOffNumber && settings.ScrollOffSet >= 1 {
			settings.ScrollOffSet--
		}

	}
}

func (b *Buffer) MoveCursor(rowDelta, colDelta int) {
	newY := b.CursorY + rowDelta
	newX := b.CursorX + colDelta

	// Garante que o cursor não ultrapasse o tamnho da tela e nem a quantidade de linhas
	if newY < 0 {
		newY = 0
	} else if newY >= len(b.Lines) {
		newY = len(b.Lines) - 1
	}

	if newX < 0 {
		if newY > 0 {
			newY--
			newX = len(b.Lines[newY])
		} else {
			newX = 0
		}
	} else if newX > len(b.Lines[newY]) {
		if newY < len(b.Lines)-1 {
			newY++
			newX = 0
		} else {
			newX = len(b.Lines[newY])
		}
	}

	b.CursorY = newY
	b.CursorX = newX

	// Atualizar ScrollOffSet
	visibleHeight := settings.ScreenHeight - 3

	if b.CursorY < settings.ScrollOffSet {
		settings.ScrollOffSet = b.CursorY
	} else if b.CursorY >= settings.ScrollOffSet+visibleHeight {
		settings.ScrollOffSet = b.CursorY - visibleHeight + 1
	}
}
