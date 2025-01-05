package editor

import (
	"bufio"
	"os"
)

func (this *EditorStruct) ScanFile(file *os.File) {
	this.Buffer.Text = [][]rune{}
	lineIndex := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scannedLine := scanner.Text()
		this.Buffer.Text = append(this.Buffer.Text, []rune{})
		for _, ch := range scannedLine {
			this.Buffer.Text[lineIndex] = append(this.Buffer.Text[lineIndex], rune(ch))
		}
		lineIndex++
	}
	if lineIndex <= 1 {
		this.Buffer.Text = append(this.Buffer.Text, []rune{})
	}
}

func (this *EditorStruct) WriteFile() {
	file, err := os.Create(this.Currentfile)
	if err != nil {
	}

	writer := bufio.NewWriter(file)
	for _, line := range this.Buffer.Text {
		linetoWrite := string(line) + "\n"
		writer.WriteString(linetoWrite)
	}
	writer.Flush()
}

func (this *EditorStruct) SaveFile(isNewFile bool) {
	if this.Currentfile != "" {
		this.WriteFile()
		if isNewFile {
			this.Message = "Create File"
		} else {
			this.Message = "Save file"
		}
	} else {
		this.Message = "Unnamed file"
	}
}
