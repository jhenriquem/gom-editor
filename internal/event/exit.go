package event

import "github.com/jhenriquem/gom-editor/internal/core"

func CloseEditor() bool {
	for _, buffer := range core.Gom.Buffers {
		if buffer.IsModified {
			return false
		}
	}
	return true
}
