package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (cir Circle) Area() float64 {
	return math.Pi * (cir.Radius * cir.Radius)
}

func (tri Triangle) Area() float64{
	return (tri.Height * tri.Base) / 2
}
