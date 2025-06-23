package fileio

import (
	"bufio"
	"os"
)

func Scan(filename string) []string {
	text := []string{}

	file, err := os.Open(filename)
	if err != nil {
		file, _ = os.Create(filename)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scannedLine := scanner.Text()

		text = append(text, scannedLine)
	}

	return text
}
