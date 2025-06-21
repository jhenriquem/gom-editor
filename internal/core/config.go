package core

import g "github.com/jhenriquem/gom-editor/internal/gom"

var Running bool = true

var Gom g.Editor

func Init() {
	Gom = g.Editor{}
	newbuffer := g.Buffer{
		Lines:   []string{""},
		CursorX: 0,
		CursorY: 0,
	}

	Gom.Buffers = append(Gom.Buffers, newbuffer)

	Gom.IndexBuffer = 0
}
