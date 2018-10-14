// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package poly

import "math"

// HermiteH returns the nth unnormalised, or physics, Hermite polynomial, which is
// related to the normalised Hermite polynomial by
//
//  H(n, x) = √2**n He(n, x√2)
//
// where He is the normalised Hermite polynomial.
//
// See http://mathworld.wolfram.com/HermitePolynomial.html for more information.
func HermiteH(n int, x float64) float64 {
	return math.Exp(float64(n)*math.Ln2/2) * HermiteHe(n, math.Sqrt2*x)
}

// HermiteHe returns the nth normalised Hermite polynomial, which is
// related to the "physics" Hermite polynomial by
//
//  H(n, x) = √2**n He(n, x√2)
//
// where H is the unnormalised, or physics, Hermite polynomial.
//
// See http://mathworld.wolfram.com/HermitePolynomial.html for more information.
func HermiteHe(n int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0:
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return x
	}

	tmp := 1.0
	res := x
	for k := 1; k < n; k++ {
		res, tmp = x*res-float64(k)*tmp, res
	}
	return res
}
