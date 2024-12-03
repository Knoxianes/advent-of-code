package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	isSafe bool
	data   []int
	isAsc  bool
}

func main() {
	data, _ := os.ReadFile("data.txt")

	splitedDAta := strings.Split(string(data), "\n")

	reports := make([]Report, len(splitedDAta))

	for i, v := range splitedDAta {
		var tmpReport Report
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

		reports[i] = tmpReport
	}

	scan(&reports)

	fmt.Println(countSafe(reports))
}

func scan(reports *[]Report) {
	for j, report := range *reports {

		tmpReports := make([]Report, len(report.data))

		for i := 0; i < len(report.data); i++ {
			var newReport Report
			copyOfOriginalSlice := make([]int, len(report.data))
			copy(copyOfOriginalSlice, report.data)

			newReport.data = remove(copyOfOriginalSlice, i)
			newReport.isSafe = true
			if newReport.data[0] > newReport.data[1] {
				newReport.isAsc = false
			} else if newReport.data[0] < newReport.data[1] {
				newReport.isAsc = true
			}
			tmpReports[i] = newReport
		}

		for k, tmpReport := range tmpReports {
			for i := 0; i < len(tmpReport.data)-1; i++ {
				firstNum := tmpReport.data[i]
				secondNum := tmpReport.data[i+1]

				if firstNum == secondNum {
					tmpReport.isSafe = false
					break
				}

				if tmpReport.isAsc && (firstNum > secondNum || secondNum-firstNum > 3) {
					tmpReport.isSafe = false
					break
				}
				if !tmpReport.isAsc && (firstNum < secondNum || firstNum-secondNum > 3) {
					tmpReport.isSafe = false
					break
				}
			}
			tmpReports[k] = tmpReport
		}

		isSafe := false
		for _, tmpReport := range tmpReports {
			if tmpReport.isSafe {
				isSafe = true
				break
			}
		}
		report.isSafe = isSafe
		(*reports)[j] = report
	}
}

func isAsc(report Report) bool {
	countAsc := 0
	countDesc := 0
	for i := 0; i < len(report.data)-1; i++ {
		firstNum := report.data[i]
		secondNum := report.data[i+1]

		if firstNum < secondNum {
			countAsc++
		}
		if firstNum > secondNum {
			countDesc++
		}
	}
	return countAsc > countDesc
}

func countSafe(reports []Report) int {
	safeCount := 0
	for _, report := range reports {
		if report.isSafe {
			safeCount++
		}
	}
	return safeCount
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
