package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards    []string
	Bid      int
	Strength int
}

func (hand *Hand) CalculateStrength() {
	hand.Strength = 1
	joinCards := strings.Join(hand.Cards, "")
	jokerCount := strings.Count(joinCards, "J")
	joinCards = strings.Replace(joinCards, "J", "", -1)

	for i := 0; i < 5; i++ {
		count := strings.Count(joinCards, hand.Cards[i])
		joinCards = strings.Replace(joinCards, hand.Cards[i], "", -1)
		tmpStrength := 1
		switch count {
		case 5:
			tmpStrength = 7
		case 4:
			tmpStrength = 6
		case 3:
			if hand.Strength == 2 {
				tmpStrength = 5
			} else {
				tmpStrength = 4
			}
		case 2:
			if hand.Strength == 2 {
				tmpStrength = 3
			} else if hand.Strength == 4 {
				tmpStrength = 5
			} else {
				tmpStrength = 2
			}

		}
		if tmpStrength > hand.Strength {
			hand.Strength = tmpStrength
		}
	}
	switch jokerCount {
	case 5:
		hand.Strength = 7
	case 4:
		hand.Strength = 7
	case 3:
		if hand.Strength == 2 {
			hand.Strength = 7
		} else {
			hand.Strength = 6
		}
	case 2:
		if hand.Strength == 4 {
			hand.Strength = 7
		} else if hand.Strength == 2 {
			hand.Strength = 6
		} else {
			hand.Strength = 4
		}
	case 1:
		switch hand.Strength {
		case 6:
            hand.Strength = 7
		case 4:
            hand.Strength = 6
		case 3:
            hand.Strength = 5
		case 2:
            hand.Strength = 4
		case 1:
            hand.Strength = 2
		}
	}
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	splited := strings.Split(string(dat), "\n")
	hands := []Hand{}
	for _, value := range splited {
		if value == "" {
			continue
		}
		tmp := strings.Split(value, " ")
		bid, _ := strconv.Atoi(tmp[1])
		hand := Hand{
			Cards: strings.Split(tmp[0], ""),
			Bid:   bid,
		}
		hand.CalculateStrength()
		hands = append(hands, hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		return CompareHands(hands[i], hands[j])
	})
	res := 0
	for i, value := range hands {
		res += (i + 1) * value.Bid
	}
	fmt.Println(res)

}
func CompareHands(a Hand, b Hand) bool {
	if a.Strength == b.Strength {
		return CheckCards(a.Cards, b.Cards)
	}
	return a.Strength < b.Strength
}
func CheckCards(a []string, b []string) bool {
	strengthA := getValue(a[0])
	strengthB := getValue(b[0])
	if strengthA == strengthB {
		return CheckCards(a[1:], b[1:])
	}
	return strengthA < strengthB
}

func getValue(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		switch str {
		case "T":
			return 10
		case "J":
			return 1
		case "Q":
			return 13
		case "K":
			return 14
		case "A":
			return 15
		}
	}
	return value
}
