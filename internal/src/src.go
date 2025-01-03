package src

import (
	"bufio"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/jhenriquem/go-nvim/internal/editor"
	"github.com/jhenriquem/go-nvim/internal/keymaps"
	"github.com/jhenriquem/go-nvim/internal/screen"
)

func RunEditor() {
	screen.ScreenInitializer()
	defer screen.Screen.Fini()

	for editor.Editor.Running {

		if editor.Editor.Mode != "COMMAND" {
			editor.Editor.Buffer.RenderBuffer()
		} else {
			editor.Editor.CommandLine()
		}
		editor.Editor.StatusLine()

		ev := screen.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:

			screen.Screen.Sync()

		case *tcell.EventKey:

			file, err := os.Create("log.txt")
			if err != nil {
			}
			writer := bufio.NewWriter(file)
			for _, line := range editor.Editor.Buffer.Text {
				linetoWrite := string(line) + "\n"
				writer.WriteString(linetoWrite)
			}
			writer.Flush()

			keymaps.KeymapsLogicModes(ev)
		}
	}
}
