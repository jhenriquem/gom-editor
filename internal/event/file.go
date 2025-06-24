package event

import (
	"bufio"
	"os"

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

		for n, line := range originalText {
			if line != buffer.Lines[n] {
				core.Gom.Buffers[i].IsModified = true
				return
			}
		}
		core.Gom.Buffers[i].IsModified = false
	}
}

func Savefile() {
	buffer := core.Gom.GetCurrentBuffer()

	if buffer.Filename == "" {
		buffer.Filename = "no_name.txt"
	}
	file, _ := os.Create(buffer.Filename)

	writer := bufio.NewWriter(file)
	for _, line := range buffer.Lines {
		linetoWrite := line + "\n"
		writer.WriteString(linetoWrite)
	}
	writer.Flush()
}
