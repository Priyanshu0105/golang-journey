package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("The current time is", presentTime)

	// formatting time
	fmt.Println("Formatted time is", presentTime.Format("01-02-2006 15:04:05 Monday"))

	// creating a specific time
	createdDate := time.Date(2004, time.April, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created date is", createdDate)

	fmt.Println("Formatted created date is", createdDate.Format("01-02-2006 Monday"))
}
