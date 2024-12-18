package modes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
)

func KeymapsEventsForInsertMode(eventKey *tcell.EventKey) {
	switch eventKey.Key() {

	case tcell.KeyEscape:
		CurrentMODE = "NORMAL"

	case tcell.KeyDelete:
		if global.CurrentColumn < len((global.Lines)[global.CurrentLine]) {
			(global.Lines)[global.CurrentLine] = append((global.Lines)[global.CurrentLine][:global.CurrentColumn], (global.Lines)[global.CurrentLine][global.CurrentColumn+1:]...)
		}

	case tcell.KeyEnter:
		currLineText := (global.Lines)[global.CurrentLine][global.CurrentColumn:]
		(global.Lines)[global.CurrentLine] = (global.Lines)[global.CurrentLine][:global.CurrentColumn]

		newLines := append([][]rune{}, (global.Lines)[:global.CurrentLine+1]...)
		newLines = append(newLines, currLineText)
		global.Lines = append(newLines, (global.Lines)[global.CurrentLine+1:]...)

		global.CurrentLine++
		global.CurrentColumn = 0

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if global.CurrentColumn > 0 {
			global.CurrentColumn--
			(global.Lines)[global.CurrentLine] = append((global.Lines)[global.CurrentLine][:global.CurrentColumn], (global.Lines)[global.CurrentLine][global.CurrentColumn+1:]...)
		} else if global.CurrentLine > 0 {
			prevLine := (global.Lines)[global.CurrentLine-1]
			global.CurrentColumn = len(prevLine)
			(global.Lines)[global.CurrentLine-1] = append(prevLine, (global.Lines)[global.CurrentLine]...)
			global.Lines = append((global.Lines)[:global.CurrentLine], (global.Lines)[global.CurrentLine+1:]...)
			global.CurrentLine--
		}

	case tcell.KeyLeft:
		if global.CurrentColumn > 0 {
			global.CurrentColumn--
		} else if global.CurrentLine > 0 {
			global.CurrentLine--
			global.CurrentColumn = len((global.Lines)[global.CurrentLine])
		}

	case tcell.KeyRight:
		if global.CurrentColumn < len((global.Lines)[global.CurrentLine]) {
			global.CurrentColumn++
		} else if global.CurrentLine < len(global.Lines)-1 {
			global.CurrentLine++
			global.CurrentColumn = 0
		}

	case tcell.KeyUp:
		if global.CurrentLine > 0 {
			global.CurrentLine--
			if global.CurrentColumn > len((global.Lines)[global.CurrentLine]) {
				global.CurrentColumn = len((global.Lines)[global.CurrentLine])
			}
		}

	case tcell.KeyDown:
		if global.CurrentLine < len(global.Lines)-1 {
			global.CurrentLine++
			if global.CurrentColumn > len((global.Lines)[global.CurrentLine]) {
				global.CurrentColumn = len((global.Lines)[global.CurrentLine])
			}
		}

	default:
		if eventKey.Rune() != 0 {
			line := (global.Lines)[global.CurrentLine]
			(global.Lines)[global.CurrentLine] = append(line[:global.CurrentColumn], append([]rune{eventKey.Rune()}, line[global.CurrentColumn:]...)...)
			global.CurrentColumn++
		}
	}
}
