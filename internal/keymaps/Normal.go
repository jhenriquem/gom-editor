package keymaps

import (
	"github.com/gdamore/tcell/v2"

	"github.com/jhenriquem/go-neovim/internal/editor"
)

func KeymapsNormal(eventKey *tcell.EventKey) {
	switch eventKey.Rune() {

	case 'i':
		editor.Editor.Mode = "INSERT"

	case 'h':
		if editor.Editor.Buffer.CurrentColumn > 0 {
			editor.Editor.Buffer.CurrentColumn--
		} else if editor.Editor.Buffer.CurrentLine > 0 {
			editor.Editor.Buffer.CurrentLine--
			editor.Editor.Buffer.CurrentColumn = len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine])
		}

	case 'l':
		if editor.Editor.Buffer.CurrentColumn < len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
			editor.Editor.Buffer.CurrentColumn++
		} else if editor.Editor.Buffer.CurrentLine < len(editor.Editor.Buffer.Text)-1 {
			editor.Editor.Buffer.CurrentLine++
			editor.Editor.Buffer.CurrentColumn = 0
		}

	case 'k':
		if editor.Editor.Buffer.CurrentLine > 0 {
			editor.Editor.Buffer.CurrentLine--
			if editor.Editor.Buffer.CurrentColumn > len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
				editor.Editor.Buffer.CurrentColumn = len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine])
			}
		}

	case 'j':
		if editor.Editor.Buffer.CurrentLine < len(editor.Editor.Buffer.Text)-1 {
			editor.Editor.Buffer.CurrentLine++
			if editor.Editor.Buffer.CurrentColumn > len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
				editor.Editor.Buffer.CurrentColumn = len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine])
			}
		}

	case ':':
		editor.Editor.Buffer.CurrentColumn = 1
		editor.Editor.Mode = "COMMAND"
	}
}
