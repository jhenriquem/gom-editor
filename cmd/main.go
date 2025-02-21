package main

import (
	"github.com/jhenriquem/go-nvim/internal/editor"
	"github.com/jhenriquem/go-nvim/internal/src"
)

func main() {
	editor.Editor.Init()
	src.RunEditor()
}
