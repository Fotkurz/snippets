package testutils

import "testing"

func AssertEqual(t *testing.T, want, got interface{}) {
	if got != want {
		t.Errorf("expected=(%d), got=(%d)", want, got)
	}
}

func AssertNil(t *testing.T, v interface{}) {
	if v != nil {
		t.Errorf("expected nil but got=%v", v)
	}
}
