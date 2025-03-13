// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// GammaRatio returns the ratio of products of Gamma functions, i.e.
//
//	                  n-1             m-1
//	GammaRatio(x, y) = ∏ Gamma(x[i]) / ∏ Gamma(y[j])
//	                  i=0             j=0
//
// where x = {x[0], ..., x[n-1]} and y = {y[0], ..., y[m-1]} are slices of length n
// and m respectively. The result is NaN if x and y contain a different number of
// infinite or NaN values.
//
// See http://mathworld.wolfram.com/GammaFunction.html for more information.
func GammaRatio(x, y []float64) float64 {
	lg, sg := LgammaRatio(x, y)
	return float64(sg) * math.Exp(lg)
}

// LgammaRatio returns the natural logarithm and sign of the ratio of products of Gamma functions, i.e.
//
//	                        n-1             m-1
//	LgammaRatio(x, y) = Log| ∏ Gamma(x[i]) / ∏ Gamma(y[j]) |
//	                        i=0             j=0
//
// where x = {x[0], ..., x[n-1]} and y = {y[0], ..., y[m-1]} are slices of length n
// and m respectively. The result is NaN if x and y contain a different number of
// infinite or NaN values.
//
// See http://mathworld.wolfram.com/GammaFunction.html for more information.
func LgammaRatio(x, y []float64) (float64, int) {
	x, y, nx, ny := removeCommonElements(x, y)

	s := 1
	res := 0.0
	npolex := 0
	npoley := 0
	for i := 0; i < nx; i++ {
		if xi := x[i]; math.IsInf(xi, 0) || math.IsNaN(xi) {
			return math.NaN(), 1
		} else if xi <= 0 && xi == math.Trunc(xi) {
			npolex++
			lg, _ := math.Lgamma(1 - xi)
			res -= lg
			s *= GammaSign(xi)
		} else {
			lg, sg := math.Lgamma(xi)
			res += lg
			s *= sg
		}
	}
	for i := 0; i < ny; i++ {
		if yi := y[i]; math.IsInf(yi, 0) || math.IsNaN(yi) {
			return math.NaN(), 1
		} else if yi <= 0 && yi == math.Trunc(yi) {
			npoley++
			lg, _ := math.Lgamma(1 - yi)
			res += lg
			s *= GammaSign(yi)
		} else {
			lg, sg := math.Lgamma(yi)
			res -= lg
			s *= sg
		}
	}

	if npolex < npoley {
		return math.Inf(-1), s
	}
	if npolex > npoley {
		return math.Inf(1), s
	}

	return res, s
}

/*
// lgrclose returns an estimate of ln(gamma(y+eps)/gamma(y)) for small eps.
func lgrclose(y, eps float64) float64 {
	return eps * (Digamma(y) + eps*(Trigamma(y)/2+eps*(Polygamma(2, y)/6+eps*(Polygamma(3, y)/24+eps*(Polygamma(4, y)/120+eps*(Polygamma(5, y)/720))))))
}
*/

// Poch returns the kth Pochhammer symbol of x, defined by
//
//	Poch(x, k) = Gamma(x+k) / Gamma(x)
//
// See http://mathworld.wolfram.com/PochhammerSymbol.html for more information.
func Poch(x, k float64) float64 {
	if math.IsInf(x, 0) {
		if x > 0 && k > 0 {
			return x
		}
		return 0
	}
	if math.IsInf(k, 0) {
		if x <= 0 && x == math.Trunc(x) {
			return 0
		}
		if k < 0 {
			return math.NaN()
		}
		return k * float64(GammaSign(x))
	}
	return GammaRatio([]float64{x + k}, []float64{x})
}
