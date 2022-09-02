package advanced

import "fmt"

func Execute() {
	fmt.Println("=== Testing marshal ===")
	marshal()
	fmt.Println("=== Testing goroutines ===")
	goroutines()
	fmt.Println("=== Testing mutex ===")
	mutex()
	fmt.Println("=== Testing channels ===")
	channel()
}
