package dwz

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	r := New(1234, 5)

	want := "1234-5"
	got := fmt.Sprintf("%v", r)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestNextLengthError(t *testing.T) {
	r := New(1234, 5)

	want := fmt.Errorf("too many points for too few opponents")
	_, got := r.Next(3.5, []Rating{New(1000, 1)})
	if got.Error() != want.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}

func TestInvalidResultError(t *testing.T) {
	r := New(1234, 5)

	want := fmt.Errorf("result must end with .0 or .5")
	_, got := r.Next(0.12345, []Rating{New(1000, 1)})
	if got.Error() != want.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}
