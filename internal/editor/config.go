package editor

import (
	"github.com/jhenriquem/Gom/internal/buffer"
)

type EditorStruct struct {
	CrrMode     string
	CrrBuffer   *buffer.BufferStruct
	Buffers     []buffer.BufferStruct
	Running     bool
	CrrCommand  []rune
	CrrBffIndex int
}

var GOM EditorStruct = EditorStruct{
	Running:     true,
	CrrMode:     "NORMAL",
	CrrBffIndex: 0,
	CrrCommand:  []rune{':'},
}

func Inicializer() {
	GOM.CrrBuffer = GOM.InicializeBuffer()
	GOM.LoadArgsFile()
}
