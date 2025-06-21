package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/gom-editor/internal/settings"
)

func Cursor(cod_X, cod_Y int) {
	screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)

	cursorScreenRow := cod_Y - settings.ScrollOffSet

	screen.ShowCursor(cod_X+7, cursorScreenRow)
}
