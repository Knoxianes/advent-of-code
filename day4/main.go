package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Compare(leftSide []string, rightSide []string, gameID int) int {
	numberOfWinningNumbers := 0
	for _, leftNumber := range leftSide {
		if leftNumber == "" {
			continue
		}
		for _, rightNumber := range rightSide {
			if rightNumber == "" {
				continue
			}
			if leftNumber == rightNumber {
				numberOfWinningNumbers++
				break
			}
		}
	}
	for i := 1; i <= numberOfWinningNumbers; i++ {
		index := gameID + i
		cards[index] += 1 * cards[gameID]
	}
	return cards[gameID]
}

var cards = map[int]int{}

func main() {
	dat, _ := os.ReadFile("input.txt")
	splited := strings.Split(string(dat), "\n")
	ret := 0
	for _, row := range splited {
		if row == "" {
			continue
		}
		card := strings.Split(strings.Split(row, ":")[1], "|")
		tmp := strings.Split(strings.Split(row, ":")[0], " ")
		gameID, err := strconv.Atoi(tmp[1])
		if err != nil {
			gameID, err = strconv.Atoi(tmp[2])
			if err != nil {
				gameID, _ = strconv.Atoi(tmp[3])
			}
		}
		cards[gameID] += 1
		ret += Compare(strings.Split(card[0], " "), strings.Split(card[1], " "), gameID)
	}
	fmt.Println(cards)
	fmt.Println(ret)

}
