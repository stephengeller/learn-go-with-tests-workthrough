package internal

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Width + rectangle.Height) * 2

}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (triangle Triangle) Area() float64 {
	return triangle.Base * triangle.Height * 0.5
}
