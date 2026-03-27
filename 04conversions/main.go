package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("WELCOME TO OUR PIZZA APP")
	fmt.Println("Please rate our app from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //string to number conversion
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating", numRating+1)
	}
}
