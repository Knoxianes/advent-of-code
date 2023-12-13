package main

import (
	"fmt"
	"os"
	"strings"
)

type Wire struct{
    Name string
    Value uint16
}
func Exists(wires []Wire,name string) Wire{
    for   _,value := range wires{
        if value.Name == name{
            return  value
        }
    }
    return Wire{Name:""}
}

func NOT (wires []Wire, inName string, outName string) []Wire {
    return  wires
}

func main(){
    wires := []Wire{}
    dat, _ := os.ReadFile("input.txt") 
    splitedDat := strings.Split(string(dat),"\n")

    for _, row:= range splitedDat{
        if row == ""{
            continue
        }

        splitedRow := strings.Split(row," ")
        if splitedRow[0] == "NOT"{
            NOT(wires,splitedRow[1],splitedRow[3])
        }else{
            if splitedRow[1] == "OR"{
            }else if splitedRow[1] == "AND"{
            }else if splitedRow[1] == "LSHIFT"{
            }else if splitedRow[1] == "RSHIFT"{
            }else{
            }
        }
    }



    var ret uint16 = 0
    for _, value := range wires{
        if value.Name == "a"{
            ret = value.Value
            break    
        }
    }
    fmt.Println(ret)

}
