package editor

import (
	"os"

	"github.com/jhenriquem/go-nvim/internal/buffer"
)

type EditorStruct struct {
	Mode           string
	Currentfile    string
	Buffer         buffer.BufferStruct
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

	this.Buffer.Text = append(this.Buffer.Text, []rune{})
	this.CurrentCommand = []rune{':'}

	this.Mode = "NORMAL"

	this.LoadArgsFile()
}

var Editor EditorStruct = EditorStruct{}
