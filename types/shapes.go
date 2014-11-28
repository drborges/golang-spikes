package main

import "fmt"

type Vector struct {
	x, y float64
}

type Shape interface {
	Area() float64
}

type Square struct {
	p1, p2, p3, p4 Vector
}

func (s Square) Height() float64 {
	return s.p3.y - s.p1.y
}

func (s Square) Width() float64 {
	return s.p2.x - s.p1.x
}

func (s Square) Area() float64 {
	return s.Height() * s.Width()
}

func TotalArea(shapes []Shape) float64 {
	var totalArea float64
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {
	p1 := Vector{0, 0}
	p2 := Vector{1, 0}
	p3 := Vector{1, 1}
	p4 := Vector{0, 1}
	p5 := Vector{-1, 0}
	p6 := Vector{-1, -2}
	p7 := Vector{2, 0}

	s1 := Square{p1, p2, p3, p4}
	s2 := Square{p1, p5, p6, p7}

	fmt.Println("Area s1: ", s1.Area())
	fmt.Println("Area s2: ", s2.Area())
	fmt.Println("Area s1 + s2: ", TotalArea([]Shape{s1, s2}))
}
