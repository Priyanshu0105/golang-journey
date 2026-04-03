package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["RB"] = "Ruby"

	fmt.Println("The value of languages is", languages)
	fmt.Println("The value of languages[GO] is", languages["GO"])

	delete(languages, "RB")
	fmt.Println("The value of languages after deletion is", languages)

	for key, value := range languages {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

}
