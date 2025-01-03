package editor

type BufferStruct struct {
	Text          [][]rune
	CurrentColumn int
	CurrentLine   int
	IsModificed   bool
}

type EditorStruct struct {
	Mode           string
	Currentfile    string
	Buffer         BufferStruct
	Running        bool
	CurrentCommand []rune
}

func (this *EditorStruct) Init() {
	this.Running = true
	this.Buffer.CurrentColumn, this.Buffer.CurrentLine = 0, 0

	this.Buffer.Text = append(this.Buffer.Text, []rune{})
	this.CurrentCommand = []rune{':'}

	this.Mode = "NORMAL"
}

var Editor EditorStruct = EditorStruct{}
