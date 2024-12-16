package modes

import (
	"github.com/gdamore/tcell/v2"
)

func KeymapsEventsForNormalMode(eventKey *tcell.EventKey) {
	switch eventKey.Rune() {
	case 105:
		CurrentMODE = "INSERT"
	}
}
