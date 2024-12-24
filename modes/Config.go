package modes

import "github.com/gdamore/tcell/v2"

var CurrentMODE string = "NORMAL"

func KeymapsLogicModes(eventKey *tcell.EventKey) {
	switch CurrentMODE {
	case "NORMAL":
		KeymapsNormal(eventKey)
	case "INSERT":
		KeymapsInsert(eventKey)
	case "COMMAND":
		KeymapsCommand(eventKey)
	}
}
