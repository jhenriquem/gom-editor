package editor

import (
	"os"

	"github.com/jhenriquem/Gom/internal/buffer"
)

type EditorStruct struct {
	Mode           string
	CrrBuffer      *buffer.BufferStruct
	Buffers        []buffer.BufferStruct
	Running        bool
	CurrentCommand []rune
	CrrBufferIndex int
}

func (this *EditorStruct) LoadArgsFile() {
	if len(os.Args) > 1 {
		this.CrrBuffer.NameFile = os.Args[1]
	}

	file, err := os.Open(this.CrrBuffer.NameFile)

	if err == nil {
		this.ScanFile(file)
	}
}

func (this *EditorStruct) InicializeBuffer() {
	// Garante que há pelo menos um buffer disponível
	if len(this.Buffers) == 0 {
		newBuffer := buffer.BufferStruct{}
		this.Buffers = append(this.Buffers, newBuffer)
	}

	// Define CrrBuffer para apontar para o primeiro buffer
	this.CrrBuffer = &this.Buffers[0]
	this.CrrBufferIndex = 0
}

func (this *EditorStruct) Init() {
	this.Running = true

	this.InicializeBuffer()

	this.LoadArgsFile()

	this.CrrBuffer.Text = append(this.CrrBuffer.Text, []rune{})
	this.CurrentCommand = []rune{':'}

	this.Mode = "NORMAL"
}

var Editor EditorStruct = EditorStruct{}
