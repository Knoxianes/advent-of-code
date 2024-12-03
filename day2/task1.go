package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report2 struct {
	isSafe bool
	isAsc  bool
	data   []int
}

func main2() {
	data, _ := os.ReadFile("data.txt")

	splitedDAta := strings.Split(string(data), "\n")

	reports := make([]Report2, len(splitedDAta))

	for i, v := range splitedDAta {
		var tmpReport Report2
		var values []int
		if v == "" {
			continue
		}

		splitedV := strings.Split(v, " ")
		for _, value := range splitedV {
			convertedValue, _ := strconv.Atoi(value)
			values = append(values, convertedValue)
		}

		tmpReport.data = values
		tmpReport.isSafe = true
		if values[0] > values[1] {
			tmpReport.isAsc = false
		} else if values[0] < values[1] {
			tmpReport.isAsc = true
		}

		reports[i] = tmpReport
	}

	scan2(&reports)

	fmt.Println(countSafe2(reports))
}

func scan2(reports *[]Report2) {
	for j, report := range *reports {
		for i := 0; i < len(report.data)-1; i++ {
			firstNum := report.data[i]
			secondNum := report.data[i+1]

			if firstNum == secondNum {
				report.isSafe = false
				break
			}

			if report.isAsc && (firstNum > secondNum || secondNum-firstNum > 3) {
				report.isSafe = false
				break
			}
			if !report.isAsc && (firstNum < secondNum || firstNum-secondNum > 3) {
				report.isSafe = false
				break
			}
		}

		(*reports)[j] = report

	}
}

func countSafe2(reports []Report2) int {
	safeCount := 0
	for _, report := range reports {
		if report.isSafe {
			safeCount++
		}
	}
	return safeCount
}
