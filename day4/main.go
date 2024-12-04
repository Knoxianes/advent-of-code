package main

import (
	"fmt"
	"os"
	"strings"
)

type Cordinate struct {
	x int
	y int
}

type Word struct {
	Cordinates []Cordinate
	word       string
}

var maze [][]rune

func main() {
	data, _ := os.ReadFile("data.txt")
	dataStringSplited := strings.Split(string(data), "\n")
	for _, line := range dataStringSplited {
		if line == "" {
			continue
		}
		maze = append(maze, []rune(line))
	}

	words := findWords()
	fmt.Println(len(words))
}

func findWords() []Word {
	var words []Word
	for i := 1; i < len(maze)-1; i++ {
		for j := 1; j < len(maze[i])-1; j++ {
			if maze[i][j] != 'A' {
				continue
			}
			traverse(i, j, &words)
		}
	}
	return words
}

func traverse(i, j int, words *[]Word) {
	topLeftChar := maze[i-1][j-1]
	topRightChar := maze[i-1][j+1]
	bottomLeftChar := maze[i+1][j-1]
	bottomRightChar := maze[i+1][j+1]

	mainDiagonal := false
	otherDiagonal := false

	if (topLeftChar == 'M' && bottomRightChar == 'S') || (topLeftChar == 'S' && bottomRightChar == 'M') {
		mainDiagonal = true
	}

	if (topRightChar == 'M' && bottomLeftChar == 'S') || (topRightChar == 'S' && bottomLeftChar == 'M') {
		otherDiagonal = true
	}

	if mainDiagonal && otherDiagonal {
		*words = append(*words, Word{})
	}
}
