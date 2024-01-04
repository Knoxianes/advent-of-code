package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Cordinate struct {
	X int
	Y int
}

var start Cordinate
var count = 0

func main() {
	dat, _ := os.ReadFile("test.txt")
	maze := createMaze(string(dat))
	O := []Cordinate{}
	Walk(maze, findStart(maze), Cordinate{X: -1, Y: -1}, 1, &O)
	CountInside(maze, O)
	fmt.Println(count)

}

func isInside(maze [][]byte, current Cordinate, O []Cordinate) bool {
	if slices.Contains(O, current) {
		return false
	}
	numberOfCrosses := 0
	for i := current.X; i < len(maze[0]); i++ {
		tmp := Cordinate{X: i, Y: current.Y}
		if slices.Contains(O, tmp) {
			up := Cordinate{X: i, Y: current.Y - 1}
			down := Cordinate{X: i, Y: current.Y + 1}
			

			if slices.Contains(O, up) {
				if maze[up.Y][up.X] == 124 || maze[up.Y][up.X] == 70 || maze[up.Y][up.X] == 55 {
					numberOfCrosses++
					continue
				}
			}
			if slices.Contains(O, down) {
				if maze[down.Y][down.X] == 124 || maze[down.Y][down.X] == 74 || maze[down.Y][down.X] == 76 {
					numberOfCrosses++
					continue
				}
			}
		}
	}

	if numberOfCrosses%2 == 0 {
		return false
	}
	return true
}
func CountInside(maze [][]byte, O []Cordinate) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if isInside(maze, Cordinate{Y: i, X: j}, O) {
				count++
			}
		}
	}
}

func findStart(maze [][]byte) Cordinate {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == 83 {

				return Cordinate{X: x, Y: y}
			}
		}
	}
	return Cordinate{}
}
func createMaze(input string) [][]byte {
	maze := [][]byte{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		tmp := []byte{}
		rowSplited := strings.Split(row, "")
		for _, value := range rowSplited {
			tmp = append(tmp, value[0])
		}
		maze = append(maze, tmp)
	}

	return maze
}

func Walk(maze [][]byte, startPosition Cordinate, lastSeen Cordinate, steps int, O *[]Cordinate) {
	nextPosition := Cordinate{}
	if start == startPosition {
		return
	}
	switch maze[startPosition.Y][startPosition.X] {
	case 124: // |
		if lastSeen.Y == startPosition.Y-1 {
			nextPosition = Cordinate{X: startPosition.X, Y: startPosition.Y + 1}
		} else {
			nextPosition = Cordinate{X: startPosition.X, Y: startPosition.Y - 1}
		}
	case 45: // -
		if lastSeen.X == startPosition.X-1 {
			nextPosition = Cordinate{X: startPosition.X + 1, Y: startPosition.Y}
		} else {
			nextPosition = Cordinate{X: startPosition.X - 1, Y: startPosition.Y}
		}
	case 76: // L
		if lastSeen.Y == startPosition.Y-1 {
			nextPosition = Cordinate{X: startPosition.X + 1, Y: startPosition.Y}
		} else {
			nextPosition = Cordinate{X: startPosition.X, Y: startPosition.Y - 1}
		}
	case 74: // J
		if lastSeen.Y == startPosition.Y-1 {
			nextPosition = Cordinate{X: startPosition.X - 1, Y: startPosition.Y}
		} else {
			nextPosition = Cordinate{X: startPosition.X, Y: startPosition.Y - 1}
		}
	case 55: // 7
		if lastSeen.Y == startPosition.Y+1 {
			nextPosition = Cordinate{X: startPosition.X - 1, Y: startPosition.Y}
		} else {
			nextPosition = Cordinate{X: startPosition.X, Y: startPosition.Y + 1}
		}
	case 70: // F
		if lastSeen.Y == startPosition.Y+1 {
			nextPosition = Cordinate{X: startPosition.X + 1, Y: startPosition.Y}
		} else {
			nextPosition = Cordinate{X: startPosition.X, Y: startPosition.Y + 1}
		}
	case 83: // S
		possiblePosition := findPossibleForStart(maze, startPosition.X, startPosition.Y)
		start = startPosition
		(*O) = append(*O, startPosition)
		(*O) = append(*O, possiblePosition[0])
		Walk(maze, possiblePosition[0], startPosition, steps, O)
		return
	}
	steps++
	(*O) = append(*O, nextPosition)
	Walk(maze, nextPosition, startPosition, steps, O)
}

func findPossibleForStart(maze [][]byte, x int, y int) []Cordinate {
	ret := []Cordinate{}
	if y-1 >= 0 {
		if maze[y-1][x] == 124 || maze[y-1][x] == 70 || maze[y-1][x] == 55 {
			ret = append(ret, Cordinate{X: x, Y: y - 1})
		}
	}
	if y+1 < len(maze) {
		if maze[y+1][x] == 124 || maze[y+1][x] == 74 || maze[y+1][x] == 76 {
			ret = append(ret, Cordinate{X: x, Y: y + 1})
		}
	}
	if x+1 < len(maze[y]) {
		if maze[y][x+1] == 45 || maze[y][x+1] == 74 || maze[y][x+1] == 55 {
			ret = append(ret, Cordinate{X: x + 1, Y: y})
		}
	}

	if x-1 >= 0 {
		if maze[y][x-1] == 45 || maze[y][x-1] == 70 || maze[y][x-1] == 76 {
			ret = append(ret, Cordinate{X: x - 1, Y: y})
		}
	}
	return ret
}
