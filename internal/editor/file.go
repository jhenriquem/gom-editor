package editor

import (
	"bufio"
	"os"

	"github.com/jhenriquem/go-neovim/global"
)

func WriteFile() {
	file, err := os.Create(global.CurrentFileName)
	if err != nil {
	}
	writer := bufio.NewWriter(file)
	for _, line := range global.Lines {
		linetoWrite := string(line) + "\n"
		writer.WriteString(linetoWrite)
	}
	writer.Flush()
}

func SaveFile(isNewFile bool) {
	if !isNewFile {
		global.Command = []rune{'s', 'a', 'v', 'e', ' ', 'f', 'i', 'l', 'e'}
	} else {
		global.Command = []rune{'c', 'r', 'e', 'a', 't', 'e', ' ', 'f', 'i', 'l', 'e'}
	}
	WriteFile()
}
