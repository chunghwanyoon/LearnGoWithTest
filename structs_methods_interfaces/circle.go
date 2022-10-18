package structs_methods_interfaces

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius // Pi * r^2
}
