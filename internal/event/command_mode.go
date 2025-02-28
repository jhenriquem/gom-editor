package event

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/Gom/internal/editor"
	"github.com/jhenriquem/Gom/internal/screen"
	"github.com/jhenriquem/Gom/internal/utils"
)

func InsertInCommand(eventKey *tcell.EventKey) {
	editor.GOM.CrrCommand = append(
		editor.GOM.CrrCommand[:editor.GOM.CrrBuffer.CurrentColumn],
		append([]rune{eventKey.Rune()}, editor.GOM.CrrCommand[editor.GOM.CrrBuffer.CurrentColumn:]...)...,
	)
	editor.GOM.CrrBuffer.CurrentColumn++
}

func BackSpaceInCommand() {
	if len(editor.GOM.CrrCommand) > 1 && editor.GOM.CrrBuffer.CurrentColumn > 1 {
		editor.GOM.CrrBuffer.CurrentColumn--
		editor.GOM.CrrCommand = append(editor.GOM.CrrCommand[:editor.GOM.CrrBuffer.CurrentColumn], editor.GOM.CrrCommand[editor.GOM.CrrBuffer.CurrentColumn+1:]...)
	}
}

func EnterInCommand() {
	editor.GOM.CrrMode = "NORMAL"
	editor.GOM.CrrCommand = []rune{':'}
	editor.GOM.CrrBuffer.CurrentLine = 0
	screen.Screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	editor.GOM.CrrBuffer.CurrentColumn = 0
}

func RunCommand() {
	command := strings.Split(string(editor.GOM.CrrCommand), " ")

	utils.MessageCommand(command[0])

	if len(command) == 1 {
		command = append(command, "")
	}

	switch command[0] {
	case ":q":
		editor.GOM.Running = false
	case ":w":
		saveFile(command[1])
	case ":e":
		openFile(command[1])
	case ":bp":
		prevBuffer()
	case ":bn":
		nextBuffer()
	case ":bd":
		deleteBuffer()
	}
}

func saveFile(nameFile string) {
	isNew := false
	if nameFile != "" {
		editor.GOM.CrrBuffer.NameFile = nameFile
		isNew = true
	}

	// Imprimi a message de retorno do comando
	utils.MessageCommand(editor.GOM.SaveFile(isNew))
}

func openFile(nameFile string) {
	if nameFile != "" {
		file, err := os.Open(nameFile)
		if err != nil {
			file, err = os.Create(nameFile)
		}
		editor.GOM.ScanFile(file)
	}

	// Imprimi a message de retorno do comando
	utils.MessageCommand("No name file")
}

func prevBuffer() {
	var index *int = &editor.GOM.CrrBffIndex
	if *index--; *index < 0 {
		*index = len(editor.GOM.Buffers) - 1
	}
	editor.GOM.CrrBuffer = &editor.GOM.Buffers[*index]
}

func nextBuffer() {
	var index *int = &editor.GOM.CrrBffIndex
	if *index++; *index >= len(editor.GOM.Buffers) {
		*index = 0
	}
	editor.GOM.CrrBuffer = &editor.GOM.Buffers[*index]
}

func deleteBuffer() {
	var index *int = &editor.GOM.CrrBffIndex
	editor.GOM.Buffers = append(editor.GOM.Buffers[:*index], editor.GOM.Buffers[*index+1:]...)

	if *index--; *index < 0 {
		*index = 0
	}
	editor.GOM.CrrBuffer = &editor.GOM.Buffers[*index]
}
