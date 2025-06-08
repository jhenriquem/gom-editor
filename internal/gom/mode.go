package gom

import (
	"strings"
)

var modes = [...]string{
	"NORMAL",
	"EDIT",
	"COMMAND",
	"EXPLORER",
}

var current_mode string = modes[0]

func Get_mode() string {
	return current_mode
}

func Set_mode(nextMode string) {
	mode := strings.ToUpper(nextMode)

	for i, m := range modes {
		if mode == m {
			current_mode = modes[i]
		}
	}
}
