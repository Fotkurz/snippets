package pangram_test

import (
	"pangrams/pangram"
	"testing"
)

func TestIsPangram(t *testing.T) {
	t.Run("is pangram", func(t *testing.T) {
		input := "We promptly judged antique ivory buckles for the next prize"
		want := "pangram"
		got := pangram.IsPangram(input)

		if want != got {
			t.Errorf("expected=(%s), got=(%s)", want, got)
		}
	})

	t.Run("not pangram", func(t *testing.T) {
		input := "We promptly judged antique ivory buckles for the prize"
		want := "not pangram"
		got := pangram.IsPangram(input)

		if want != got {
			t.Errorf("expected=(%s), got=(%s)", want, got)
		}
	})
}
