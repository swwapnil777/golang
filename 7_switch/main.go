package main

import (
	"fmt"
	"time"
)

func main() {

	i := 4

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")

	default:
		fmt.Println("Other")

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Print("Its Weekend")

	default:
		fmt.Println("Its Workday")
	}
}
