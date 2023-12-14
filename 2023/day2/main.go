package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Possible(game string) int{
    var max =  map[string]int{}
    max["green"] = 0
    max["blue"] = 0
    max["red"] = 0
    gameSplitted := strings.Split(game,";")
    for _, miniGame := range gameSplitted{
        miniGameSplited := strings.Split(miniGame,",")
        for _, color := range miniGameSplited{
            colorSplited := strings.Split(color, " ")
            value, _ := strconv.Atoi(colorSplited[1])
            if max[colorSplited[2]] <  value{
                max[colorSplited[2]]= value
            }
        }
    }
    return max["green"] * max["blue"] * max["red"]
}


func main(){
    dat, _ := os.ReadFile("input.txt")
    ret := 0
    splited := strings.Split(string(dat),"\n")
    for _, row := range splited{
        if row == ""{
            continue
        }
        game := strings.Split(row,":")[1]
        ret += Possible(game)
        
    }
    fmt.Println(ret)
}

