// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package poly

import "math"

// LegendreAP returns the nth associated Legendre polynomial of the first kind
// with parameter m at x.
//
// See http://mathworld.wolfram.com/AssociatedLegendrePolynomial.html for more information.
func LegendreAP(n, m int, x float64) float64 {
	// Calculate P(|m|, |m|, x) and then use the recurrence
	// relation to get P(n, |m|, x). Finally, if m < 0, use the
	// reflection formula between m and -m to get P(n, m, x).

	if n < 0 {
		n = -n - 1
	}

	sign := 1 - 2*(m&1)
	switch {
	case sign < 0 && math.Abs(x) > 1:
		return math.NaN()
	case m > n || m < -n:
		return 0
	}

	// Calculate the prefactor for cases with m < 0 using the reflection formula.
	negmfac := 1.0
	if m < 0 {
		m = -m
		negmfac = float64(sign) * math.Gamma(float64(n-m+1)) / math.Gamma(float64(n+m+1))
	}

	// P(|m|, |m|, x)
	res := float64(sign) * math.Pow(1-x*x, float64(m)/2) * math.Gamma(float64(m)+0.5) / math.SqrtPi * float64(int(1<<uint(m)))
	if n > m {
		// Use recurrance formula to go from P(|m|, |m|, x) to P(n, |m|, x),
		// using a special case to get P(|m|+1, |m|, x).
		tmp := res
		res = x * float64(2*m+1) * res
		for k := m + 1; k < n; k++ {
			res, tmp = (x*float64(2*k+1)*res-float64(k+m)*tmp)/float64(k-m+1), res
		}
	}
	return negmfac * res
}
