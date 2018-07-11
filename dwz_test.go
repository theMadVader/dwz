package dwz

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	r := Rating{1234, 5, 20}

	want := "1234-5"
	got := fmt.Sprintf("%v", r)
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestNew(t *testing.T) {
	want := Rating{1234, 5, 20}
	got, err := New(1234, 5, 20)
	if err != nil {
		t.Error(err)
	}
	if *got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

	var wanterr error

	_, err = New(-4, 5, 20)
	wanterr = fmt.Errorf("rating value cannot be less than zero")
	if wanterr.Error() != err.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}

	_, err = New(1234, -1, 20)
	wanterr = fmt.Errorf("rating index cannot be less than zero")
	if wanterr.Error() != err.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}

	_, err = New(1234, 5, -2)
	wanterr = fmt.Errorf("age cannot be less than zero")
	if wanterr.Error() != err.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}

func TestNextLengthError(t *testing.T) {
	r := Rating{1234, 5, 20}

	want := fmt.Errorf("too many points for too few opponents")
	_, got := r.Next(3.5, []Rating{Rating{1000, 1, 20}})
	if got.Error() != want.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}

func TestInvalidResultError(t *testing.T) {
	r := Rating{1234, 5, 20}

	want := fmt.Errorf("result must end with .0 or .5")
	_, got := r.Next(0.12345, []Rating{Rating{1000, 1, 20}})
	if got.Error() != want.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}
