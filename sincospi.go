// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// sin(πx) and cos(πx) with true zeros

func sinPi(x float64) float64 {
	if isInt(x) {
		return 0
	}
	return math.Sin(math.Pi * math.Remainder(x, 2))
}

func cosPi(x float64) float64 {
	if isInt(x + 0.5) {
		return 0
	}
	return math.Cos(math.Pi * math.Remainder(x, 2))
}

func sincosPi(x float64) (sin, cos float64) {
	switch {
	case isInt(x):
		sin = 0
		cos = math.Cos(math.Pi * math.Remainder(x, 2))
	case isInt(x + 0.5):
		sin = math.Sin(math.Pi * math.Remainder(x, 2))
		cos = 0
	default:
		sin, cos = math.Sincos(math.Pi * math.Remainder(x, 2))
	}
	return
}

func tanPi(x float64) float64 {
	sin, cos := sincosPi(x)
	return sin / cos
}

func cotPi(x float64) float64 {
	sin, cos := sincosPi(x)
	return cos / sin
}
