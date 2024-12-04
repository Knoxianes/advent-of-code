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
	fmt.Println(calculateAnswer(words))
}

func calculateAnswer(words []Word) int {
	var tmpWords []Word
	for _, word := range words {
		if !isCoridnatesInList(word.Cordinates, tmpWords) {
			tmpWords = append(tmpWords, word)
		}
	}

	return len(tmpWords)
}

func isCoridnatesInList(coridnates []Cordinate, words []Word) bool {
	ret := false
	for _, word := range words {
		numberOfExists := 0
		for _, coridnate := range coridnates {
			exists := false
			for _, wordCoridnate := range word.Cordinates {
				if coridnate.x == wordCoridnate.x && coridnate.y == wordCoridnate.y {
					exists = true
					break
				}
			}
			if exists {
				numberOfExists++
			}
		}
		if numberOfExists == 4 {
			return true
		}
	}
	return ret
}

func findWords() []Word {
	var words []Word
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			traverse(i, j, &words)
		}
	}
	return words
}

func traverse(i, j int, words *[]Word) {
	maxLeft := j - 3
	maxRight := j + 3
	maxUp := i - 3
	maxDown := i + 3

	if maxLeft >= 0 {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := j; k >= maxLeft; k-- {
			runes = append(runes, maze[i][k])
			cordinates = append(cordinates, Cordinate{x: k, y: i})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}

	if maxRight < len(maze[i]) {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := j; k <= maxRight; k++ {
			runes = append(runes, maze[i][k])
			cordinates = append(cordinates, Cordinate{x: k, y: i})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}

	if maxUp >= 0 {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := i; k >= maxUp; k-- {
			runes = append(runes, maze[k][j])
			cordinates = append(cordinates, Cordinate{x: j, y: k})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}

	if maxDown < len(maze) {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := i; k <= maxDown; k++ {
			runes = append(runes, maze[k][j])
			cordinates = append(cordinates, Cordinate{x: j, y: k})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}

	if maxLeft >= 0 && maxUp >= 0 {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := 0; k <= 3; k++ {
			runes = append(runes, maze[i-k][j-k])
			cordinates = append(cordinates, Cordinate{x: j - k, y: i - k})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}

	if maxLeft >= 0 && maxDown < len(maze) {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := 0; k <= 3; k++ {
			runes = append(runes, maze[i+k][j-k])
			cordinates = append(cordinates, Cordinate{x: j - k, y: i + k})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}

	if maxRight < len(maze[i]) && maxUp >= 0 {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := 0; k <= 3; k++ {
			runes = append(runes, maze[i-k][j+k])
			cordinates = append(cordinates, Cordinate{x: j + k, y: i - k})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)

	}

	if maxRight < len(maze[i]) && maxDown < len(maze) {
		var word Word
		var cordinates []Cordinate
		runes := []rune{}
		for k := 0; k <= 3; k++ {
			runes = append(runes, maze[i+k][j+k])
			cordinates = append(cordinates, Cordinate{x: j + k, y: i + k})
		}
		word.word = string(runes)
		word.Cordinates = cordinates
		*words = append(*words, word)
	}
}
