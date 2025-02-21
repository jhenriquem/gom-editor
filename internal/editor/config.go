package editor

import (
	"os"
)

type BufferStruct struct {
	Text          [][]rune
	CurrentColumn int
	CurrentLine   int
	IsModificed   bool
}

type EditorStruct struct {
	Message        string
	Mode           string
	Currentfile    string
	Buffer         BufferStruct
	Running        bool
	CurrentCommand []rune
}

func (this *EditorStruct) LoadArgsFile() {
	if len(os.Args) > 1 {
		this.Currentfile = os.Args[1]
	}

	file, err := os.Open(this.Currentfile)

	if err == nil {
		this.ScanFile(file)
	}
}

func (this *EditorStruct) Init() {
	this.Running = true
	this.Buffer.CurrentColumn, this.Buffer.CurrentLine = 0, 0

	this.Message = "Hello World, GO Nvim"

	this.Buffer.Text = append(this.Buffer.Text, []rune{})
	this.CurrentCommand = []rune{':'}

	this.Mode = "NORMAL"

	this.LoadArgsFile()
}

var Editor EditorStruct = EditorStruct{}
