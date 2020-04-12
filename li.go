// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Li returns the logarithmic integral of x.
//
// Special cases are:
//  Li(0) = 0
//  Li(1) = -Inf
//  Li(+Inf) = +Inf
//  Li(x < 0) = NaN
//
func Li(x float64) float64 {
	switch {
	case math.IsInf(x, 1):
		return x
	case x < 0 || math.IsNaN(x):
		return nan
	}

	l := math.Log(x)

	if l <= 40 {
		return Ei(math.Log(x))
	}

	// For large x, use the asymptotic series for Ei
	// with the substitution x -> ln(x) to avoid additional
	// rounding errors.

	//           x    ∞     k!
	// Li(x) ~ -----  Σ  --------
	//         ln(x) k=0 ln(x)**k

	y := 1 / l
	return x * y * hyp2F0(1, 1, y)
}
