package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculate(value string) int {
	firstnumber := 0
	lastnumber := 0
	firstNumberIndex := -1
	lastNumberIndex := -1
	for i, char := range value {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		if firstnumber == 0 {
			firstnumber = number
			lastnumber = number
			firstNumberIndex = i
			lastNumberIndex = i
		}
		if lastnumber == number {
            lastNumberIndex = i
			continue
		}
		lastnumber = number
		lastNumberIndex = i
	}
    numbers := CheckForSpelling(value)
    for _, number := range numbers{
        if firstNumberIndex == -1{
            firstNumberIndex = number.Index
            firstnumber = number.Value
            lastnumber = number.Value
            lastNumberIndex = number.Index
            continue
        }
        if firstNumberIndex > number.Index{
            firstnumber = number.Value
            firstNumberIndex = number.Index
            continue
        }
        if lastNumberIndex < number.Index{
            lastnumber = number.Value
            lastNumberIndex = number.Index
        }

    }

	ret, err := strconv.Atoi(strconv.Itoa(firstnumber) + strconv.Itoa(lastnumber))
	if err != nil {
		return 0
	}
	fmt.Println(value, firstnumber, lastnumber, ret)
	return ret
}

func CheckForSpelling(value string) []Number {
	ret := []Number{}
	if strings.Contains(value, "one") {
		for i := range value {
            if i == len(value)-2{
                break
            }
			if value[i] == 111 && value[i+1] == 110 && value[i+2] == 101 {
				ret = append(ret, Number{Index: i, Value: 1})
			}
		}
	}
	if strings.Contains(value, "two") {
		for i := range value {
            if i == len(value)-2{
                break
            }
			if value[i] == 116 && value[i+1] == 119 && value[i+2] == 111 {
				ret = append(ret, Number{Index: i, Value: 2})
			}
		}
	}
	if strings.Contains(value, "three") {
		for i := range value {
            if i == len(value)-4{
                break
            }
			if value[i] == 116 && value[i+1] == 104 && value[i+2] == 114 && value[i+3] == 101 && value[i+4] == 101 {
				ret = append(ret, Number{Index: i, Value: 3})
			}
		}
	}
	if strings.Contains(value, "four") {
		for i := range value {
            if i == len(value)-3{
                break
            }
			if value[i] == 102 && value[i+1] == 111 && value[i+2] == 117 && value[i+3] == 114 {
				ret = append(ret, Number{Index: i, Value: 4})
			}
		}
	}
	if strings.Contains(value, "five") {
		for i := range value {
            if i == len(value)-3{
                break
            }
			if value[i] == 102 && value[i+1] == 105 && value[i+2] == 118 && value[i+3] == 101 {
				ret = append(ret, Number{Index: i, Value: 5})
			}
		}
	}
	if strings.Contains(value, "six") {
		for i := range value {
            if i == len(value)-2{
                break
            }
			if value[i] == 115 && value[i+1] == 105 && value[i+2] == 120 {
				ret = append(ret, Number{Index: i, Value: 6})
			}
		}
	}
	if strings.Contains(value, "seven") {
		for i := range value {
            if i == len(value)-4{
                break
            }
			if value[i] == 115 && value[i+1] == 101 && value[i+2] == 118 && value[i+3] == 101 && value[i+4] == 110 {
				ret = append(ret, Number{Index: i, Value: 7})
			}
		}
	}
	if strings.Contains(value, "eight") {
		for i := range value {
            if i == len(value)-4{
                break
            }
			if value[i] == 101 && value[i+1] == 105 && value[i+2] == 103 && value[i+3] == 104 && value[i+4] == 116 {
				ret = append(ret, Number{Index: i, Value: 8})
			}
		}
	}
	if strings.Contains(value, "nine") {
		for i := range value {
            if i == len(value)-3{
                break
            }
			if value[i] == 110 && value[i+1] == 105 && value[i+2] == 110 && value[i+3] == 101 {
				ret = append(ret, Number{Index: i, Value: 9})
			}
		}
	}
    fmt.Println(value, ret)

	return ret
}

type Number struct {
	Index int
	Value int
}

func main() {

	dat, _ := os.ReadFile("input.txt")
	splitted := strings.Split(string(dat), "\n")

	ret := 0
	for _, value := range splitted {
		if value == "" {
			continue
		}
		ret += Calculate(value)

	}

	fmt.Println(ret)

}
