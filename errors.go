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
)

var (
	// When constructing a new Rating with New(), a non-negative rating value
	// must be provided
	errorNegativeValue = fmt.Errorf("rating value cannot be less than zero")

	// When constructing a new Rating with New(), a non-negative rating index
	// must be provided
	errorNegativeIndex = fmt.Errorf("rating index cannot be less than zero")

	// When constructing a new Rating with New(), a non-negative age
	// must be provided
	errorNegativeAge = fmt.Errorf("age cannot be less than zero")

	// The given result as float64 must be a sum of chess game results.
	// Since chess games end in wins (1.0 points), draws (0.5 points)
	// or losses (0.0 points), the given result must be a sum of these values
	errorInvalidResult = fmt.Errorf("result must end with .0 or .5")

	// The maximum number of points in the result is equal to the number
	// of opponents, since 1.0 is the maximum score in a chess game
	errorMorePointsThanGames = fmt.Errorf("too many points for too few opponents")
)
