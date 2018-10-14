// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package poly

import "math"

// LaguerreAL returns the nth associated Laguerre polynomial with parameter a at x.
//
// See http://mathworld.wolfram.com/AssociatedLaguerrePolynomial.html for more information.
func LaguerreAL(n int, a, x float64) float64 {
	switch {
	case math.IsNaN(a) || math.IsNaN(x) || n < 0:
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return 1 + a - x
	}

	scale := 1.0
	if a < 0 && a == math.Trunc(a) {
		// scale = Gamma(-n+|a|) / Gamma(-n)
		if -float64(n)-a <= 0 {
			// scale = (-1)**|a| n! / (n-|a|)!
			scale *= float64(1 - 2*(int(a)&1))
			lnfac, snfac := math.Lgamma(float64(n + 1))
			lanfac, sanfac := math.Lgamma(float64(n) + a + 1)
			scale *= float64(snfac*sanfac) * math.Exp(lanfac-lnfac)
		} else {
			return math.NaN()
		}
		n, a = n+int(a), -a
		scale *= math.Pow(x, a)
	}

	tmp := 1.0
	res := 1 + a - x
	for k := 1; k < n; k++ {
		p := a + float64(2*k) + 1 - x
		q := a + float64(k)
		r := float64(k + 1)
		res, tmp = (p*res-q*tmp)/r, res
	}
	return scale * res
}
