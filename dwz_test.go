package dwz

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	r := Rating{1234, 5, 20}

	want := "1234-5 (20)"
	got := fmt.Sprintf("%v (%d)", r, r.age)
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

func TestNextInvalidResultError(t *testing.T) {
	r := Rating{1234, 5, 20}

	wanterr := fmt.Errorf("result must end with .0 or .5")
	_, err := r.Next(0.12345, []int{1000})
	if err.Error() != wanterr.Error() {
		t.Errorf("\n got: %v\nwant: %v", err, wanterr)
	}
}

type dwzTestData struct {
	oldRating  Rating
	result     float64
	oppRatings []int
	newRating  Rating
}

var (
	dwzTestSet = []dwzTestData{
		dwzTestData{
			Rating{1566, 30, 34},
			2.5,
			[]int{1619, 1524, 1389, 1688, 1808, 1679},
			Rating{1563, 31, 34},
		},
		dwzTestData{
			Rating{1214, 3, 14},
			2.5,
			[]int{1682, 1447, 1495, 1370, 1563, 768},
			Rating{1271, 4, 14},
		},
	}
)

const (
	maxDeviation = 1
)

func TestNextNoError(t *testing.T) {
	for _, testdat := range dwzTestSet {
		curr := testdat.oldRating

		want := testdat.newRating
		got, _ := curr.Next(testdat.result, testdat.oppRatings)
		if *got != want {
			// Allow rating value to be off by 1
			if abs(got.Current()-want.Current()) > maxDeviation {
				t.Errorf("got: %v, want: %v", got, want)
			}
		}
	}
}

func TestExpectedPoints(t *testing.T) {
	want := 2.592
	got := dwzTestSet[0].oldRating.expectedPoints(dwzTestSet[0].oppRatings)
	if !almostEqual(got, want, 1e-3) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func almostEqual(x, y, eps float64) bool {
	return math.Abs(x-y) < eps
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
