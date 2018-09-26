// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

func Chi(x float64) float64 {
	return Ei(x) - Shi(x)
}

func Shi(x float64) float64 {
	if x == 0 {
		return 0
	}

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	if x < 35 {
		z := x * x
		a := 1.0
		s := 1.0
		k := 2.0
		for math.Abs(a/s) > 1e-16 {
			a *= z / k
			k++
			a /= k
			s += a / k
			k++
		}
		return float64(sign) * s * x
	}
	return float64(sign) * Ei(x) / 2
}
