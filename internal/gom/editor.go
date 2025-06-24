package gom

import fileio "github.com/jhenriquem/gom-editor/internal/file_io"

type Editor struct {
	Buffers     []Buffer
	IndexBuffer int
}

func (e *Editor) GetCurrentBuffer() *Buffer {
	return &e.Buffers[e.IndexBuffer]
}

func (e *Editor) OpenFile(filename string) {
	text := fileio.Scan(filename)

	if len(text) == 0 {
		text = []string{""}
	}
	e.GetCurrentBuffer().Filename = filename
	e.GetCurrentBuffer().Lines = text
}
