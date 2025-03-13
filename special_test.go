// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"math"
)

const tol = 1e-10

var (
	nan = math.NaN()
	inf = math.Inf(1)
)

func equalFloat64(x float64, y float64) bool {
	if math.IsNaN(y) {
		return math.IsNaN(x)
	}

	if math.IsInf(y, 1) || math.IsInf(y, -1) {
		return x == y
	}

	if y == 0 {
		return math.Abs(x) < tol
	}

	return math.Abs((x-y)/y) < tol
}
