package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}