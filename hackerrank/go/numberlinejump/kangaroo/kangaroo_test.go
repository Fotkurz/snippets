package kangaroo_test

import (
	"numberlinejump/kangaroo"
	"testing"
)

func TestKangaroo(t *testing.T) {
	noAns := "NO"
	yesAns := "YES"

	t.Run("yes when 0 3 4 2", func(t *testing.T) {
		got := kangaroo.Kangaroo(0, 3, 4, 2)
		assertEqual(t, got, yesAns)
	})

	t.Run("NO when 0 2 5 3", func(t *testing.T) {
		got := kangaroo.Kangaroo(0, 2, 5, 3)
		assertEqual(t, got, noAns)
	})
}

func TestKangarooV2(t *testing.T) {
	noAns := "NO"
	yesAns := "YES"

	t.Run("yes when 0 3 4 2", func(t *testing.T) {
		got := kangaroo.KangarooV2(0, 3, 4, 2)
		assertEqual(t, got, yesAns)
	})

	t.Run("NO when 0 2 5 3", func(t *testing.T) {
		got := kangaroo.KangarooV2(0, 2, 5, 3)
		assertEqual(t, got, noAns)
	})
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got=(%s), want=(%s)", got, want)
	}
}
