package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data.txt")
	maze := createMaze(data)
	x, y := findStart(maze)
	direction := "up"
	maze[y][x] = "x"
	for {
		// printMaze(maze)
		// fmt.Println("x:", x, "y:", y, "direction:", direction, "len(maze):", len(maze), "len(maze[y]):", len(maze[y]))
		if direction == "up" {
			y--
			if y < 0 {
				break
			}
			if maze[y][x] == "#" {
				y++
				direction = "right"
				continue
			}
		}

		if direction == "down" {
			y++
			if y >= len(maze)-1 {
				break
			}
			if maze[y][x] == "#" {
				y--
				direction = "left"
				continue
			}
		}

		if direction == "left" {
			x--
			if x < 0 {
				break
			}
			if maze[y][x] == "#" {
				x++
				direction = "up"
				continue
			}
		}
		if direction == "right" {
			x++
			if x >= len(maze[y]) {
				break
			}
			if maze[y][x] == "#" {
				x--
				direction = "down"
				continue
			}
		}
		maze[y][x] = "x"
	}

	fmt.Println(countX(maze))
}

func printMaze(maze [][]string) {
	for _, row := range maze {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func createMaze(data []byte) [][]string {
	lines := strings.Split(string(data), "\n")
	maze := make([][]string, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		maze[i] = strings.Split(line, "")
	}
	return maze
}

func countX(maze [][]string) int {
	count := 0
	for _, row := range maze {
		for _, cell := range row {
			if cell == "x" {
				count++
			}
		}
	}
	return count
}

func findStart(maze [][]string) (int, int) {
	for y, row := range maze {
		for x, cell := range row {
			if cell == "^" {
				return x, y
			}
		}
	}
	return -1, -1
}
