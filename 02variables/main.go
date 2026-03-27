package main

import "fmt"

func main() {
	var name string = " Priyanshu"
	fmt.Println(name)
	var age int = 21
	fmt.Println(age)
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of this var is: %T \n", isLoggedIn)
	var smallINT uint8 = 255
	fmt.Println(smallINT)
	fmt.Printf("The type of this var is: %T \n", smallINT)
	var smallFloat float32 = 255.4555555555555555555555555555555555
	fmt.Println(smallFloat)
	fmt.Printf("The type of this var is: %T \n", smallFloat)
}
