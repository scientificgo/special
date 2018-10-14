// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package poly

import "math"

// SphericalHarmonicY returns the angular portion of the solutions to Laplace's equation
// in spherical coordinates, where theta is in [0, π], phi is in [0, 2π] and |l| ≤ m.
//
// See http://mathworld.wolfram.com/SphericalHarmonic.html for more information.
func SphericalHarmonicY(l, m int, theta, phi float64) (float64, float64) {
	st, ct := math.Sincos(theta)
	switch {
	case m > l || m < -l || l < 0 || math.IsNaN(theta) || math.IsNaN(phi):
		return math.NaN(), math.NaN()
	case m == 0:
		re := math.Sqrt(float64(2*l+1)) * LegendreP(l, ct) / (2 * math.SqrtPi)
		return re, 0
	case m == -l:
		s := math.Sqrt(math.Gamma(2*float64(l+1))) / math.Gamma(float64(l+1))
		s *= math.Pow(st/2, float64(l)) / (2 * math.SqrtPi)
		im, re := math.Sincos(phi * float64(l))
		return re * s, -im * s
	}

	reflect := false
	if m < 0 {
		m = -m
		reflect = true
	}

	im, re := math.Sincos(phi * float64(m))
	s := math.Sqrt(float64(2*l+1) * math.Gamma(float64(l-m+1)) / math.Gamma(float64(l+m+1)))
	s *= LegendreAP(l, m, ct) / (2 * math.SqrtPi)
	if reflect {
		im = -im
		s *= float64(1 - 2*(m&1))
	}
	return re * s, im * s
}
