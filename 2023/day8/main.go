package main

import (
	"fmt"
	"os"
	"strings"
)
type Node struct{
    Name string
    Left string
    Right string
}

func main(){
    dat, _ := os.ReadFile("input.txt")
    splitted := strings.Split(string(dat), "\n\n")
    mapa, start := addMap(splitted[1]) 
    repeatNumber := make([]int,len(start))
    for i, value := range start{
        tmp := 0
        Walk(mapa,splitted[0],value,&tmp)
        repeatNumber[i] = tmp
    }
    fmt.Println(LCM(repeatNumber[0],repeatNumber[1],repeatNumber...))

}
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func addMap(input string)(map[string]Node,[]string){
    split := strings.Split(input,"\n")
    start := []string{}
    ret := map[string]Node{}
    for _,value := range split{
        
        if value == ""{
            continue
        }
        tmp := Node{}
        splitValue := strings.Split(value," = ")

        name := splitValue[0]
        tmp.Name = name
        if name[2] == 65{
            start = append(start,name)
        }

        splitValue[1] = strings.Replace(splitValue[1],"(","",-1)
        splitValue[1] = strings.Replace(splitValue[1],")","",-1)

        leftRight := strings.Split(splitValue[1],", ")
        tmp.Left = leftRight[0]
        tmp.Right = leftRight[1]
        ret[name] = tmp
    }

    return ret,start
}
func Walk(mapa map[string]Node,instructions string, startPoint string, steps *int){
    current := startPoint
     
    for _,instruction := range instructions{
        if instruction == 76{   // Left
                current = mapa[current].Left
            (*steps)++
            if current[2] == 90{
                return
            }
        }else{                  // Right
            current = mapa[current].Right
            (*steps)++
            if current[2] == 90{
                return
            }
        }
    }
    Walk(mapa,instructions,current,steps)
}
