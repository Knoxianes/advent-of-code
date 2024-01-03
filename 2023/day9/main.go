package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, _ := os.ReadFile("input.txt")
	input := strings.Split(string(dat), "\n")
	ret := 0
	for _, value := range input {
		if value == "" {
			continue
		}
        tmp := []int{}
        for _, value = range strings.Split(value," "){
            convertedValue, _ := strconv.Atoi(value)
            tmp = append(tmp, convertedValue)
        }
        ret += FindNext(tmp)
	}
	fmt.Println(ret)
}
func isEnd(input []int) bool {
    for _, value := range input{
        if value != input[0]{
            return false
        }
    }
    return true
}
func FindNext(input []int) int {
    if isEnd(input){
        return input[0]
    }
    nextArray := make([]int,len(input)-1)
    for i:=0; i< len(input)-1; i++{
       nextArray[i] = input[i+1]-input[i] 
    }
    nextValue := input[0] - FindNext(nextArray)
    return nextValue
}
