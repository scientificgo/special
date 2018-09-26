// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// GammaRegP returns the regularised lower incomplete gamma function,
// defined by
//
//                                   x
//  GammaRegP(a, x) = [1 / Gamma(a)] ∫ dt Exp(-t) * t**(a-1)
//                                  t=0
//
// The regularised lower incomplete gamma function has a series representation
//                                 ∞
//  GammaRegP(a, x) = Exp(-x) a**x ∑ x**k / Gamma(a+k+1)
//                                k=0
//
// and also satisfies the identity
//
//  GammaRegP(a, x) + GammaRegQ(a, x) = 1
//
// where GammaRegQ is the regularised upper incomplete gamma function.
//
// See http://mathworld.wolfram.com/RegularizedGammaFunction.html
// for more information.
func GammaRegP(a, x float64) float64 {
	// Special cases.
	switch {
	case x < 0 || math.IsNaN(x) || math.IsNaN(a) || math.IsInf(a, -1):
		return math.NaN()
	case x == 0:
		return 0
	case math.IsInf(a, 1):
		if math.IsInf(x, 1) {
			return 0.5
		}
		return 0
	case math.IsInf(x, 1) || (a <= 0 && a == math.Trunc(a)):
		return 1
	case a == 1:
		return 1 - math.Exp(-x)
	}

	//
	// Use gammaQ as primary function and calculate using
	// the continued fraction representation.
	//

	if x > a && !(x < 2 && a > -10) {
		return 1 - gammaQcf(a, x)
	}

	//
	// Calculate with the explicit power series when the continued
	// fraction for gammaQ can't be used.
	//

	return gammaPseries(a, x)
}

// GammaRegQ returns the regularised upper incomplete gamma function,
// defined by
//
//                                   ∞
//  GammaRegQ(a, x) = [1 / Gamma(a)] ∫ dt Exp(-t) * t**(a-1)
//                                  t=x
//
// GammaRegQ also satisfies the identity
//
//  GammaRegP(a, x) + GammaRegQ(a, x) = 1
//
// where GammaRegP is the regularised lower incomplete gamma function.
//
// See http://mathworld.wolfram.com/RegularizedGammaFunction.html
// for more information.
func GammaRegQ(a, x float64) float64 {
	// Special cases.
	switch {
	case x < 0 || math.IsNaN(x) || math.IsNaN(a) || math.IsInf(a, -1):
		return math.NaN()
	case x == 0:
		return 1
	case math.IsInf(a, 1):
		if math.IsInf(x, 1) {
			return 0.5
		}
		return 1
	case math.IsInf(x, 1) || (a <= 0 && a == math.Trunc(a)):
		return 0
	case a == 1:
		return math.Exp(-x)
	}

	//
	// Calculate with the explicit power series when the continued
	// fraction for gammaQ can't be used.
	//

	if x < a || (x < 2 && a > -10) {
		return 1 - gammaPseries(a, x)
	}

	//
	// Use the continued fraction representation to calculate efficiently
	// whenever possible.
	//

	return gammaQcf(a, x)
}

// gammaPseries returns GammaRegP using the hypergeometric series definition
func gammaPseries(a, x float64) float64 {
	const (
		maxiter = 2000
		rtol    = 1e-20
	)

	// Note that each term is proportional to the last by a factor of x/(a+k)
	// and hence it is unnecessary to calculate Γ(a+k+1) at each iteration.
	res := 1.0
	for k, tmp := 1, 1.0; k < maxiter && math.Abs(tmp/res) > rtol; k++ {
		tmp *= x / (a + float64(k))
		res += tmp
	}

	lga1, sga1 := math.Lgamma(a + 1)
	res *= float64(sga1) * math.Exp(a*math.Log(x)-x-lga1)
	if a > 0 {
		res = math.Min(res, 1)
	}
	return res
}

// gammaQcf returns GammaRegQ using a continued fraction.
func gammaQcf(a, x float64) float64 {
	lga, sga := math.Lgamma(a)
	s := math.Copysign(1, x)
	lx := math.Log(math.Abs(x))
	xma := x - a

	d := cfgammaQdepth(a, x)
	cf := xma + float64(d<<1+1)
	for i := d; i > 0; i-- {
		j := (i-1)<<1 + 1
		ai := float64(i) * (a - float64(i))
		bj := xma + float64(j)
		cf = bj + ai/cf
	}
	return s * float64(sga) * math.Exp(a*lx-x-lga) / cf
}

// cfdepth returns the depth required for convergence for the continued fraction for GammaRegQ.
func cfgammaQdepth(a, x float64) int {
	switch y := x / a; {
	case y > 1.5:
		return 10
	case y > 1.3:
		return 20
	case y > 1.1:
		return 40
	case y > 1.05:
		return 50
	default:
		return 100
	}
}

// GammaIncL returns the lower incomplete gamma function,
// defined by
//
//                    x
//  GammaIncL(a, x) = ∫ dt Exp(-t) * t**(a-1)
//                   t=0
//
// GammaIncL also satisfies the identity
//
//  GammaIncL(a, x) + GammaIncU(a, x) = Gamma(a)
//
// where GammaIncU is the upper incomplete gamma function.
//
// See http://mathworld.wolfram.com/IncompleteGammaFunction.html
// for more information.
func GammaIncL(a, x float64) float64 {
	lga, sga := math.Lgamma(a)
	gp := GammaRegP(a, x)
	if gp < 0 {
		gp = -gp
		sga = -sga
	}
	return float64(sga) * math.Exp(lga+math.Log(gp))
}

// GammaIncU returns the lower incomplete gamma function,
// defined by
//
//                    ∞
//  GammaIncU(a, x) = ∫ dt Exp(-t) * t**(a-1)
//                   t=x
//
// GammaIncU also satisfies the identity
//
//  GammaIncL(a, x) + GammaIncU(a, x) = Gamma(a)
//
// where GammaIncL is the lower incomplete gamma function.
//
// See http://mathworld.wolfram.com/IncompleteGammaFunction.html
// for more information.
func GammaIncU(a, x float64) float64 {
	switch {
	case a <= 0 && math.Trunc(a) == a:
		return math.Pow(x, a) * En(int(1-a), x)
	default:
		lga, sga := math.Lgamma(a)
		gq := GammaRegQ(a, x)
		if gq < 0 {
			gq = -gq
			sga = -sga
		}
		return float64(sga) * math.Exp(lga+math.Log(gq))
	}
}
