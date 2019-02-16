package main

func (board *Board) get(x, y int) string {
	return (*board)[y][x]
}

func (board *Board) toArray() []string {
	var allChars []string

	for _, row := range *board {
		for _, col := range row {
			allChars = append(allChars, col)
		}
	}

	return allChars
}

func (board *Board) findCharPositions(char string) []Position {
	var positions []Position

	for y, row := range *board {
		for x := range row {
			if board.get(x, y) == char {
				pos := Position { x, y }
				positions = append(positions, pos)
			}
		}
	}

	return positions
}

func (pos *Position) getNeighbours(board *Board) []Position {
	var n []Position
	boardSize := len(*board)

	if pos.y-1 >= 0 && pos.x-1 >= 0               { n = append(n, Position {pos.x-1, pos.y-1}) }
	if pos.y-1 >= 0 && pos.x+1 < boardSize        { n = append(n, Position {pos.x+1, pos.y-1}) }
	if pos.y-1 >= 0                               { n = append(n, Position {pos.x, pos.y-1}) }

	if pos.x-1 >= 0                               { n = append(n, Position {pos.x-1, pos.y}) }
	if pos.x+1 < boardSize                        { n = append(n, Position {pos.x+1, pos.y}) }

	if pos.y+1 < boardSize && pos.x-1 >= 0        { n = append(n, Position {pos.x-1, pos.y+1}) }
	if pos.y+1 < boardSize                        { n = append(n, Position {pos.x, pos.y+1}) }
	if pos.y+1 < boardSize && pos.x+1 < boardSize { n = append(n, Position {pos.x+1, pos.y+1}) }

	return n
}

func findWords(board *Board, words *[]string) []SolverResult {
	var result []SolverResult
	allChars := board.toArray()

	for _, word := range *words {
		var chars []string
		wordLength := len([]rune(word))

		for _, c := range word {
			chars = append(chars, string(c))
		}

		if !isSubset(&allChars, &chars) {
			continue
		}

		// Do board search
		startPositions := board.findCharPositions(chars[0])

		for _, startPos := range startPositions {
			traversedPositions := []Position { startPos }
			charIndex := 1

			for charIndex < wordLength {
				currentPosition := traversedPositions[len(traversedPositions) - 1]
				neighbours := currentPosition.getNeighbours(board)
				currentChar := string([]rune(word)[charIndex])
				charFound := false

				for _, neighbour := range neighbours {
					skip := false

					for _, traversedPos := range traversedPositions {
						if neighbour == traversedPos {
							skip = true
							break
						}
					}

					if skip {
						continue
					}

					neighbourChar := board.get(neighbour.x, neighbour.y)

					if neighbourChar == currentChar {
						traversedPositions = append(traversedPositions, Position { neighbour.x, neighbour.y })
						charIndex++
						charFound = true
						break
					}
				}

				if !charFound {
					break
				}
			}

			if charIndex == wordLength {
				result = append(result, SolverResult{ word, traversedPositions })
				break
			}
		}
	}

	return result
}