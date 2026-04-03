package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruit = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Printf("type of fruitlist: %T\n", fruit)
	fmt.Println("The value of fruit is", fruit)

	fruit = append(fruit, "Elderberry", "Fig", "Grape")
	fmt.Println("The new value of fruit is", fruit)

	fruit = append(fruit[:3])
	fmt.Println(fruit)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 50
	highScores[2] = 25
	highScores[3] = 75
	fmt.Println("High Scores:", highScores)

	highScores = append(highScores, 99, 88, 77)
	fmt.Println("Updated High Scores:", highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove a value from slice based on index
	var fruitSlice = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("Original slice:", fruitSlice)
	indexToRemove := 1 // index of "Banana"
	fruitSlice = append(fruitSlice[:indexToRemove], fruitSlice[indexToRemove+1:]...)
	fmt.Println("Slice after removal:", fruitSlice)
}
