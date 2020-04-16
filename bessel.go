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
//  In(+Inf, 0) = 0
//  In(-Inf, 0) = NaN
//  In(n > 0, 0) = 0
//  In(n < 0, 0) = 0 for integer n
//  In(n < 0, 0) = +Inf for non-integer n with even integer part
//  In(n < 0, 0) = -Inf for non-integer n with odd integer part
//  In(n, x < 0) = NaN for non-integer n
//  In(±Inf, +Inf) = +Inf
//  In(±Inf, -Inf) = NaN
//  In(±Inf, x) = 0
//  In(n, ±Inf) = ±Inf
//
func In(n, x float64) float64 {
	if isNonPosInt(n) {
		n = -n
	}

	// special cases
	switch {
	case x < 0:
		if !isInt(n) {
			return nan
		}
		if ni := math.Trunc(n); math.Mod(ni, 2) == 0 {
			return In(n, -x)
		}
		return -In(n, -x)
	case x == 0:
		switch {
		case math.IsInf(n, 0):
			if n < 0 {
				return nan
			}
			return 0
		case n > 0:
			return 0
		case n == 0:
			return 1
		default: // n is negative non-integer
			in := math.Trunc(n)
			if math.Mod(in, 2) == 0 { // even integer part
				return +inf
			}
			return -inf // odd integer part
		}
	case math.IsInf(n, 0):
		if math.IsInf(x, 1) {
			return x
		}
		return 0
	case math.IsInf(x, 1):
		return x
	}

	if n > 170 {
		// avoid overflow in Gamma function using Stirling's approx
		//
		//   (x/2)**n    ((e*x)/(2*n))**n
		//  ---------- ~ ----------------
		//  Gamma(n+1)    Sqrt(n) * S(n)
		//
		//  S(n) = Sqrt(2*π) * (1 + 1/(12*n) + 1/(288*n**2) + ...)
		//
		return math.Pow(0.5*x*math.E/n, n) / (stirling(n) * math.Sqrt(n)) * hyp0F1(1+n, 0.25*x*x)
	}
	return math.Pow(0.5*x, n) / math.Gamma(1+n) * hyp0F1(1+n, 0.25*x*x)
}
