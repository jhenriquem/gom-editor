package ui

import "github.com/jhenriquem/gom-editor/internal/settings"

func Cursor(cod_X, cod_Y int) {
	cursorScreenRow := cod_Y - settings.ScrollOffSet
	screen.ShowCursor(cod_X+7, cursorScreenRow)
	screen.Show()
}
