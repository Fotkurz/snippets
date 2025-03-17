package shape

type Shape interface {
	Clone() Shape
	Equals(s Shape) bool
}

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width, height int) Shape {
	return &Rectangle{
		width:  width,
		height: height,
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
	//	object fields to generate a new one
	newR := Rectangle{
		width:  r.width,
		height: r.height,
	}

	return &newR
}

type Circle struct {
	radius int
}

func NewCircle(radius int) Shape {
	return &Circle{
		radius: radius,
	}
}

// Clone implements Shape.
func (c *Circle) Clone() Shape {
	newC := &Circle{radius: c.radius}

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
