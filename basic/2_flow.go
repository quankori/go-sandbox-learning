package basic

import (
	"fmt"
	"time"
)

func condition() {
	fmt.Print("Check variable even or odd")
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func loop() {
	fmt.Println("Loop for")
	array := [3]int{12, 78, 50}
	for i := 1; i <= 10; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println()
	fmt.Println("Loop range")

	for i, s := range array {
		fmt.Println(i, s)
	}
}

func switchCase() {
	fmt.Print("Check day: ")
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}
}
