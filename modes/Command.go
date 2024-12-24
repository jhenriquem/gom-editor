package modes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/global"
)

func CommandsLogic() {
	commandString := string(global.Command)

	switch commandString {
	case ":q":
		global.RunningApp = false
	}
}

func KeymapsCommand(eventKey *tcell.EventKey) {
	switch eventKey.Key() {
	default:
		if eventKey.Rune() != 0 {
			global.Command = append(
				global.Command[:global.CurrentColumn],
				append([]rune{eventKey.Rune()}, global.Command[global.CurrentColumn:]...)...,
			)
			global.CurrentColumn++
		}
	case tcell.KeyEnter:
		CommandsLogic()
	case tcell.KeyEscape:
		CurrentMODE = "NORMAL"
		global.Command = []rune{':'}
		global.CurrentLine = 0
		global.CurrentColumn = 1

	case tcell.KeyLeft:
		if global.CurrentColumn >= 1 {
			global.CurrentColumn--
		}

	case tcell.KeyRight:
		if global.CurrentColumn < len(global.Command) {
			global.CurrentColumn++
		}

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(global.Command) > 1 && global.CurrentColumn > 1 {
			global.CurrentColumn--
			global.Command = append(global.Command[:global.CurrentColumn], global.Command[global.CurrentColumn+1:]...)
		}

	}
}
