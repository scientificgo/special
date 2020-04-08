// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Beta returns the Beta function of x and y.
//
// Special cases are:
//  Beta(±Inf, y > 0) = 0
//  Beta(±Inf, y ≤ 0) = NaN for integer y
//  Beta(-Inf, y < 0) = NaN for non-integer y
//  Beta(+Inf, y < 0) = ±Inf for non-integer y
//  Beta(x ≤ 0, y ≤ 0) = NaN for integer x and y
//  Beta(x ≤ 0, y) = NaN for integer x and non-integer y or integer y > -x
//  Beta(x, y) = 0 for non-integer x and y and integer x+y ≤ 0
//
func Beta(x, y float64) float64 {
	// special cases
	if math.IsInf(x, 0) || math.IsInf(y, 0) {
		if math.IsInf(y, 0) {
			x, y = y, x
		}
		switch {
		case y > 0:
			return 0
		case x < 0:
			return nan
		case isNonPosInt(y):
			return nan
		case math.Mod(math.Trunc(y-1), 2) == 0: // if even when rounded away from 0
			return +inf
		default:
			return -inf
		}
	}

	z := x + y

	ix := isNonPosInt(x)
	iy := isNonPosInt(y)
	iz := isNonPosInt(z)

	if iy && !ix { // wlog, swap x and y
		x, y = y, x
		ix, iy = iy, ix
	}

	sign := 1
	// non-positive integer cases
	if ix && (iy || !iz) { // all terms diverge or only Gamma(x) diverges
		return nan
	}
	if iz {
		if !ix { // only Gamma(x+y) diverges
			return 0
		}

		// Both Gamma(x) and Gamma(x+y) diverge, but their ratio
		// has a finite limit:
		//
		//   Gamma(x)    Gamma(1-x-y)       y
		//  ---------- = ------------ * (-1)
		//  Gamma(x+y)    Gamma(1-x)
		//

		if math.Mod(y, 2) == 1 {
			sign = -sign
		}

		// If x+y=0, the above simplifies to (-1)**y / Gamma(1+y)
		// so the function is just (-1)**y / y.

		if z == 0 {
			return float64(sign) / y
		}

		x, z = 1-z, 1-x
	}

	// Beta(x, y) = Gamma(x) * Gamma(y) / Gamma(x+y)

	if y < x { // let y ≥ x wlog
		x, y = y, x
	}

	// avoid overflows and improve accuracy by reducing arguments
	if math.Abs(x+y) > 50 || math.Abs(x) > 50 || math.Abs(y) > 50 {

		// use the Gauss multiplication formula with x = 2*x', etc., giving
		//
		//               Beta(x', y') * Beta(x'+1/2, y'+1/2)    Gamma(z'+1)
		//  Beta(x, y) = ----------------------------------- * -------------
		//                           2 * Sqrt(π)               Gamma(z'+1/2)
		//

		x /= 2
		y /= 2
		z /= 2

		// if z is still very large, use an asymptotic series for the ratio of Gamma functions
		if z > 170 {
			s := z + poleval(1/z, 1./8, 1./128, -5./1024, -21./32768, 399./262144)
			s /= math.Sqrt(z)
			return 0.5 * Beta(x, y) * Beta(x+0.5, y+0.5) * s / math.SqrtPi * float64(sign)
		}

		return 0.5 * Beta(x, y) * Beta(x+0.5, y+0.5) * math.Gamma(z+1) / (math.Gamma(z+0.5) * math.SqrtPi) * float64(sign)
	}

	// direct formula
	gx := math.Gamma(x)
	gy := math.Gamma(y)
	gz := math.Gamma(z)
	return gx * (gy / gz) * float64(sign)
}
