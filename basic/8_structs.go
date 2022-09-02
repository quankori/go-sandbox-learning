package basic

import "fmt"

// Person Defining a struct type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Structures func
func structures() {
	var p Person // Declare p with struct Person
	p.LastName = "Nguyễn"
	p.FirstName = "Quân"
	fmt.Println("Person :", p)

	p1 := Person{"A", "Nuyễn Quân", 26} // Style 1
	fmt.Println("Person1: ", p1)

	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	} // Style 2
	fmt.Println("Person2: ", p2)
}
