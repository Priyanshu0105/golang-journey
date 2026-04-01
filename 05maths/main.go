package main

import (
	"fmt"
	"math"
	"math/big"

	//"math/rand/v2"
	"crypto/rand"
)

func main() {
	fmt.Println("The square root of 16 is", math.Sqrt(16))
	fmt.Println("The value of Pi is", math.Pi)
	fmt.Println("The value of E is", math.E)
	fmt.Println("The value of the golden ratio is", math.Phi)
	// random number

	//repowind fmt.Println(rand.IntN(5) + 1) // generates a random number between 0 and 4, so we add 1 to get a number between 1 and 5
	//random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(6)) //generates a random number between 0 and 5
	fmt.Println(myRandomNumber)
	//repowind remind me to play with big.Int and big.Float in the future
}
