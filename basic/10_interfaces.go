package basic

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	x, y, r float64
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// Interfaces func
func interfaces() {
	c := circle{0, 0, 5}
	fmt.Println(totalArea(&c))
}
