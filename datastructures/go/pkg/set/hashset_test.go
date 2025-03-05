package set_test

import (
	"slices"
	"testing"

	"github.com/Fotkurz/snippets/datastructures/go/pkg/set"
)

func Test_Add(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		wantRet bool
	}{
		{"false for non existing value", []string{"orange"}, false},
		{"true for existing value", []string{"orange", "orange"}, true},
		{"false for non existing values", []string{"orange", "apple", "banana"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := set.NewHashSet[string]()

			var got bool
			for _, a := range tt.arg {
				got = hs.Add(a)
			}

			if got != tt.wantRet {
				t.Errorf("expected %v but got %v", tt.wantRet, got)
			}
		})
	}
}

func Test_Remove(t *testing.T) {
	tests := []struct {
		name     string
		toAdd    []string
		toRemove string
		wantRet  bool
	}{
		{"false for non existing value", []string{"orange"}, "apple", false},
		{"true for existing value", []string{"orange"}, "orange", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := set.NewHashSet[string]()

			for _, a := range tt.toAdd {
				hs.Add(a)
			}
			got := hs.Remove(tt.toRemove)

			if got != tt.wantRet {
				t.Errorf("expected %v but got %v", tt.wantRet, got)
			}
		})
	}
}

func Test_Contains(t *testing.T) {
	tests := []struct {
		name     string
		toAdd    []string
		contains string
		wantRet  bool
	}{
		{"true for existing value", []string{"orange"}, "orange", true},
		{"false for non existing value", []string{"orange"}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := set.NewHashSet[string]()

			for _, a := range tt.toAdd {
				hs.Add(a)
			}

			got := hs.Contains(tt.contains)
			if got != tt.wantRet {
				t.Errorf("expected %v but got %v", tt.wantRet, got)
			}
		})
	}
}

func Test_Slice(t *testing.T) {
	tests := []struct {
		name  string
		toAdd []string
		want  []string
	}{
		{"slice with one item", []string{"orange"}, []string{"orange"}},
		{"slice with more than one item", []string{"orange", "apple", "banana"}, []string{"orange", "apple", "banana"}},
		{"empty slice", []string{}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := set.NewHashSet[string]()

			for _, a := range tt.toAdd {
				hs.Add(a)
			}

			got := hs.Slice()
			if len(got) != len(tt.want) {
				t.Errorf("expected %v but got %v", tt.want, got)
			}
			for _, v := range got {
				if !slices.Contains(tt.want, v) {
					t.Errorf("expected %v but got %v", tt.want, got)
				}
			}
		})
	}
}

func TestHashSet(t *testing.T) {
	hs := set.NewHashSet[string]()

	values := []string{"orange", "apple", "banana"}
	expectFalse := true
	for _, v := range values {
		expectFalse = hs.Add(v)
	}

	if expectFalse {
		t.Errorf("expected to return false when adding new values to hashset, but got true")
	}

	expectTrue := hs.Add("orange")
	if !expectTrue {
		t.Errorf("expected to return true when adding values already existing in hashset, but got false")
	}

	slc := hs.Slice()
	for _, v := range values {
		if !slices.Contains(slc, v) {
			t.Errorf("expected hashset to contains %s but this is not true", v)
		}
	}

	expectTrue = hs.Remove("orange")
	if !expectTrue {
		t.Errorf("expects true when removing existing values from hashset, but got false")
	}
	expectFalse = hs.Remove("airplance")
	if expectFalse {
		t.Errorf("expects false when removing non-existing values from hashset, but got true")
	}
}

func BenchmarkHashSet(b *testing.B) {
	hs := set.NewHashSet[int]()
	for i := 0; i < 1000000; i++ {
		hs.Add(i)
	}
}
