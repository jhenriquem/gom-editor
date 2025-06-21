package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/gom"
)

var screen tcell.Screen

func GetScreen() tcell.Screen {
	return screen
}

func InitScreen() tcell.Screen {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	return screen
}

func Load(buffer *gom.Buffer) {
	lines := buffer.Lines
	coordinates := []int{buffer.CursorX, buffer.CursorY}

	screen.Clear()

	Buffer(lines, coordinates[1])

	Cursor(coordinates[0], coordinates[1])

	Status(buffer.Filename, buffer.IsModified, buffer.CursorX, buffer.CursorY)
}
