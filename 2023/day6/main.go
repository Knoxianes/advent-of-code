package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct{
    Time int
    Distance int
}

func main(){
    dat, _:= os.ReadFile("input.txt")
    splited := strings.Split(string(dat),"\n")
    timeValues := strings.Split(splited[0], " ")
    distanceValues := strings.Split(splited[1], " ")
    time, _ :=  strconv.Atoi(timeValues[1])
    distance, _ := strconv.Atoi(distanceValues[1])
    fmt.Println(Win(time, distance))
}

func Win(time int, distance int) int{
    ret := 0
    for i:=1 ; i<time; i++{
        if distance < i*(time -i){
            ret++
        }
    }

    return ret
}
