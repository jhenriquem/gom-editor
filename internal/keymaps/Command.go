package keymaps

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/internal/editor"
)

func CommandsLogic() {
	commandString := strings.Split(string(editor.Editor.CurrentCommand), " ")
	switch commandString[0] {
	case ":q":
		editor.Editor.Running = false

	case ":open":
		if len(commandString) == 2 {
			file, err := os.Open(commandString[1])
			if err != nil {
				file, err = os.Create(commandString[1])
			}
			editor.Editor.Currentfile = commandString[1]
			editor.Editor.ScanFile(file)
		}

	case ":w":
		isNew := false
		if len(commandString) == 2 {
			editor.Editor.Currentfile = commandString[1]
			isNew = true
		}

		editor.Editor.SaveFile(isNew)
	}
}

func KeymapsCommand(eventKey *tcell.EventKey) {
	switch eventKey.Key() {
	default:
		if eventKey.Rune() != 0 {
			editor.Editor.CurrentCommand = append(
				editor.Editor.CurrentCommand[:editor.Editor.Buffer.CurrentColumn],
				append([]rune{eventKey.Rune()}, editor.Editor.CurrentCommand[editor.Editor.Buffer.CurrentColumn:]...)...,
			)
			editor.Editor.Buffer.CurrentColumn++
		}
	case tcell.KeyEnter:
		CommandsLogic()
		editor.Editor.Mode = "NORMAL"
	case tcell.KeyEscape:
		editor.Editor.Mode = "NORMAL"
		editor.Editor.CurrentCommand = []rune{':'}
		editor.Editor.Buffer.CurrentLine = 0
		editor.Editor.Buffer.CurrentColumn = 1

	case tcell.KeyLeft:
		if editor.Editor.Buffer.CurrentColumn >= 1 {
			editor.Editor.Buffer.CurrentColumn--
		}

	case tcell.KeyRight:
		if editor.Editor.Buffer.CurrentColumn < len(editor.Editor.CurrentCommand) {
			editor.Editor.Buffer.CurrentColumn++
		}

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(editor.Editor.CurrentCommand) > 1 && editor.Editor.Buffer.CurrentColumn > 1 {
			editor.Editor.Buffer.CurrentColumn--
			editor.Editor.CurrentCommand = append(editor.Editor.CurrentCommand[:editor.Editor.Buffer.CurrentColumn], editor.Editor.CurrentCommand[editor.Editor.Buffer.CurrentColumn+1:]...)
		}

	}
}
