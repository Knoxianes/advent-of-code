package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	L int
	W int
	H int
}

func (cube *Cube) Calculate() int {
	var tmp []int
	tmp = append(tmp, cube.L, cube.W, cube.H)
	ret := 2*cube.H + 2*cube.W + 2*cube.L
	max := tmp[0]
	for _, value := range tmp {
		if value > max {
			max = value
		}
	}
	ret -= 2 * max
	ribbon := cube.L * cube.H * cube.W
	return ret + ribbon
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	splittedData := strings.Split(string(dat), "\n")
	answere := 0
	for _, value := range splittedData {
		if value == "" {
			continue
		}
		tmp := strings.Split(value, "x")
		l, _ := strconv.Atoi(tmp[0])
		w, _ := strconv.Atoi(tmp[1])
		h, _ := strconv.Atoi(tmp[2])
		cube := Cube{
			L: l,
			W: w,
			H: h,
		}
		answere += cube.Calculate()
	}
	fmt.Println(answere)

}
