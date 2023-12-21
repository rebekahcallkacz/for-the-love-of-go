package mytypes_test

import (
	"mytypes"
	"testing"
)
func TestTwice(t *testing.T) {
	t.Parallel()

	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()

	if want != got {
		t.Errorf("twice %d: want %d, got %d", input, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()

	input := mytypes.MyString("test")
	want := 4
	got := input.Len()

	if want != got {
		t.Errorf("%q: want %d, got %d", input, want, got)
	}
}