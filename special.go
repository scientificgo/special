// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

// Package special contains special mathematical functions
// and constants to supplement the standard math package.
package special // import "scientificgo.org/special"

import "math"

// Mathematical constants.
const (
	Euler = 0.577215664901532860606512090082402431042159335939923598805767234884867726777664670936947063291746749 // https://oeis.org/A001620
	LnPi  = 1.144729885849400174143427351353058711647294812915311571513623071472137769884826079783623270275489707 // https://oeis.org/A053510
	LnPhi = 0.481211825059603447497758913424368423135184334385660519661018168840163867608221774412009429122723474 // https://oeis.org/A002390
)

var (
	nan     = math.NaN()
	inf     = math.Inf(1)
	macheps = math.Nextafter(1, 2) - 1 // machine epsilon, or ε
)

// isInt returns true if x is exactly zero or within ε of a non-zero integer.
func isInt(x float64) (int float64, is bool) {
	r, f := math.Modf(x)
	int = r
	if f < 0 {
		r = -r
		f = -f
	}
	if r+macheps < 1 {
		is = f == 0
		return
	}
	is = f < macheps || f > 1-macheps
	return
}

// isNonPosInt returns true if x is a finite non-positive integer.
func isNonPosInt(x float64) (int float64, is bool) {
	int, is = isInt(x)
	is = is && int <= 0
	return
}

// chebeval evaluates the Chebyshev series defined by
//               n-1
//  y = cs[0]/2 + Σ cs[k] * T_k(x)
//               k=1
// where T_k(x) is the k-th Chebyshev polynomial of the
// first kind of x.
//
// Since the Chebyshev polynomials are only defined over
// (-1, 1), x must be transformed into that domain
// before calling this function. The transformation required
// depends on the interval used to calculate the coefficients.
//
// If the coefficients were calculated for (a, b):
//  x' = (2*x - b - a) / (b - a)
//
// If the coefficients were calculated for (1/b, 1/a):
//  x' = (2*a*b/x - b - a) / (b - a)
//
func chebeval(x float64, cs ...float64) float64 {
	var b0, b1, b2 float64
	n := len(cs)
	b0 = cs[n-1]
	for i := n - 2; i >= 0; i-- {
		b1, b2 = b0, b1
		b0 = cs[i] + 2*x*b1 - b2
	}
	return (b0 - b2) / 2
}

// poleval evaluates the polynomial cs[0] + cs[1] * x + ... + cs[n] * x**n
func poleval(x float64, cs ...float64) (p float64) {
	for i := len(cs) - 1; i >= 0; i-- {
		p = cs[i] + p*x
	}
	return
}

// factorialseries returns the series
//      n
//  y = Σ k! * x**k
//     k=0
// where n is the smallest positive integer such that
// |(n+1) * x| > 1 or |n! * x**n / y| < ε, where ε is
// the machine precision.
//
// This cutoff is necessary since the series is divergent for all
// x in the limit n → ∞.
func factorialseries(x float64) float64 {
	s := 0.

	for i, t, ti := 1, 1., x; math.Abs(ti) < 1 && math.Abs(t/s) > macheps; i++ {
		t *= ti
		s += t
		ti = x * float64(i+1)
	}

	return s + 1
}
