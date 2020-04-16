// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

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
