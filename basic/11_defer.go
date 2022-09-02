package basic

import "fmt"

func defers() {
	defer fmt.Println("World") // Pause print "World"
	fmt.Println("Hello")       // Print text "Hello"

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// Result
	// 9 8 7 6 5 4 3 2 1 0
}
