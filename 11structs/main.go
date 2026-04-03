package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	//repowind no inheritance in golang; no super or parent class; no sub or child class; no is a relationship; only has a relationship
	priyanshu := User{"Priyanshu", "priyanshu4996@gmail.com", true, 21}
	fmt.Println(priyanshu)
	fmt.Printf("Name is %v and email is %v\n", priyanshu.Name, priyanshu.Email)
}
