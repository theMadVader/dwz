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
		return nil, errorNegativeValue
	}
	if index < 0 {
		return nil, errorNegativeIndex
	}
	if age < 0 {
		return nil, errorNegativeAge
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
func (r *Rating) Next(result float64, oppRatings []int) (*Rating, error) {
	if !isValidResult(result) {
		return nil, errorInvalidResult
	}
	if result > float64(len(oppRatings)) {
		return nil, errorMorePointsThanGames
	}

	next := *r
	// R_n = R_o + 800 * (W - W_e) / (E + n)
	// R_n: new Rating
	// R_o: old Rating
	// W: Wins (Draw = 0.5)
	// W_e: expectedPoints
	// n: number of games
	// E: development coefficient, E = E_0 * f_B + S_Br
	We := r.expectedPoints(oppRatings)
	coeff := r.coeff(result, We)
	next.current = next.current + int(math.Round(800.0*(result-We)/(coeff+float64(len(oppRatings)))))

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
func (r *Rating) expectedValue(oppr int) float64 {
	diff := r.current - oppr
	return 1.0 / (1.0 + math.Pow10(-diff/400))
}

// expectedPoints is the sum of expected Values against a list of opponents
func (r *Rating) expectedPoints(opps []int) float64 {
	sum := 0.0
	for _, opr := range opps {
		sum += r.expectedValue(opr)
	}
	return sum
}

func (r *Rating) coeff(W, We float64) float64 {
	// E: development coefficient, E = E_0 * f_B + S_Br

	// J = {5, if age <= 20; 10, if 21 <= age <= 25; 15, if age > 25}
	var J float64
	switch {
	case r.age <= 20:
		J = 5.0
	case r.age > 20 && r.age <= 25:
		J = 10.0
	default:
		J = 15.0
	}

	// E_0 = (R_o / 1000)^4 + J
	E0 := math.Pow(float64(r.current)/1000.0, 4.0) + J

	// f_B = {0.5 <= (R_o / 2000) <= 1.0, if age <= 20 && W >= W_e; 1.0 in every other case}
	fB := 1.0
	if fb := float64(r.current) / 2000.0; W >= We && r.age <= 20 && fb < 1.0 {
		if fb < 0.5 {
			fB = 0.5
		} else {
			fB = fb
		}
	}

	// S_Br = {exp((1300-R_o)/150) - 1, if R_o < 1300 && W < W_e; 0.0 in every other case}
	SBr := 0.0
	if W < We && r.current < 1300 {
		SBr = math.Exp((1300.0 - float64(r.current)) / 150.0)
	}

	return math.Round(E0*fB + SBr)
}
