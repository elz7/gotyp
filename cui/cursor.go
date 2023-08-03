package cui

import (
	"math"

	"github.com/awesome-gocui/gocui"
)

func cursorUp(v *gocui.View) int {
	_, curLine := v.Cursor()
	_, curOrig := v.Origin()

	_, viewLinesCount := v.Size()
	bufferLinesCount := v.LinesHeight()

	switch {
	case curLine == 0 && curOrig == 0: // the first line
		curOrig = viewLinesCount * (int(math.Ceil(float64(bufferLinesCount)/float64(viewLinesCount))) - 1)
		curLine = viewLinesCount - curOrig

	case curLine == 0 && curOrig != 0: // the first line somewhere in the middle
		curOrig = curOrig - viewLinesCount
		curLine = viewLinesCount - 1

	default:
		curLine = curLine - 1
	}

	v.SetOrigin(0, curOrig)
	v.SetCursor(0, curLine)

	return curOrig + curLine
}

func cursorDown(v *gocui.View) int {
	_, curLine := v.Cursor()
	_, curOrig := v.Origin()

	curLine = curLine + curOrig // current line within the buffer, not the view

	_, viewLinesCount := v.Size()
	bufferLinesCount := v.LinesHeight()

	switch {
	case curLine == bufferLinesCount-1: // the last line
		curOrig = 0
		curLine = 0
	case (curLine+1)%viewLinesCount == 0: // last line somewhere in the middle
		curOrig = curOrig + viewLinesCount
		curLine = 0
	default:
		curLine = curLine + 1
	}

	v.SetOrigin(0, curOrig)
	v.SetCursor(0, curLine)

	return curOrig + curLine
}

func cursorPos(v *gocui.View) int {
	_, y := v.Cursor()
	_, o := v.Origin()
	return y + o
}
