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
