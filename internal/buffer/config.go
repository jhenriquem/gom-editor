package buffer

type BufferStruct struct {
	NameFile      string
	Text          [][]rune
	CurrentColumn int
	CurrentLine   int
	IsModificed   bool
}
