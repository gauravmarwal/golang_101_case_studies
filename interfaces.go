package main

import (
	"fmt"
	"math"
)

// Go Interface - `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Struct type `Rectangle` - implements the `Shape` interface by implementing all its methods.
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Struct type `Circle` - implements the `Shape` interface by implementing all its methods.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// Generic function to calculate the total area of multiple shapes of different types
func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {
	var s Shape = Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())

	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 5}, Circle{10})
	fmt.Println("Total area = ", totalArea)
}
