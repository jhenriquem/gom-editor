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
