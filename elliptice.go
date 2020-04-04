// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// EllipticE returns the complete elliptical integral of
// the second kind with paramter x = k**2.
//
// Special cases are:
//  EllipticE(1) = 1
//  EllipticE(0) = π/2
//  EllipticE(-Inf) = 0
//  EllipticE(x > 1) = NaN
//
func EllipticE(x float64) float64 {
	// special cases
	switch {
	case x >= 1 || math.IsNaN(x):
		if x == 1 {
			return 1
		}
		return nan
	case math.IsInf(x, -1):
		return +inf
	}

	switch {
	case x <= -1e6: // series expansion for large negative x

		// Well approximated by the first few terms of
		// http://functions.wolfram.com/08.01.06.0030.01

		x = -x
		s := math.Sqrt(x)
		l := 4*math.Ln2 + math.Log(x)

		c0 := 1.
		c1 := 0.25 + 0.25*l
		c2 := 0.046875 - 0.03125*l
		return s * poleval(1/x, c0, c1, c2)

	case x >= 0.95: // x is near 1, the singular point

		// Use the asymptotic series https://dlmf.nist.gov/19.12#E2
		//
		//                           ∞
		//  EllipticE(x) = 1 + y/2 * Σ c[k] * (l + d[k] - r[k])
		//                          k=0
		//
		// where y = 1-x, l = -Log(y)/2, and
		//
		//  c[k] = Poch(1/2, k) * Poch(3/2, k) / Poch(2, k) / k! * y**k
		//  d[k] = Digamma(k+1) - Digamma(k+1/2)
		//  r[k] = 1 / (2*k+1) / (2*k+2)
		//
		// where Poch(z, k) is a Pochhammer symbol.
		//

		y := 1 - x
		l := -math.Log(y) / 2

		c := 1.
		d := 1.3862943611198906 // d[0] = Digamma(1) - Digamma(1/2)
		s := 0.
		for i := 0.; c/s > macheps; i++ {
			r := 1 / (2*i + 1) / (2*i + 2)
			s += c * (l + d - r)

			// c[k+1] = c[k] * ((2*k+1)*(2*k+3)) / (4*(k+1)*(k+2)) * y
			c *= ((2*i + 1) * (2*i + 3)) / (4 * (i + 1) * (i + 2)) * y

			// d[k+1] = d[k] - 2 * r[k]
			d -= 2 * r
		}
		return 1 + s*y/2

	default: // arithmetic-geometric mean algorithm

		// The result is given by https://dlmf.nist.gov/19.8#E6
		//
		//                     π/2            ∞
		//  EllipticE(x) = ----------- * (1 - Σ c[i]**2 * 2**(i-1))
		//                 AGM(1, 1-x)       i=0
		//
		// where a[i] and g[i] the i-th iterations of AGM(1, 1-x)
		// and c[i+1] = sqrt(a[i+1]**2 - g[i+1]**2) = (a[i] - g[i])/2.
		//

		// initial values
		a := math.Sqrt(1 - x)
		g := 1.

		if a < g {
			a, g = g, a
		}

		s := 1 - x/2 // c[0]**2 = 1 - (1 - x) = x
		for k := 1; 1-g/a > macheps; k *= 2 {
			c := (a - g) / 2
			s -= c * c * float64(k)
			a, g = (a+g)/2, math.Sqrt(a*g)
		}
		return math.Pi / (a + g) * s
	}
}
