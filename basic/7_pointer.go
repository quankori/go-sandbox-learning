package basic

import "fmt"

func pointer() {
	var a = 100
	var p *int
	fmt.Println(p) // Default is null
	p = &a
	fmt.Println("a = ", a)
	fmt.Println("p = ", p)   // Index of pointer a
	fmt.Println("*p = ", *p) // Value of pointer a

	*p = 2000
	fmt.Println("a (after) = ", a) // Change by pointer
}
