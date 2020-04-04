// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Digamma returns the Digamma function of x.
//
// Special cases are:
//  Digamma(1) = -Gamma
//  Digamma(+Inf) = +Inf
//  Digamma(k ≤ 0) = NaN for integer k
//  Digamma(-Inf) = NaN
//
func Digamma(x float64) float64 {
	// special cases
	switch {
	case isNonPosInt(x) || math.IsInf(x, -1) || math.IsNaN(x):
		return nan
	case x == 1:
		return -Euler
	case math.IsInf(x, 1):
		return x
	}

	digamma := 0.
	for math.Abs(x) <= 8 { // Digamma(x) = Digamma(x+1) - 1/x
		digamma -= 1 / x
		x++
	}

	if x < 0 { // Digamma(x) = Digamma(1-x) - π*cot(π*x)
		theta := math.Pi * math.Remainder(x, 1)
		digamma -= math.Pi / math.Tan(theta)
		x = 1 - x
	}

	//                        1    ∞  B_2n      1
	// Digamma(x) ~ Log(x) - --- + Σ  ---- * --------
	//                       2*x  n=1  2*n   x**(2*n)

	y := 1 / x
	y2 := y * y
	digamma += math.Log(x) - y/2 + y2*poleval(y2, _digamma...)

	return digamma
}

// coefficients for Digamma series expansion
var _digamma = []float64{
	-0.08333333333333333,  // -1. / 12
	0.008333333333333333,  // 1. / 120
	-0.003968253968253968, // -1. / 252
	0.004166666666666667,  // 1. / 240
	-0.007575757575757576, // -1. / 132
	0.021092796092796094,  // 691. / 32760
	-0.08333333333333333,  // -1. / 12
	0.4432598039215686,    // 3617. / 8160
	-3.0539543302701198,   // -43867. / 14364
}
