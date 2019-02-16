package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    var board Board
    words := loadWords("kotus_sanat.txt")

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf("Type board characters: ")
    scanner.Scan()
    boardStr := strings.ToUpper(scanner.Text())
    rows := chunkString(boardStr, 4)

    for _, row := range rows {
        board = append(board, chunkString(row, 1))
    }

    start := time.Now()
    results := findWords(&board, &words)
    delta := time.Since(start)

    fmt.Printf("\nFound words: %v\nTook: %v\n\n", len(results), delta)

    for _, wordStruct := range results {
        fmt.Println("Solving " + wordStruct.word)
        performWordMouseDrag(wordStruct)
    }
}