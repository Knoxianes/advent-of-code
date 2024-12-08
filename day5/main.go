package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	res := 0
	nodesMap := make(map[int][]int)

	data, _ := os.ReadFile("data.txt")
	splitedData := strings.Split(string(data), "\n\n")

	orderingRules := strings.Split(splitedData[0], "\n")
	input := strings.Split(splitedData[1], "\n")

	for _, rule := range orderingRules {
		if rule == "" {
			continue
		}
		splitedRule := strings.Split(rule, "|")
		firstNumber, _ := strconv.Atoi(splitedRule[0])
		secondNumber, _ := strconv.Atoi(splitedRule[1])
		makeMap(firstNumber, secondNumber, &nodesMap)
	}

	for _, line := range input {
		if line == "" {
			continue
		}

		numberList := []int{}
		for _, numberString := range strings.Split(line, ",") {
			number, _ := strconv.Atoi(numberString)
			numberList = append(numberList, number)
		}

		isSafe, goodOrder := isSafe(&nodesMap, numberList)
		if !isSafe {
			res += goodOrder[len(goodOrder)/2]
		}
	}

	fmt.Println(res)
}

func isSafe(nodesMap *map[int][]int, numberList []int) (bool, []int) {
	ret := true
	for i, x := range numberList {
		listOfNumbers, found := (*nodesMap)[x]
		if !found {
			continue
		}
		for j, y := range numberList {
			if i < j && slices.Contains(listOfNumbers, y) {
				ret = false

				break
			}
		}
	}

	if !ret {
		for i := 0; i < len(numberList); i++ {
			listOfNumbers, found := (*nodesMap)[numberList[i]]
			if !found {
				continue
			}
			for j := 0; j < len(numberList); j++ {
				if i < j && slices.Contains(listOfNumbers, numberList[j]) {
					numberList[i], numberList[j] = numberList[j], numberList[i]
					i = -1
					break
				}
			}
		}
	}

	return ret, numberList
}

func makeMap(firstNumber, secondNumber int, nodesMap *map[int][]int) {
	secondNumberList, secondFound := (*nodesMap)[secondNumber]

	if !secondFound {
		(*nodesMap)[secondNumber] = []int{firstNumber}
	} else {
		(*nodesMap)[secondNumber] = append(secondNumberList, firstNumber)
	}
}
