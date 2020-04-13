// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// hyp0F1 returns the 0,1-hypergeometric function of x
// with parameter b.
//
func hyp0F1(b, x float64) float64 {
	t := x / b
	s := t
	r := 0.
	for i := 2; math.Abs(t/s) > macheps; i++ {
		b++
		t *= x / (b * float64(i))

		// KBN summation
		tmp := s + t
		if math.Abs(s) >= math.Abs(t) {
			r += (s - tmp) + t
		} else {
			r += (t - tmp) + s
		}
		s = tmp
	}
	s += 1 + r
	return s
}
