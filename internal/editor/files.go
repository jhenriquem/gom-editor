package editor

import (
	"bufio"
	"os"

	"github.com/jhenriquem/Gom/internal/buffer"
)

func (this *EditorStruct) ScanFile(file *os.File) {
	// Se o arquivo existir na lista de buffers
	Exist := false

	for i, buffer := range this.Buffers {
		if buffer.NameFile == file.Name() {
			this.CrrBuffer = &this.Buffers[i]
			this.CrrBffIndex = i
			Exist = true
		}
	}

	if Exist {
		return
	}

	// Se o arquivo não existir na lista de buffers
	newBuffer := buffer.BufferStruct{NameFile: file.Name()}

	this.Buffers = append(this.Buffers, newBuffer)
	this.CrrBuffer = &this.Buffers[len(this.Buffers)-1]

	this.CrrBffIndex = len(this.Buffers) - 1

	lineIndex := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scannedLine := scanner.Text()
		this.CrrBuffer.Text = append(this.CrrBuffer.Text, []rune{})
		for _, ch := range scannedLine {
			this.CrrBuffer.Text[lineIndex] = append(this.CrrBuffer.Text[lineIndex], rune(ch))
		}
		lineIndex++
	}
	if lineIndex <= 1 {
		this.CrrBuffer.Text = append(this.CrrBuffer.Text, []rune{})
	}
}

func (this *EditorStruct) WriteFile() {
	file, err := os.Create(this.CrrBuffer.NameFile)
	if err != nil {
	}

	writer := bufio.NewWriter(file)
	for _, line := range this.CrrBuffer.Text {
		linetoWrite := string(line) + "\n"
		writer.WriteString(linetoWrite)
	}
	writer.Flush()
}

func (this *EditorStruct) SaveFile(isNewFile bool) string {
	if this.CrrBuffer.NameFile != "" {
		this.WriteFile()
		if isNewFile {
			return "create"
		} else {
			return "save"
		}
	} else {
		return "unnamed file"
	}
}
