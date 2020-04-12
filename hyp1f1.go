// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// hyp1F1 returns the 1,1-hypergeometric function
// of the first kind of x with parameters a and b.
//
// Special cases are:
//  Hyp1F1(a, b, 0) = 1
//  Hyp1F1(a, a, x) = Exp(x)
//  Hyp1F1(a > 0, b ≤ 0, x) = NaN for integer b
//  Hyp1F1(a ≤ 0, b ≤ 0, x) = NaN for for integer a, b with -a < -b
//
func hyp1F1(a, b, x float64) float64 {
	switch {
	case x == 0:
		return 1
	case a == b:
		return math.Exp(x)
	case isNonPosInt(b):
		if !isNonPosInt(a) {
			return nan
		}
		if a < b {
			return nan
		}
	}

	if x < -1 && !isNonPosInt(a) {
		return math.Exp(x) * hyp1F1(b-a, b, -x)
	}

	if a == 1 {
		return hyp1F1a1(b, x)
	}

	t := x * (a / b)
	s := t
	r := 0.
	for i := 2; math.Abs(t/s) > macheps; i++ {
		a++
		b++
		t *= x * a / (b * float64(i))

		// KBN summation
		tmp := s + t
		if math.Abs(s/t) >= math.Abs(t) {
			r += (s - tmp) + t
		} else {
			r += (t - tmp) + s
		}
		s = tmp
	}
	s += 1 + r
	return s
}

// hyp1F1a1 evaluates hyp1F1(a=1, b, x).
func hyp1F1a1(b, x float64) float64 {
	t := x / b
	s := t
	r := 0.
	for math.Abs(t/s) > macheps {
		b++
		t *= x / b

		// KBN summation
		tmp := s + t
		if math.Abs(s/t) >= math.Abs(t) {
			r += (s - tmp) + t
		} else {
			r += (t - tmp) + s
		}
		s = tmp
	}
	s += 1 + r
	return s
}
