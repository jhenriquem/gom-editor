package modes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
)

func KeymapsEventsForNormalMode(eventKey *tcell.EventKey) {
	switch eventKey.Rune() {

	case 'q':
		global.RunningApp = false
	case 'i':
		CurrentMODE = "INSERT"

	case 'h':
		if global.CurrentColumn > 0 {
			global.CurrentColumn--
		} else if global.CurrentLine > 0 {
			global.CurrentLine--
			global.CurrentColumn = len((global.Lines)[global.CurrentLine])
		}

	case 'l':
		if global.CurrentColumn < len((global.Lines)[global.CurrentLine]) {
			global.CurrentColumn++
		} else if global.CurrentLine < len(global.Lines)-1 {
			global.CurrentLine++
			global.CurrentColumn = 0
		}

	case 'k':
		if global.CurrentLine > 0 {
			global.CurrentLine--
			if global.CurrentColumn > len((global.Lines)[global.CurrentLine]) {
				global.CurrentColumn = len((global.Lines)[global.CurrentLine])
			}
		}

	case 'j':
		if global.CurrentLine < len(global.Lines)-1 {
			global.CurrentLine++
			if global.CurrentColumn > len((global.Lines)[global.CurrentLine]) {
				global.CurrentColumn = len((global.Lines)[global.CurrentLine])
			}
		}

	}
}
