// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Hyp2F1 returns the 2,1-hypergeometric function
// of x with parameters a, b and c.
//
func hyp2F1(a, b, c, x float64) float64 {
	if a > b { // symmetric in a, b so let a ≤ b
		a, b = b, a
	}

	t := x * a * (b / c)
	s := t

	r := 0.
	for i := 2; math.Abs(t/s) > macheps; i++ {
		a++
		b++
		c++
		t *= x * a * (b / (c * float64(i)))

		// KBN summation
		tmp := s + t
		if math.Abs(s) > math.Abs(t) {
			r += (s - tmp) + t
		} else {
			r += (t - tmp) + s
		}
		s = tmp
	}
	s += 1 + r
	return s
}
