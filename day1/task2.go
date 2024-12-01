package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")

	var ret int

	dataSplited := strings.Split(string(data), "\n")

	list1 := []string{}
	list2 := []string{}

	for _, value := range dataSplited {
		if value == "" {
			continue
		}
		splitedValue := strings.Split(value, "   ")
		list1 = append(list1, splitedValue[0])
		list2 = append(list2, splitedValue[1])
	}

	list1Joined := strings.Join(list1, " ")

	for _, value := range list2 {
		intValue, _ := strconv.Atoi(value)
		ret += strings.Count(list1Joined, value) * intValue
	}

	fmt.Println(ret)
}
