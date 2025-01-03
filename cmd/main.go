package main

import (
	"github.com/jhenriquem/go-neovim/internal/editor"
	"github.com/jhenriquem/go-neovim/internal/src"
)

func main() {
	editor.Editor.Init()
	src.RunEditor()
}
