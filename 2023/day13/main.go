package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	splitedData := strings.Split(string(data), "\n\n")
	ret := 0
	for _, value := range splitedData {
		if value == "" {
			continue
		}
		matrix := [][]byte{}
		splitedInput := strings.Split(value, "\n")
		for _, row := range splitedInput {
			if row == "" {
				continue
			}
			tmp := []byte{}
			for i := range row {
				tmp = append(tmp, row[i])
			}
			matrix = append(matrix, tmp)
		}
		ret += Calculate(matrix)
	}
	fmt.Println(ret)
}
func PrintMatrix(input [][]byte) {
	for _, value := range input {
		fmt.Println(string(value))
	}
	fmt.Println()
}

func Calculate(input [][]byte) int {
	rows := Count(input)
	if rows != 0 {
		return rows * 100
	}
	rotatedInput := rotateMatrix(input)
	columns := Count(rotatedInput)
	if columns != 0 {
		return columns
	}
	fmt.Println("----------------")

	for _, value:= range input{
		fmt.Println(string(value))
	}

	fmt.Println()
	for _, value:= range rotatedInput{
		fmt.Println(string(value))
	}
	fmt.Println("----------------")

	return 0
}

func Count(matrix [][]byte) int {
	firstRow := matrix[0]
	secondRow := matrix[1]
	lastRow := matrix[len(matrix)-1]
	secondToLastRow := matrix[len(matrix)-2]
	if slices.Compare(firstRow, lastRow) == 0 {
		return len(matrix) / 2
	} else if slices.Compare(firstRow, secondToLastRow) == 0 {
		return len(matrix) / 2
	} else if slices.Compare(secondRow, lastRow) == 0 {
		return len(matrix)/2 + 1
	} else if slices.Compare(secondRow, secondToLastRow) == 0 {
		return len(matrix)/2 + 1
	}
	return 0
}
func rotateMatrix(matrix [][]byte) (out [][]byte) {
	// Create and populate a new out matrix full of zero
	for i := 0; i < len(matrix[0]); i++ {
		out = append(out, []byte{})
		for j := 0; j < len(matrix); j++ {
			out[i] = append(out[i], 0)
		}
	}
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(out[0]); j++ {
			out[i][j] = matrix[j][i]
		}
	}
	return
}
