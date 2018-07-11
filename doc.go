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

// Package dwz implements the calculation of DWZ (Deutsche Wertungszahl) chess ratings.
//
// This package is intended for quick calculation of the update of a single rating,
// not for evaluating entire tournaments.
//
// Also: Games againts opponents without a DWZ
// rating cannot be evaluated and must be excluded when calling the Next() method.
package dwz
