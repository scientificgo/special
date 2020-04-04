// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// EllipticK returns the complete elliptical integral of
// the first kind with parameter x = k**2.
//
// Special cases are:
//  EllipticK(1) = +Inf
//  EllipticK(0) = π/2
//  EllipticK(-Inf) = 0
//  EllipticK(x > 1) = NaN
//
func EllipticK(x float64) float64 {
	// apply https://dlmf.nist.gov/19.8#E5

	//                     π/2
	//  EllipticK(x) = -----------
	//                 AGM(1, 1-x)

	x = math.Sqrt(1 - x)
	return math.Pi / AGM(x, 1) / 2
}
