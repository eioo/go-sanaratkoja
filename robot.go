package main

import "github.com/go-vgo/robotgo"

func performWordMouseDrag(wordStruct SolverResult) {
	tileSize := 125
	baseX := 97  // TODO: Don't hard-code. Nox window should be at 0, 0 with size: 600, 1020
	baseY := 483

	for i, pos := range wordStruct.traversedPositions {
		destX := baseX + pos.x * tileSize
		destY := baseY + pos.y * tileSize

		if i == 0 {
			robotgo.MoveMouse(destX, destY)
			robotgo.MouseToggle("down")
		} else {
			robotgo.DragMouse(destX, destY)
		}

		robotgo.MilliSleep(100)
	}

	robotgo.MouseToggle("up")
}
