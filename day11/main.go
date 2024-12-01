package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Cordinate struct {
	Y int
	X int
}

type Connection struct {
	Length int
	Galaxy *Galaxy
}

type Galaxy struct {
	Number      int
	Connections map[int]Connection
	Cordinate   Cordinate
}
var numberToExpand = 1000000-1
var columnsToExpand = []int{}
var rowsToExpand = []int{}
func main() {
	data, _ := os.ReadFile("input.txt")
	_, mapOfGalaxy := MakeSpace(string(data))
    mapOfGalaxy = Recalculate(mapOfGalaxy)
	mapOfGalaxy = ConnectAllGalaxies(mapOfGalaxy)
    fmt.Println(columnsToExpand,rowsToExpand)
	fmt.Println(Calculate(mapOfGalaxy))

}
func Recalculate(mapOfGalaxy map[int]Galaxy)map[int]Galaxy{
    for i := range mapOfGalaxy{
        galaxy := mapOfGalaxy[i]
        currentRow := galaxy.Cordinate.Y
        currentColumn := galaxy.Cordinate.X
        for _, value:= range rowsToExpand{
            if currentRow > value{
                galaxy.Cordinate.Y += numberToExpand
            }
        }
        for _, value:= range columnsToExpand{
            if currentColumn > value{
                galaxy.Cordinate.X += numberToExpand
            }
        }

        mapOfGalaxy[i] = galaxy

    }
    return mapOfGalaxy
}
func Calculate(mapOfGalaxy map[int]Galaxy) int {
	ret := 0
	for i := 1; i < len(mapOfGalaxy); i++ {
		for j := i + 1; j <= len(mapOfGalaxy); j++ {
			ret += mapOfGalaxy[i].Connections[j].Length
		}
	}
	return ret
}
func ConnectAllGalaxies(mapOfGalaxy map[int]Galaxy) map[int]Galaxy {
	for i := range mapOfGalaxy {
		galaxy := mapOfGalaxy[i]
		for j := range mapOfGalaxy {
			otherGalaxy := mapOfGalaxy[j]
			if galaxy.Number == otherGalaxy.Number {
				continue
			}

			cordinatesOfOtherGalaxy := otherGalaxy.Cordinate
			cordinatesGalaxy := galaxy.Cordinate
			tmpLength := int(math.Abs(float64(cordinatesOfOtherGalaxy.Y-cordinatesGalaxy.Y))) +
				int(math.Abs(float64(cordinatesOfOtherGalaxy.X-cordinatesGalaxy.X)))

			galaxy.Connections[otherGalaxy.Number] = Connection{
				Length: tmpLength,
				Galaxy: &otherGalaxy,
			}

		}
		mapOfGalaxy[i] = galaxy
	}
	return mapOfGalaxy
}
func ExpandColumns(input string) []string {
	rows := strings.Split(input, "\n")
	lenRow := len(rows[0])
	for x := 0; x < lenRow; x++ {
		isColumnEmpty := true
		for y := 0; y < len(rows); y++ {
			if rows[y] == "" {
				continue
			}
			if rows[y][x] == 35 {
				isColumnEmpty = false
			}
		}

		if isColumnEmpty {
            columnsToExpand = append(columnsToExpand,x)
		}
	}
	return rows
}
func ExpandRows(rows []string)[]string{
    length := len(rows)
    for i:= 0; i< length; i++{
        if rows[i] == ""{
            continue
        }
        if strings.Contains(rows[i],"#"){
            continue
        }
        rowsToExpand = append(rowsToExpand,i)

    }
    return rows
}
func MakeSpace(input string) ([][]int, map[int]Galaxy) {
	rows := ExpandColumns(input)
    rows = ExpandRows(rows)
	space := [][]int{}
	mapOfGalaxy := map[int]Galaxy{}
	count := 1
	for y, row := range rows {
		if row == "" {
			continue
		}
		tmp := []int{}
		for i, value := range row {
			if value == 35 {
				tmp = append(tmp, count)
				mapOfGalaxy[count] = Galaxy{
					Number:      count,
					Connections: map[int]Connection{},
					Cordinate:   Cordinate{Y: y, X: i},
				}
				count++
			} else {
				tmp = append(tmp, 0)
			}
		}
		space = append(space, tmp)

	}
	return space, mapOfGalaxy
}
