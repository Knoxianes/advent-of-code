package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Check(value string) bool {
	ret := Rule3(value)
	if ret == false {
		return false
	}

	ret = Rule2(value)
	if ret == false {
		return false
	}
	ret, _ = Rule1(value)
	if ret == false {
		return false
	}
	return true
}

func Rule1(value string) (bool, error) {
	return regexp.MatchString("^.*[aeiou].*[aeiou].*[aeiou].*$", value)
}
func Rule2(value string) bool {
	for i, tmp := range value {
		if i == len(value)-1 {
			break
		}
		if byte(tmp) == value[i+1] {
			return true
		}
	}
	return false
}
func Rule3(value string) bool {
	return !strings.Contains(value, "ab") || !strings.Contains(value, "cd") || !strings.Contains(value, "pq") || !strings.Contains(value, "xy")
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	splittedData := strings.Split(string(dat), "\n")
	numberOfNice := 0
	for _, value := range splittedData {
		if value == "" {
			continue
		}
		if Check(value) {
			numberOfNice++
		}
	}
	fmt.Println(numberOfNice)

}
