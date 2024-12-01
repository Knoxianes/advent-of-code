package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main1() {
	data, _ := os.ReadFile("data.txt")

	var ret int

	dataSplited := strings.Split(string(data), "\n")

	list1 := []int{}
	list2 := []int{}

	for _, value := range dataSplited {
		if value == "" {
			continue
		}
		splitedValue := strings.Split(value, "   ")
		value1, _ := strconv.Atoi(splitedValue[0])
		value2, _ := strconv.Atoi(splitedValue[1])
		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	for i := 0; i < len(list1); i++ {
		ret += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(ret)
}
