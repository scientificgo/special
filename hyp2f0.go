// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// hyp2F0 returns the 2,0-hypergeometric function of x
// with parameters a and b.
//
// This series formally diverges for all non-zero x, but it can
// provide useful values as an aymptotic expansion if truncated
// after n terms when either:
//  1. |t_{n+1}/t_{n}| = |x * (a+n) * (b+n) / (n+1)| ≥ 1, or
//  2. |t_{n}/s_{n}| < ε
//  3. a or b is a non-positive integer, then the series
//     becomes a polynomial of order |a| or |b|.
// where t_{k} is the k-th term in the series and s_{k} is the
// partial sum after k terms.
//
func hyp2F0(a, b, x float64) float64 {
	if b < a {
		a, b = b, a
	}

	t := x * a * b
	s := t
	r := 0.

	ispoly := isNonPosInt(a) || isNonPosInt(b)

	for i := 2; math.Abs(t/s) > macheps || (ispoly && a*b != 0); i++ {
		a++
		b++

		if ti := x * a * b / float64(i); math.Abs(ti) < 1 || ispoly {
			t *= ti
		} else {
			break
		}

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
