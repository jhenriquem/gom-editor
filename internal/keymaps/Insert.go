package keymaps

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-neovim/config"
	"github.com/jhenriquem/go-neovim/internal/editor"
	"github.com/jhenriquem/go-neovim/internal/screen"
)

func KeymapsInsert(eventKey *tcell.EventKey) {
	_, screenHeight := screen.Screen.Size()
	switch eventKey.Key() {

	case tcell.KeyEscape:
		editor.Editor.Mode = "NORMAL"

	case tcell.KeyDelete:
		if editor.Editor.Buffer.CurrentColumn < len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
			(editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine] = append((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine][:editor.Editor.Buffer.CurrentColumn], (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine][editor.Editor.Buffer.CurrentColumn+1:]...)
		}

	case tcell.KeyEnter:

		currLineText := (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine][editor.Editor.Buffer.CurrentColumn:]
		(editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine] = (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine][:editor.Editor.Buffer.CurrentColumn]

		newLines := append([][]rune{}, (editor.Editor.Buffer.Text)[:editor.Editor.Buffer.CurrentLine+1]...)
		newLines = append(newLines, currLineText)
		editor.Editor.Buffer.Text = append(newLines, (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine+1:]...)

		editor.Editor.Buffer.CurrentLine++
		editor.Editor.Buffer.CurrentColumn = 0

		// aplicação do scroll
		if editor.Editor.Buffer.CurrentLine >= config.ScrollOffSet+screenHeight-config.ScrollOffNumber {
			config.ScrollOffSet++
		}

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if editor.Editor.Buffer.CurrentColumn > 0 {

			editor.Editor.Buffer.CurrentColumn--
			(editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine] = append((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine][:editor.Editor.Buffer.CurrentColumn], (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine][editor.Editor.Buffer.CurrentColumn+1:]...)

		} else if editor.Editor.Buffer.CurrentLine > 0 {

			prevLine := (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine-1]
			editor.Editor.Buffer.CurrentColumn = len(prevLine)

			(editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine-1] = append(prevLine, (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]...)

			editor.Editor.Buffer.Text = append((editor.Editor.Buffer.Text)[:editor.Editor.Buffer.CurrentLine], (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine+1:]...)
			editor.Editor.Buffer.CurrentLine--

			// aplicação do scroll
			if editor.Editor.Buffer.CurrentLine < config.ScrollOffSet+config.ScrollOffNumber && config.ScrollOffSet >= 1 {
				config.ScrollOffSet--
			}

		}

	case tcell.KeyLeft:
		if editor.Editor.Buffer.CurrentColumn > 0 {
			editor.Editor.Buffer.CurrentColumn--
		} else if editor.Editor.Buffer.CurrentLine > 0 {
			editor.Editor.Buffer.CurrentLine--
			editor.Editor.Buffer.CurrentColumn = len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine])
		}

	case tcell.KeyRight:
		if editor.Editor.Buffer.CurrentColumn < len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
			editor.Editor.Buffer.CurrentColumn++
		} else if editor.Editor.Buffer.CurrentLine < len(editor.Editor.Buffer.Text)-1 {
			editor.Editor.Buffer.CurrentLine++
			editor.Editor.Buffer.CurrentColumn = 0
		}

	case tcell.KeyUp:

		if editor.Editor.Buffer.CurrentLine > 0 {
			editor.Editor.Buffer.CurrentLine--
			if editor.Editor.Buffer.CurrentLine < config.ScrollOffSet+config.ScrollOffNumber && config.ScrollOffSet >= 1 {
				config.ScrollOffSet--
			}
			if editor.Editor.Buffer.CurrentColumn > len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
				editor.Editor.Buffer.CurrentColumn = len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine])
			}
		}

	case tcell.KeyDown:
		if editor.Editor.Buffer.CurrentLine < len(editor.Editor.Buffer.Text)-1 {
			editor.Editor.Buffer.CurrentLine++
			if editor.Editor.Buffer.CurrentColumn > len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]) {
				editor.Editor.Buffer.CurrentColumn = len((editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine])
			}
			if editor.Editor.Buffer.CurrentLine >= config.ScrollOffSet+screenHeight-config.ScrollOffNumber {
				config.ScrollOffSet++
			}
		}

	default:
		if eventKey.Rune() != 0 {
			line := (editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine]
			(editor.Editor.Buffer.Text)[editor.Editor.Buffer.CurrentLine] = append(line[:editor.Editor.Buffer.CurrentColumn], append([]rune{eventKey.Rune()}, line[editor.Editor.Buffer.CurrentColumn:]...)...)
			editor.Editor.Buffer.CurrentColumn++
		}
	}
}
