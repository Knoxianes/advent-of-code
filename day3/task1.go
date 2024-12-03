package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	ret := 0
	data, _ := os.ReadFile("data.txt")

	longString := string(data)

	for {
		indexDont := strings.Index(longString, "don't()")
		indexDo := strings.Index(longString, "do()")

		if indexDont == -1 {
			break
		}

		if indexDo == -1 {
			longString = longString[:indexDont]
			break
		}

		if indexDo < indexDont {
			longString = longString[:indexDo] + longString[indexDo+3:]
			continue
		}

		longString = longString[:indexDont] + longString[indexDo+3:]
	}

	pattern := `mul\(\d{1,3},\d{1,3}\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(longString, -1)

	for _, match := range matches {
		numberPattern := `(\d{1,3})`
		newRe := regexp.MustCompile(numberPattern)
		numbers := newRe.FindAllString(match, -1)
		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])
		ret += number1 * number2
	}

	// newPattern := `(?<=don't\(\)).*?mul\(\d{1,3},\d{1,3}\).*?(?=do\(\))`

	fmt.Println(ret)
}
