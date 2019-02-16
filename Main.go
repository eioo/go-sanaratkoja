package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Position struct {
	x, y int
}

type Board [][]string

var board Board

func loadWords(filename string) []string {
	var words []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())

		if len(word) > 0 {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words
}

func isSubset(of, arr *[]string) bool {
	set := make(map[string]int)

	for _, value := range *of {
		set[value] += 1
	}

	for _, value := range *arr {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

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

func (pos *Position) getNeighbours() []Position {
	var n []Position
	boardSize := len(board)

	if pos.y-1 >= 0 && pos.x-1 >= 0					{ n = append(n, Position {pos.x-1, pos.y-1}) }
	if pos.y-1 >= 0 && pos.x+1 < boardSize 			{ n = append(n, Position {pos.x+1, pos.y-1}) }
	if pos.y-1 >= 0 					  			{ n = append(n, Position {pos.x, pos.y-1}) }

	if pos.x-1 >= 0 								{ n = append(n, Position {pos.x-1, pos.y}) }
	if pos.x+1 < boardSize							{ n = append(n, Position {pos.x+1, pos.y}) }

	if pos.y+1 < boardSize && pos.x-1 >= 0        	{ n = append(n, Position {pos.x-1, pos.y+1}) }
	if pos.y+1 < boardSize && pos.x+1 < boardSize 	{ n = append(n, Position {pos.x+1, pos.y+1}) }
	if pos.y+1 < boardSize 							{ n = append(n, Position {pos.x, pos.y+1}) }

	return n
}

func findWords(board *Board, words *[]string) []string {
	var foundWords []string
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
				neighbours := currentPosition.getNeighbours()
				currentChar := string([]rune(word)[charIndex])
				charFound := false

				for _, neighbour := range neighbours {
					skip := false

					for _, traversedPos := range traversedPositions {
						if neighbour.x == traversedPos.x && neighbour.y == traversedPos.y {
							skip = true
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

			// Word possible
			if charIndex == wordLength {
				foundWords = append(foundWords, word)
				break
			}
		}
	}

	return foundWords
}

func main() {
	words := loadWords("kotus_sanat.txt")

	board = Board {
		{ "S", "U", "L", "N" },
		{ "I", "I", "I", "K" },
		{ "Ä", "D", "N", "Ä" },
		{ "T", "K", "Ä", "K" },
	}

	start := time.Now()
	foundWords := findWords(&board, &words)

	fmt.Println(strings.Join(foundWords, "\n"))
	fmt.Println("\nTook:", time.Since(start))
}
