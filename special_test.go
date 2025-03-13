// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"math"
)

const tol = 1e-10

var (
	nan = math.NaN()
	inf = math.Inf(1)
)

func equalFloat64(x float64, y float64) bool {
	if math.IsNaN(y) {
		return math.IsNaN(x)
	}

	if math.IsInf(y, 1) || math.IsInf(y, -1) {
		return x == y
	}

	if y == 0 {
		return math.Abs(x) < tol
	}

	return math.Abs((x-y)/y) < tol

}

// func equalComplex128(x complex128, y complex128) (bool, float64) {
// 	if x == y || (cmplx.IsNaN(x) && cmplx.IsNaN(y)) {
// 		return true, 0
// 	}
// 	if cmplx.Abs(y) == 0 {
// 		diff := cmplx.Abs(x)
// 		return diff < tol, diff
// 	} else {
// 		diff := cmplx.Abs((x - y) / y)
// 		return diff < tol, diff
// 	}
// }
