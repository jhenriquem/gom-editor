package gom

type Editor struct {
	Buffers     []Buffer
	IndexBuffer int
}

func (e *Editor) GetCurrentBuffer() *Buffer {
	return &e.Buffers[e.IndexBuffer]
}
