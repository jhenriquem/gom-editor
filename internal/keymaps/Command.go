package keymaps

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
)

func CommandsLogic() {
	commandString := strings.Split(string(editor.Editor.CurrentCommand), " ")
	switch commandString[0] {
	case ":q":
		editor.Editor.Running = false

	case ":e":
		if len(commandString) == 2 {
			file, err := os.Open(commandString[1])
			if err != nil {
				file, err = os.Create(commandString[1])
			}
			editor.Editor.ScanFile(file)
		}

	case ":bd":
		var index *int = &editor.Editor.CrrBufferIndex
		editor.Editor.Buffers = append(editor.Editor.Buffers[:*index], editor.Editor.Buffers[*index:]...)

		if *index--; *index < 0 {
			*index = 0
		}
		editor.Editor.CrrBuffer = &editor.Editor.Buffers[*index]

	case ":bn":
		var index *int = &editor.Editor.CrrBufferIndex

		if *index++; *index >= len(editor.Editor.Buffers) {
			*index = 0
		}

		editor.Editor.CrrBuffer = &editor.Editor.Buffers[*index]

	case ":bp":
		var index *int = &editor.Editor.CrrBufferIndex

		if *index--; *index < 0 {
			*index = len(editor.Editor.Buffers) - 1
		}

		editor.Editor.CrrBuffer = &editor.Editor.Buffers[*index]

	case ":w":
		isNew := false
		if len(commandString) == 2 {
			editor.Editor.CrrBuffer.NameFile = commandString[1]
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
				editor.Editor.CurrentCommand[:editor.Editor.CrrBuffer.CurrentColumn],
				append([]rune{eventKey.Rune()}, editor.Editor.CurrentCommand[editor.Editor.CrrBuffer.CurrentColumn:]...)...,
			)
			editor.Editor.CrrBuffer.CurrentColumn++
		}
	case tcell.KeyEnter:
		CommandsLogic()
		editor.Editor.Mode = "NORMAL"
		editor.Editor.CurrentCommand = []rune{':'}
		editor.Editor.CrrBuffer.CurrentLine = 0
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		editor.Editor.CrrBuffer.CurrentColumn = 0

	case tcell.KeyEscape:
		screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		editor.Editor.Mode = "NORMAL"
		editor.Editor.CurrentCommand = []rune{':'}
		editor.Editor.CrrBuffer.CurrentLine = 0
		editor.Editor.CrrBuffer.CurrentColumn = 0

	case tcell.KeyLeft:
		editor.Editor.CrrBuffer.MoveCursor(0, -1)

	case tcell.KeyRight:
		editor.Editor.CrrBuffer.MoveCursor(0, 1)

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if len(editor.Editor.CurrentCommand) > 1 && editor.Editor.CrrBuffer.CurrentColumn > 1 {
			editor.Editor.CrrBuffer.CurrentColumn--
			editor.Editor.CurrentCommand = append(editor.Editor.CurrentCommand[:editor.Editor.CrrBuffer.CurrentColumn], editor.Editor.CurrentCommand[editor.Editor.CrrBuffer.CurrentColumn+1:]...)
		}

	}
}
