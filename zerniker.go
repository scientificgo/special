// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// ZernikeR returns the nth Zernike polynomial with parameter m at x.
//
// See http://mathworld.wolfram.com/ZernikePolynomial.html for more information.
func ZernikeR(n, m int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0 || m < 0:
		return math.NaN()
	case (n-m)&1 == 1 || n < m:
		return 0
	default:
		s := (n - m) / 2
		s = 1 - 2*(s&1)
		return float64(s) * math.Pow(x, float64(m)) * JacobiP((n-m)/2, float64(m), 0, 1-2*x*x)
	}
}
