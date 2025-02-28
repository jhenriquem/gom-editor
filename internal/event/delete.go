package event

import "github.com/jhenriquem/Gom/internal/editor"

func Delete() {
	line := editor.GOM.CrrBuffer.Text[editor.GOM.CrrBuffer.CurrentLine]
	if editor.GOM.CrrBuffer.CurrentColumn < len(line) {
		editor.GOM.CrrBuffer.Text[editor.GOM.CrrBuffer.CurrentLine] = append(line[:editor.GOM.CrrBuffer.CurrentColumn], line[editor.GOM.CrrBuffer.CurrentColumn+1:]...)
	}
}
