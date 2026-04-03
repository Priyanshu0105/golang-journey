package main

import (
	"fmt"
)

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for _, day := range days {
		fmt.Println(day)
	}
	// while loop is not available in golang but we can use for loop to achieve the same functionality

	i := 0
	for i < len(days) {
		fmt.Println(days[i])
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// break and continue
	for i := 0; i < len(days); i++ {
		if days[i] == "Thursday" {
			break
		}
		fmt.Println(days[i])
	}

	for i := 0; i < len(days); i++ {
		if days[i] == "Saturday" {
			continue
		}
		fmt.Println(days[i])
	}
	for i := 0; i < len(days); i++ {
		if days[i] == "Saturday" {
			goto huhahaha
		}
		fmt.Println(days[i])
	}
huhahaha:
	fmt.Println("Goto statement is used to jump to a specific line of code")
	// range loop with index and value
	for i, day := range days {
		fmt.Println(day, "is at index", i)
	}
}
