package main

import (
	"github.com/jheriquem/Gom/internal/editor"
	"github.com/jheriquem/Gom/internal/src"
	// "github.com/jhenriquem/Gom/internal/editor"
	// "github.com/jhenriquem/Gom/internal/src"
)

func main() {
	editor.Editor.Init()
	src.RunEditor()
}
