package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Number struct {
	value      int
	startIndex int
	endIndex   int
	rowIndex   int
}

func calculate(ret *int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] != 42 {
				continue
			}
			*ret += FoundStar(i, j)
		}
	}
}

func FoundStar(row int, column int) int {
	numbers := []int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !(graph[row+i-1][column+j-1] > 47 && graph[row+i-1][column+j-1] < 58) {
				continue
			}
			number := createNumber(row+i-1, column, &j)
			if slices.Contains(numbers,number){
				fmt.Println(numbers, number)
			}
			numbers = append(numbers, number)
		}
	}
	if len(numbers) < 2 {
		return 0
	}
	if slices.Contains(numbers, 705) {
		fmt.Println(numbers)
	}
	ret := 0
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			ret += numbers[i] * numbers[j]
		}
	}

	return ret
}
func createNumber(row int, column int, j *int) int {
	tmpNumber := []rune{graph[row][column-1+*j]}
	startLeft := column - 2 + *j  
	*j++
	for ; ; *j++ {
		if column+*j-1 == len(graph[row]) {
			break
		}
		if !(graph[row][column+*j-1] > 47 && graph[row][column+*j-1] < 58) {
			break
		}
		tmpNumber = append(tmpNumber, graph[row][column+*j-1])
	}
	for ; ; startLeft-- {
		if startLeft == -1 {
			break
		}
		if !(graph[row][startLeft] > 47 && graph[row][startLeft] < 58) {
			break
		}
		tmpNumber = append([]rune{graph[row][startLeft]}, tmpNumber...)
	}
	ret, err := strconv.Atoi(string(tmpNumber))
	if err != nil {
		return 0
	}
	return ret
}

var graph = [][]rune{}

func main() {
	dat, _ := os.ReadFile("input.txt")
	splited := strings.Split(string(dat), "\n")
	for _, row := range splited {
		var tmp []rune
		if row == "" {
			continue
		}
		for _, tmpRune := range row {
			if tmpRune == 0 {
				continue
			}
			tmp = append(tmp, tmpRune)
		}
		graph = append(graph, tmp)

	}

	ret := 0
	calculate(&ret)
	fmt.Println(ret)
}
