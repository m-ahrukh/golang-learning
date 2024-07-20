package structsmethodsandinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

/*
func Perimeter(height float64, width float64) float64 {
	return 2*height + 2*width
}

func Area(height float64, width float64) float64 {
	return height * width
}
*/

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
