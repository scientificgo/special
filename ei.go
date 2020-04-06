// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Ei returns the exponential integral of x.
//
// Special cases are:
//  Ei(0) = -Inf
//  Ei(+inf) = +Inf
//  Ei(-inf) = 0
//
func Ei(x float64) float64 {
	// special cases
	switch {
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, -1) || x < -705:
		return 0
	case math.IsInf(x, 1) || x > 716:
		return math.Inf(1)
	case x == 0:
		return math.Inf(-1)
	}

	if x < -1 { // continued fraction
		ai := func(i int) float64 {
			if i == 1 {
				return 1
			}
			return float64(i / 2)
		}
		bi := func(i int) float64 {
			if i%2 == 0 {
				return 1
			}
			return -x
		}
		return -math.Exp(x) * gcf(ai, bi)
	}

	if x >= 40 { // asymptotic expansion

		//         e**x  ∞   k!
		// Ei(x) ~ ----  Σ  ----
		//          x   k=0 x**k

		y := 1 / x
		e := math.Exp(x / 2) // avoid overflow with e**x = (e**x/2)**2
		return y * factorialseries(y) * e * e
	}

	// power series

	//          ∞    x**k
	// Ei(x) ~  Σ   ------
	//         k=1  k * k!

	s := 0.
	t := 1.
	for i := 1; math.Abs(t/s) > macheps; i++ {
		t *= x / float64(i)
		s += t / float64(i)
	}
	return Euler + math.Log(math.Abs(x)) + s
}
