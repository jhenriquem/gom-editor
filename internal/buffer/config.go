package buffer

type BufferStruct struct {
	Text          [][]rune
	CurrentColumn int
	CurrentLine   int
	IsModificed   bool
}
