package utils

import "github.com/jhenriquem/Gom/internal/editor"

func MessageCommand(message string) {
	for _, char := range message {
		editor.GOM.CrrCommand = append(editor.GOM.CrrCommand, char)
	}
}
