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
	wanterr = errorNegativeValue
	if wanterr.Error() != err.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}

	_, err = New(1234, -1, 20)
	wanterr = errorNegativeIndex
	if wanterr.Error() != err.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}

	_, err = New(1234, 5, -2)
	wanterr = errorNegativeAge
	if wanterr.Error() != err.Error() {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}

func TestNextLengthError(t *testing.T) {
	r := Rating{1234, 5, 20}

	wanterr := errorMorePointsThanGames
	_, err := r.Next(3.5, []int{1000})
	if err.Error() != wanterr.Error() {
		t.Errorf("\n got: %v\nwant: %v", err, wanterr)
	}
}

func TestInvalidResultError(t *testing.T) {
	r := Rating{1234, 5, 20}

	wanterr := fmt.Errorf("result must end with .0 or .5")
	_, err := r.Next(0.12345, []int{1000})
	if err.Error() != wanterr.Error() {
		t.Errorf("\n got: %v\nwant: %v", err, wanterr)
	}
}
