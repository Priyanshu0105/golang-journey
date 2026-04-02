package main

import (
	"fmt"
)

func main() {
	var fruit [4]string
	fruit[0] = "Apple"
	fruit[2] = "Cherry"
	fruit[3] = "Date"
	fmt.Println("The value of fruit is", fruit)
	fmt.Println("The length of fruit is", len(fruit))
}
