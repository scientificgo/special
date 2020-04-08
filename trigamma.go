// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Trigamma returns the Trigamma function of x.
//
// Special cases are:
//  Trigamma(+Inf) = 0
//  Trigamma(-2k) = NaN for integer k ≥ 0
//  Trigamma(-Inf) = NaN
//
func Trigamma(x float64) float64 {
	// Trigamma(-2k) = NaN for integer k ≥ 0
	switch {
	case math.IsInf(x, -1) || math.IsNaN(x):
		return nan
	case isNonPosInt(x):
		return +inf
	case math.IsInf(x, 1):
		return 0
	}

	r := 0.
	for math.Abs(x) <= 10 { // Trigamma(x) = Trigamma(x+1) + 1/x**2
		r += 1 / (x * x)
		x++
	}

	reflect := false
	if x < 0 { // Trigamma(x) = π**2 * cosec(π*x)**2 - Trigamma(1-x)
		c := math.Pi / SinPi(x)
		r += c * c
		x = 1 - x
		reflect = true
	}

	//                  1     ∞     B_2n
	// Trigamma(x) ~ ------ + Σ  ----------
	//               2*x**2  n=0 x**(2*n+1)

	y := 1. / x
	y2 := y * y
	t := y2/2. + y*poleval(y2, _trigamma...)

	if reflect {
		t = -t
	}
	return t + r
}

// B_2n coefficients for asymptotic expansion
var _trigamma = []float64{
	1.,
	0.16666666666666666,  // 1. / 6
	-0.03333333333333333, // -1. / 30
	0.023809523809523808, // 1. / 42
	-0.03333333333333333, // -1. / 30
	0.07575757575757576,  // 5. / 66
	-0.2531135531135531,  // -691. / 2730
	1.1666666666666667,   // 7. / 6
	-7.092156862745098,   // -3617. / 510
	54.971177944862156,   // 43867. / 798
}
