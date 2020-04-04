// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// AGM returns the arithmetic-geometric mean of x and y.
//
// Special cases are:
//  AGM(x, 0) = 0 for finite x
//  AGM(±inf, 0) = NaN
//  AGM(x > 0, +Inf) = +Inf
//  AGM(x < 0, y) = NaN for non-zero y
//
func AGM(x, y float64) float64 {
	if x < y {
		x, y = y, x
	}

	// special cases
	switch {
	case math.IsNaN(x) || math.IsNaN(y):
		return nan
	case y <= 0:
		if x == 0 || (y == 0 && !math.IsInf(x, 1)) {
			return 0
		}
		return nan
	case x == y:
		return x
	}

	// iterate until x and y converge
	for 1-y/x > macheps { // arithmetic ≥ geometric so y/x ≤ 1
		if xy := x * y; math.IsInf(xy, 1) { // x*y overflows
			x, y = (x+y)/2, math.Sqrt(x)*math.Sqrt(y)
		} else {
			x, y = (x+y)/2, math.Sqrt(xy)
		}
	}
	return x
}
