package main

import (
	"fmt"
)

func main() {
	var languages = []string{"Go", "Python", "JavaScript", "Ruby"}
	for i := 0; i < len(languages); i++ {
		if len(languages[i]) > 6 {
			fmt.Println(languages[i])
		} else if len(languages[i]) < 6 {
			fmt.Println("The length of languages is less than 6")
		} else {
			fmt.Println("The length of languages is 6")
		}
	}

	// better way for golang is to use range loop

	for _, language := range languages {
		if language == "Go" {
			fmt.Println("Go is a great language")
		} else if language == "Python" {
			fmt.Println("Python is a great language")
		} else if language == "JavaScript" {
			fmt.Println("JavaScript is a great language")
		} else if language == "Ruby" {
			fmt.Println("Ruby is a great language")
		} else {
			fmt.Println("Unknown language")
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i, "is even")
		}
	}
}
