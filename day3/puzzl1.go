package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func calculate2(ret *int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if !(graph[i][j] > 47 && graph[i][j] < 58) {
				continue
			}
			number := CreateNumber(&j, i)
			if CheckNumber(number) {
				*ret += number.value
			}
		}
	}
}

func CheckNumber(number Number) bool {
	if number.rowIndex == 0 {
		return CheckNumberFirstRow(number)
	}
	if number.rowIndex == len(graph)-1 {
		return CheckNumberLastRow(number)
	}
	if number.startIndex == 0 {
		return CheckNumberFirstInRow(number)
	}
	if number.endIndex == len(graph[number.rowIndex])-1 {
		return CheckNumberLastInRow(number)
	}
	if graph[number.rowIndex][number.startIndex-1] != 46 ||
		graph[number.rowIndex+1][number.startIndex-1] != 46 ||
		graph[number.rowIndex-1][number.startIndex-1] != 46 ||
		graph[number.rowIndex][number.endIndex+1] != 46 ||
		graph[number.rowIndex+1][number.endIndex+1] != 46 ||
		graph[number.rowIndex-1][number.endIndex+1] != 46 {

		return true
	}
	for i := number.startIndex; i <= number.endIndex; i++ {
		if graph[number.rowIndex+1][i] != 46 || graph[number.rowIndex-1][i] != 46 {
			return true
		}
	}

	return false
}
func CheckNumberFirstRow(number Number) bool {
	if graph[number.rowIndex][number.startIndex-1] != 46 ||
		graph[number.rowIndex+1][number.startIndex-1] != 46 ||
		graph[number.rowIndex][number.endIndex+1] != 46 ||
		graph[number.rowIndex+1][number.endIndex+1] != 46 {

		return true
	}
	for i := number.startIndex; i <= number.endIndex; i++ {
		if graph[number.rowIndex+1][i] != 46 {
			return true
		}
	}
	return false
}
func CheckNumberLastRow(number Number) bool {
	if graph[number.rowIndex][number.startIndex-1] != 46 ||
		graph[number.rowIndex-1][number.startIndex-1] != 46 ||
		graph[number.rowIndex][number.endIndex+1] != 46 ||
		graph[number.rowIndex-1][number.endIndex+1] != 46 {

		return true
	}
	for i := number.startIndex; i <= number.endIndex; i++ {
		if graph[number.rowIndex-1][i] != 46 {
			return true
		}
	}
	return false
}
func CheckNumberFirstInRow(number Number) bool {
	if graph[number.rowIndex][number.endIndex+1] != 46 ||
		graph[number.rowIndex+1][number.endIndex+1] != 46 ||
		graph[number.rowIndex-1][number.endIndex+1] != 46 {

		return true
	}
	for i := number.startIndex; i <= number.endIndex; i++ {
		if graph[number.rowIndex+1][i] != 46 || graph[number.rowIndex-1][i] != 46 {
			return true
		}
	}
	return false
}
func CheckNumberLastInRow(number Number) bool {
	if graph[number.rowIndex][number.startIndex-1] != 46 ||
		graph[number.rowIndex+1][number.startIndex-1] != 46 ||
		graph[number.rowIndex-1][number.startIndex-1] != 46 {

		return true
	}
	for i := number.startIndex; i <= number.endIndex; i++ {
		if graph[number.rowIndex+1][i] != 46 || graph[number.rowIndex-1][i] != 46 {
			return true
		}
	}
	return false
}

func CreateNumber(startIndex *int, rowIndex int) Number {
	var tmp []rune
	var ret Number
	ret.startIndex = *startIndex
	for ; ; *startIndex++ {
		if *startIndex == len(graph[rowIndex]) {
			break
		}
		if !(graph[rowIndex][*startIndex] > 47 && graph[rowIndex][*startIndex] < 58) {
			break
		}
		tmp = append(tmp, graph[rowIndex][*startIndex])
	}
	ret.endIndex = *startIndex - 1
	ret.rowIndex = rowIndex
	tmpValue, err := strconv.Atoi(string(tmp))
	if err != nil {
		panic("Greska kod strconv.Atoi")
	}
	ret.value = tmpValue
	return ret
}


func Start() {
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
