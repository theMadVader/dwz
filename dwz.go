// Copyright 2018 Matthias Vedder
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dwz

import (
	"fmt"
	"math"
)

// Rating encapsulates all necessary data for a DWZ rating.
type Rating struct {
	current int
	index   int
	age     int
}

func (r Rating) String() string {
	return fmt.Sprintf("%d-%d", r.current, r.index)
}

// New constructs a Rating based on the given inputs.
//
// New returns a pointer to the constructed Rating and an error.
// The error will be non-nil if any parameter is negative.
func New(current, index, age int) (*Rating, error) {
	if current < 0 {
		return nil, fmt.Errorf("rating value cannot be less than zero")
	}
	if index < 0 {
		return nil, fmt.Errorf("rating index cannot be less than zero")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be less than zero")
	}

	return &Rating{current, index, age}, nil
}

// Current returns the current amount DWZ rating.
//
// This is the first number of "1234-12"
func (r *Rating) Current() int {
	return r.current
}

// Index returns the current index of the DWZ rating.
//
// This is the second number of "1234-12"
func (r *Rating) Index() int {
	return r.index
}

// Next calculates a new Rating, based on the results
//
// The result parameter must be a float64 that ends with .0 or .5
// and it must be smaller or equal to the length of oppRatings,
// otherwise there will be a non-nil error
func (r *Rating) Next(result float64, oppRatings []Rating) (*Rating, error) {
	if !isValidResult(result) {
		return nil, fmt.Errorf("result must end with .0 or .5")
	}
	if result > float64(len(oppRatings)) {
		return nil, fmt.Errorf("too many points for too few opponents")
	}

	next := *r
	// TODO: calculate next.current
	// R_n = R_o + 800 * (W - W_c) / (E + n)
	// R_n: new Rating
	// R_o: old Rating
	// W: Wins (Draw = 0.5)
	// W_c: expectedPoints
	// E: development coefficient
	// n: number of games
	next.index++
	return &next, nil
}

// isValidResult checks whether or not result ends in a .0 or a .5
func isValidResult(r float64) bool {
	r *= 2.0
	return r == math.Round(r)
}

// expectedValue tells how high the propability is for the first player to
// defeat the second player, based on the rating difference.
func expectedValue(myr, oppr Rating) float64 {
	diff := myr.current - oppr.current
	return 1.0 / (1.0 + math.Pow10(-diff/400))
}

// expectedPoints is the sum of expected Values against a list of opponents
func expectedPoints(myr Rating, opps []Rating) float64 {
	sum := 0.0
	for _, opr := range opps {
		sum += expectedValue(myr, opr)
	}
	return sum
}
