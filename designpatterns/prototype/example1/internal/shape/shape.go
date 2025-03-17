package shape

import "math"

type Shape interface {
	Clone() Shape
	Equals(s Shape) bool
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func NewRectangle(width, height float64) Shape {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func newRectangleFrom(rect Rectangle) Shape {
	return &Rectangle{
		width:  rect.width,
		height: rect.height,
	}
}

func (r *Rectangle) Equals(s Shape) bool {
	v, ok := s.(*Rectangle)
	if !ok {
		return false
	}

	if r.height != v.height {
		return false
	}
	if r.width != v.width {
		return false
	}

	return true
}

// Clone implements Shape.
func (r *Rectangle) Clone() Shape {
	// clone is able to use private
	//	object fields to generate a new one since its on the same pkg
	newR := newRectangleFrom(*r)

	return newR
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Shape {
	return &Circle{
		radius: radius,
	}
}

func newCircleFrom(circle Circle) Shape {
	return &Circle{
		radius: circle.radius,
	}
}

func (c *Circle) Clone() Shape {
	newC := newCircleFrom(*c)

	return newC
}

func (r *Circle) Equals(s Shape) bool {
	v, ok := s.(*Circle)
	if !ok {
		return false
	}

	if r.radius != v.radius {
		return false
	}

	return true
}

func (r *Circle) Area() float64 {
	return math.Pi * math.Pow(r.radius, 2)
}
