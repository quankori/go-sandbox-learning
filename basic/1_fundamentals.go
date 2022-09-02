package basic

import "fmt"

func variable() {
	var a string = "Hello, World" // Style 1
	b := "Hello, B"               // Style 2
	b = "New"
	fmt.Println("Variable a: ", a)
	fmt.Println("Variable b: ", b)
}

func constant() {
	const (
		a = 5
	)
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a + 7i // Số phức
	fmt.Println("Variable int: ", intVar)
	fmt.Println("Variable int32: ", int32Var)
	fmt.Println("Variable float: ", float64Var)
	fmt.Println("Variable complex64: ", complex64Var)
}

func function(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
