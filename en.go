// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// En returns the order-n generalised exponential integral of x.
//
// Special cases are:
//  En(0, x) = e**(-x) / x
//  En(1, x) = -Ei(-x)
//  En(n, 0) = 1 / (n-1)
//  En(n, +Inf) = 0
//  En(n > 0, x < 0) = NaN
//  En(n < 0, x) = NaN
//
func En(n int, x float64) float64 {
	// special cases.
	switch {
	case math.IsNaN(x) || n < 0 || (n == 0 && x == 0) || (x < 0 && n > 1):
		return math.NaN()
	case math.IsInf(x, 1):
		return 0
	case n == 0:
		return math.Exp(-x) / x
	case n == 1:
		return -Ei(-x)
	case x == 0:
		return 1 / float64(n-1)
	case math.Abs(x) < 1:
		// when x is small, use recurrence
		// En(n+1, x) = (Exp(-x) - x*En(n, x)) / n
		a := math.Exp(-x)
		s := Ei(-x) // -En(1, x)
		s = a + x*s
		for i := 2; i < n; i++ {
			s = (a - x*s) / float64(i)
		}
		return s
	}

	ai := func(i int) float64 {
		if i%2 == 0 {
			return float64((i-1)/2 + n)
		}
		return float64(i / 2)
	}
	bi := func(i int) float64 {
		if i%2 == 0 {
			return 1
		}
		return x
	}
	return math.Exp(-x) * GCF(ai, bi)
}
