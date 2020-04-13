// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

//
// Modified Bessel function of the first.
//

// In returns the order-n modified Bessel function of the first kind.
//
// Special cases are:
//  In(0, 0) = 1
//  In(±Inf, 0) = NaN
//  In(n > 0, 0) = 0
//  In(n < 0, 0) = 0 for integer n
//  In(n < 0, 0) = +Inf for non-integer n with even integer part
//  In(n < 0, 0) = -Inf for non-integer n with odd integer part
//  In(±Inf, ±Inf) = ±Inf
//  In(±Inf, x) = 0
//  In(n, ±Inf) = ±Inf
//
func In(n, x float64) float64 {
	if isNonPosInt(n) {
		n = -n
	}

	// special cases
	switch {
	case x == 0:
		switch {
		case n == 0:
			return 1
		case n > 0:
			return 0
		default: // n is negative non-integer
			in := math.Trunc(n)
			if math.Mod(in, 2) == 0 { // even integer part
				return +inf
			}
			return -inf // odd integer part
		}
	case math.IsInf(n, 0):
		if math.IsInf(x, 0) {
			return x
		}
		return 0
	case math.IsInf(x, 0):
		return x
	}
	return math.Pow(0.5*x, n) / math.Gamma(1+n) * hyp0F1(1+n, 0.25*x*x)
}
