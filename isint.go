// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// isInt returns true if x is exactly zero or within ε of a non-zero integer.
func isInt(x float64) bool {
	r, f := math.Modf(x)
	if f < 0 {
		r = -r
		f = -f
	}
	if r+macheps < 1 {
		return f == 0
	}
	return f < macheps || f > 1-macheps
}

// isNonPosInt returns true if x is a finite non-positive integer.
func isNonPosInt(x float64) bool { return x <= 0 && isInt(x) }
