package event

import (
	"github.com/jhenriquem/gom-editor/internal/core"
	fileio "github.com/jhenriquem/gom-editor/internal/file_io"
)

func ModificedFile() {
	for i, buffer := range core.Gom.Buffers {
		originalText := fileio.Scan(buffer.Filename)

		if len(buffer.Lines) != len(originalText) {
			core.Gom.Buffers[i].IsModified = true
			return
		}

		for i, line := range buffer.Lines {
			if line != originalText[i] {
				core.Gom.Buffers[i].IsModified = true
				return
			}
		}
		core.Gom.Buffers[i].IsModified = false
	}
}
