package main

import (
	"fmt"
	"os"
	"slices"
)

type Position struct {
	X int
	Y int
}

func Move(value byte,currentPosition *Position, seenAtleastOnce *int,seen *[]Position)  {
	switch value {
	case 60: // left
		currentPosition.X--
	case 62: // right
		currentPosition.X++
	case 118: // down
		currentPosition.Y--
	case 94: // up
		currentPosition.Y++
	default:
		return 
	}
	if slices.Contains(*seen, *currentPosition) {
		return 
	}
	*seenAtleastOnce++
	*seen = append(*seen, *currentPosition)

}

func main() {
	dat, _ := os.ReadFile("input.txt")
	seenAtleastOnce := 1
	currentPositionSanta := Position{X: 0, Y: 0}
	currentPositionRoboSanta := Position{X: 0, Y: 0}
	seen := []Position{currentPositionSanta}
	for i, value := range dat {
        if i % 2 == 0{
            Move(value,&currentPositionSanta,&seenAtleastOnce,&seen)
        }else{
            Move(value,&currentPositionRoboSanta,&seenAtleastOnce,&seen)
        }
	}
	fmt.Println(seenAtleastOnce)
}
