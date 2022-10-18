package structs_methods_interfaces

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Name when variable is needed
func (r Rectangle) Name(prefix string) string {
	return prefix + " " + "Rectangle"
}
