package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := "Priyanshu"
	age := 21
	fmt.Println("Hello my name is", name, "and my age is", age)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age :")
	//repowind comma ok || err syntax ->input, _ := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	fmt.Println("karma is a bitch", input)
	fmt.Printf("you entered: %T", input)
}
