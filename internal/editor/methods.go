package editor

import (
	"os"

	"github.com/jhenriquem/Gom/internal/buffer"
)

func (e *EditorStruct) LoadArgsFile() {
	if len(os.Args) > 1 {
		e.CrrBuffer.NameFile = os.Args[1]
	}

	file, err := os.Open(e.CrrBuffer.NameFile)

	if err == nil {
		e.ScanFile(file)
	}
}

func (e *EditorStruct) InicializeBuffer() *buffer.BufferStruct {
	// Garante que há pelo menos um buffer disponível
	if len(e.Buffers) == 0 {
		newBuffer := buffer.BufferStruct{}
		e.Buffers = append(e.Buffers, newBuffer)
	}

	// Se CrrBuffer for nil, inicializa ele corretamente
	if e.CrrBuffer == nil {
		e.CrrBuffer = &e.Buffers[0]
	}

	e.Buffers[0].NameFile = ""

	if e.CrrBuffer.Text == nil {
		e.CrrBuffer.Text = [][]rune{}
	}

	e.Buffers[0].Text = append(e.CrrBuffer.Text, []rune{})

	return &e.Buffers[0]
}
