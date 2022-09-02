package basic

import (
	"fmt"
	"math"
)

type rectangle struct {
	length int
	width  int
}

type circleA struct {
	radius float64
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (c circleA) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circleA) changeValue(newValue float64) {
	c.radius = newValue
}

// Methods func
func methods() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.area())
	c := circleA{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.area())
	// Con trỏ
	fmt.Printf("\n\nbefore change: %f", c.radius)
	(&c).changeValue(51)
	fmt.Printf("\nafter change: %f", c.radius)
	fmt.Println()
}
