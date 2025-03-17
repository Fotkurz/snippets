package shape_test

import (
	"goprototype/internal/shape"
	"testing"
)

func Test_Rectangle_Clone(t *testing.T) {
	first := shape.NewRectangle(5, 10)

	second := first.Clone()

	if &first == &second {
		t.Errorf("Expected object %v to not be %v", &first, &second)
	}

	if !first.Equals(second) {
		t.Errorf("Object is not a clone")
	}
}

func Test_Rectangle_Area(t *testing.T) {
	rect := shape.NewRectangle(10, 10)
	got := rect.Area()
	var want float64 = 100

	if got != want {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

func Test_Circle_Clone(t *testing.T) {
	first := shape.NewCircle(10)

	second := first.Clone()

	if &first == &second {
		t.Errorf("Expected object %v to not be %v", &first, &second)
	}

	if !first.Equals(second) {
		t.Errorf("Object is not a clone")
	}
}

func Test_Circle_Area(t *testing.T) {
	circle := shape.NewCircle(10)
	got := circle.Area()
	var want float64 = 314.1592653589793

	if got != want {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
