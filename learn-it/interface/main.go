package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	var s Shape
	s = rect

	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
