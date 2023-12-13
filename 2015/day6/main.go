package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func TurnON(grid [][]int, start Position, end Position) [][]int {
    for y := start.Y; y <= end.Y; y++{
        for x:=start.X; x <= end.X; x++{
            grid[y][x]++
        }
    }
    return grid
}
func TurnOff(grid [][]int, start Position, end Position) [][]int {
    for y := start.Y; y <= end.Y; y++{
        for x:=start.X; x <= end.X; x++{
            grid[y][x]--
            if grid[y][x] < 0 {
                grid[y][x] = 0
            }
        }
    }
    return grid
}
func Toggle(grid [][]int, start Position, end Position)[][]int {
    for y := start.Y; y <= end.Y; y++{
        for x:=start.X; x <= end.X; x++{
            grid[y][x] += 2
        }
    }
    return grid
}

func main() {
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		for y := 0; y < 1000; y++ {
			grid[i] = append(grid[i], 0)
		}

	}
	dat, _ := os.ReadFile("input.txt")
	splittedData := strings.Split(string(dat), "\n")
	for _, row := range splittedData {
		if row == "" {
			continue
		}
		splittedCommand := strings.Split(row, " ")
		if splittedCommand[0] == "turn" {
			startCordinatesTmp := strings.Split(splittedCommand[2], ",")
			endCordinatesTmp := strings.Split(splittedCommand[4], ",")
			startCordinateX, _ := strconv.Atoi(startCordinatesTmp[0])
			startCordinateY, _ := strconv.Atoi(startCordinatesTmp[1])
			endCordinateX, _ := strconv.Atoi(endCordinatesTmp[0])
			endtCordinateY, _ := strconv.Atoi(endCordinatesTmp[1])

			startCordinate := Position{X: startCordinateX, Y: startCordinateY}
			endCordinate := Position{X: endCordinateX, Y: endtCordinateY}

			if splittedCommand[1] == "on" {
				grid = TurnON(grid, startCordinate, endCordinate)
			} else {
				grid = TurnOff(grid, startCordinate, endCordinate)
			}

		} else {
			startCordinatesTmp := strings.Split(splittedCommand[1], ",")
			endCordinatesTmp := strings.Split(splittedCommand[3], ",")
			startCordinateX, _ := strconv.Atoi(startCordinatesTmp[0])
			startCordinateY, _ := strconv.Atoi(startCordinatesTmp[1])
			endCordinateX, _ := strconv.Atoi(endCordinatesTmp[0])
			endtCordinateY, _ := strconv.Atoi(endCordinatesTmp[1])

			startCordinate := Position{X: startCordinateX, Y: startCordinateY}
			endCordinate := Position{X: endCordinateX, Y: endtCordinateY}

            grid = Toggle(grid,startCordinate,endCordinate)
		}
	}

	var counter = 0
	for i := 0; i < 1000; i++ {
		for y := 0; y < 1000; y++ {
            counter += grid[i][y]
		}

	}
	fmt.Println(counter)

}
