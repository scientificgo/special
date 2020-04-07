// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Utility functions for evaluating the elementary trigonometric
// functions with arguments of the form v = x * π.
//
// They have true zeros and ones (and infinities) for x within ε
// of a non-zero integer or when x is exactly zero.

// SinPi returns Sin(π*x).
//
// Special cases are:
//  SinPi(±0) = ±0
//  SinPi(x) = 0 for integer x
//  SinPi(x) = ±1 for half-integer x
//  SinPi(±Inf) = NaN
//
func SinPi(x float64) float64 {
	if i, isint := isInt(x); isint {
		if i == 0 { // correctly signed zero
			return i
		}
		return 0
	}
	return math.Sin(math.Pi * math.Remainder(x, 2))
}

// CosPi returns Cos(π*x).
//
// Special cases are:
//  CosPi(x) = ±1 for integer x
//  CosPi(x) = 0 for half-integer x
//  CosPi(±Inf) = NaN
//
func CosPi(x float64) float64 {
	if _, isint := isInt(x + 0.5); isint {
		return 0
	}
	return math.Cos(math.Pi * math.Remainder(x, 2))
}

// SincosPi returns Sin(π*x), Cos(π*x).
//
// Special cases are:
//  SincosPi(±0) = ±0, 1
//  SincosPi(x) = 0, ±1 for integer x
//  SincosPi(x) = ±1, 0 for half-integer x
//  SincosPi(±Inf) = NaN, NaN
//
func SincosPi(x float64) (sin, cos float64) {
	if i, isint := isInt(x); isint {
		cos = -1
		if math.Mod(i, 2) == 0 { // even
			cos = 1
		}
		if i == 0 { // correctly signed zero
			sin = i
		}
		return
	}
	if i, isint := isInt(x + 0.5); isint {
		sin = 1
		if math.Mod(i, 2) == 0 { // even
			sin = -sin
		}
		return
	}
	sin, cos = math.Sincos(math.Pi * math.Remainder(x, 2))
	return
}

// TanPi returns Tan(π*x).
//
// Special cases are:
//  TanPi(±0) = ±0
//  TanPi(x) = 0 for integer x
//  TanPi(x) = +Inf for half-integer x
//  TanPi(±Inf) = NaN
//
func TanPi(x float64) float64 {
	if i, isint := isInt(x); isint {
		if i == 0 { // correctly signed zero
			return i
		}
		return 0
	}
	if _, isint := isInt(x + 0.5); isint {
		return +inf
	}
	return math.Tan(math.Pi * math.Remainder(x, 1))
}

// CotPi returns Cot(π*x).
//
// Special cases are:
//  CotPi(±0) = ±Inf
//  CotPi(x) = +Inf for integer x
//  CotPi(x) = 0 for half-integer x
//  CotPi(±Inf) = NaN
//
func CotPi(x float64) float64 {
	if i, isint := isInt(x); isint {
		if i == 0 { // correctly signed infinity
			return math.Copysign(inf, i)
		}
		return +inf
	}
	if _, isint := isInt(x + 0.5); isint {
		return 0
	}

	x = math.Remainder(x, 1)
	s := math.Copysign(1, x)
	x = math.Abs(x)

	if x > 0.40 { // Laurent series about x=1/2
		x -= 0.5
		t := 0.
		xi := 1.
		for _, c := range _cotPi {
			if math.Abs(xi*c/t) < macheps {
				break
			}
			t += xi * c
			xi *= x * x
		}
		return s * t * x
	}

	sin, cos := SincosPi(x)
	return s * cos / sin
}

// Laurent series coefficients for cot(πx), given by
// 2**(2k) * (2**(2k) - 1) * B_2k * π**(2k-1) / (2k)!
var _cotPi = []float64{
	// 0,
	-3.141592653589793,
	-10.33542556009994,
	-40.80262463803753,
	-162.99995197525544,
	-651.9097561459137,
	-2607.5995051461714,
	-10430.380532928235,
	-41721.51437137118,
	-166886.05403833318,
	-667544.2146215658,
	-2.670176857805526e+06,
	-1.0680707430919562e+07,
	-4.2722829723543786e+07,
	-1.7089131889411536e+08,
	-6.835652755764349e+08,
	-2.734261102305728e+09,
	-1.0937044409222906e+10,
	-4.3748177636891624e+10,
	-1.749927105475665e+11,
	-6.99970842190266e+11,
	-2.799883368761064e+12,
	-1.1199533475044256e+13,
	-4.479813390017702e+13,
	-1.791925356007081e+14,
	-7.167701424028324e+14,
	-2.8670805696113295e+15,
	-1.1468322278445318e+16,
	-4.587328911378127e+16,
	-1.834931564551251e+17,
	-7.339726258205004e+17,
	-2.9358905032820014e+18,
	-1.1743562013128006e+19,
	-4.697424805251202e+19,
	-1.878969922100481e+20,
	-7.515879688401924e+20,
	-3.0063518753607694e+21,
	-1.2025407501443078e+22,
	-4.810163000577231e+22,
	-1.9240652002308924e+23,
	-7.69626080092357e+23,
}
