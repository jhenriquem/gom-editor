package utils

import (
	"github.com/gdamore/tcell/v2"
)

func KeymapsEvents(eventKey *tcell.EventKey, lines *[][]rune, crrCol, crrLine *int) {
	switch eventKey.Key() {

	case tcell.KeyDelete:
		if *crrCol < len((*lines)[*crrLine]) {
			(*lines)[*crrLine] = append((*lines)[*crrLine][:*crrCol], (*lines)[*crrLine][*crrCol+1:]...)
		}

	case tcell.KeyEnter:
		currLineText := (*lines)[*crrLine][*crrCol:]
		(*lines)[*crrLine] = (*lines)[*crrLine][:*crrCol]

		newLines := append([][]rune{}, (*lines)[:*crrLine+1]...)
		newLines = append(newLines, currLineText)
		*lines = append(newLines, (*lines)[*crrLine+1:]...)

		*crrLine++
		*crrCol = 0

	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if *crrCol > 0 {
			*crrCol--
			(*lines)[*crrLine] = append((*lines)[*crrLine][:*crrCol], (*lines)[*crrLine][*crrCol+1:]...)
		} else if *crrLine > 0 {
			prevLine := (*lines)[*crrLine-1]
			*crrCol = len(prevLine)
			(*lines)[*crrLine-1] = append(prevLine, (*lines)[*crrLine]...)
			*lines = append((*lines)[:*crrLine], (*lines)[*crrLine+1:]...)
			*crrLine--
		}

	case tcell.KeyLeft:
		if *crrCol > 0 {
			*crrCol--
		} else if *crrLine > 0 {
			*crrLine--
			*crrCol = len((*lines)[*crrLine])
		}

	case tcell.KeyRight:
		if *crrCol < len((*lines)[*crrLine]) {
			*crrCol++
		} else if *crrLine < len(*lines)-1 {
			*crrLine++
			*crrCol = 0
		}

	case tcell.KeyUp:
		if *crrLine > 0 {
			*crrLine--
			if *crrCol > len((*lines)[*crrLine]) {
				*crrCol = len((*lines)[*crrLine])
			}
		}

	case tcell.KeyDown:
		if *crrLine < len(*lines)-1 {
			*crrLine++
			if *crrCol > len((*lines)[*crrLine]) {
				*crrCol = len((*lines)[*crrLine])
			}
		}

	default:
		if eventKey.Rune() != 0 {
			line := (*lines)[*crrLine]
			(*lines)[*crrLine] = append(line[:*crrCol], append([]rune{eventKey.Rune()}, line[*crrCol:]...)...)
			*crrCol++
		}
	}
}
