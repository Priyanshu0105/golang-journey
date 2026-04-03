package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dieceNumber := rand.Intn(6) + 1
	switch dieceNumber {
	case 1:
		fmt.Println("you can move 1 step")
	case 2:
		fmt.Println("you can move 2 steps")
	case 3:
		fmt.Println("you can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("you can move 4 steps")
	case 5:
		fmt.Println("you can move 5 steps")
	case 6:
		fmt.Println("you can move 6 steps and roll the diece again")
	default:
		fmt.Println("invalid number")
	}

}
