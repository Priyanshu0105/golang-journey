package main

import (
	"fmt"
)

func main() {
	var ptr *int                            // declaring a pointer to an int
	fmt.Println("The value of ptr is", ptr) // prints <nil>, because ptr is not initialized
	a := 5
	b := &a // b is a pointer to a
	fmt.Println("The value of a is", a)
	fmt.Println("The value of b is", b)                               // prints the memory address of a
	fmt.Println("The value at the memory address stored in b is", *b) // dereferencing b to get the value of a

	// changing the value of a through b
	*b = 10
	fmt.Println("The new value of a is", a) // prints 10, because we changed the value of a through b

	// creating a pointer to an array
	arr := [3]int{1, 2, 3}
	arrPtr := &arr[0]
	fmt.Println("The value of arr is", arr)
	fmt.Println("The value of arrPtr is", arrPtr)                               // prints the memory address of arr
	fmt.Println("The value at the memory address stored in arrPtr is", *arrPtr) // dereferencing arrPtr to get the value of arr

	// changing the value of arr through arrPtr
	(*arrPtr) = 10
	fmt.Println("The new value of arr is", arr) // prints [10 2 3], because we changed the value of arr through arrPtr
}
