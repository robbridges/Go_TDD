package Structs

import "math"

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	getArea() float64
}

func getPerimeter(rec Rectangle) float64 {
	return 2 * (rec.width + rec.length)
}

func (rec Rectangle) getArea() float64 {
	return rec.length * rec.width
}

func (cir Circle) getArea() float64 {
	return math.Pi * (cir.radius * cir.radius)
}

func (tri Triangle) getArea() float64 {
	return .5 * (tri.base * tri.height)
}
