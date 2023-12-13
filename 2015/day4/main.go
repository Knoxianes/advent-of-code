package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)



func main(){
    input := "yzbqklnj"
    i := 1
    for{
        tmpInput := input+strconv.Itoa(i)
        result := fmt.Sprintf("%x",md5.Sum([]byte(tmpInput))) 
        if result[:6] == "000000"{
            break
        }
        i++
    }
    fmt.Println(i)
}
