package gom

type Buffer struct {
	Lines      []string
	CursorX    int
	CursorY    int
	Filename   string
	IsModified bool
}
