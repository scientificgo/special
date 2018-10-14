// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package poly

import "math"

// ChebyshevT returns the nth Chebyshev polynomial of the first kind at x.
//
// See http://mathworld.wolfram.com/ChebyshevPolynomialoftheFirstKind.html for more information.
func ChebyshevT(n int, x float64) float64 {
	s := 1
	if n < 0 {
		n = -n
	}
	if x < 0 {
		x = -x
		s = 1 - 2*(n&1)
	}

	switch {
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, 1):
		return float64(s) * x
	case x == 0:
		if n%2 == 1 {
			return 0
		}
		return float64(s) * math.Cos(math.Pi*float64(n)/2)
	case x == 1:
		return float64(s)
	case n == 0:
		return float64(s)
	case n == 1:
		return float64(s) * x
	}

	const nlarge = 45

	var res float64
	if n <= nlarge {
		tmp := 1.0
		res = x
		x *= 2
		for k := 2; k <= n; k++ {
			res, tmp = x*res-tmp, res
		}
	} else {
		// For large n, use the trigonometric definitions.
		if math.Abs(x) < 1 {
			res = math.Cos(float64(n) * math.Acos(x))
		} else {
			res = math.Cosh(float64(n) * math.Acosh(x))
		}
	}
	return float64(s) * res
}
