// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// ChebyshevU returns the nth Chebyshev polynomial of the second kind at x.
//
// See http://mathworld.wolfram.com/ChebyshevPolynomialoftheSecondKind.html for more information.
func ChebyshevU(n int, x float64) float64 {
	s := 1
	if n <= -2 {
		s = -1
		n = -n - 2
	}
	if x < 0 {
		x = -x
		s *= 1 - 2*(n&1)
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
		return float64(s * (n + 1))
	case n == -1:
		return 0
	case n == 0:
		return float64(s)
	case n == 1:
		return float64(s) * 2 * x
	}

	const nlarge = 55

	var res float64
	if n <= nlarge {
		tmp := 1.0
		x *= 2
		res = x
		for k := 2; k <= n; k++ {
			res, tmp = x*res-tmp, res
		}
	} else {
		// For large n, use the trigonometric definitions.
		if x < 1 {
			t := math.Acos(x)
			res = math.Sin(float64(n+1)*t) / math.Sin(t)
		} else {
			t := math.Acosh(x)
			res = math.Sinh(float64(n+1)*t) / math.Sinh(t)
		}
	}
	return float64(s) * res
}
