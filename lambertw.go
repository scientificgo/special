// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// The original C++ code and constants below are from
// https://github.com/DarkoVeberic/LambertW and the following publication:
// Darko Veberic, "Lambert W Function for Applications in Physics",
// Computer Physics Communications 183 (2012) 2622-2628, arXiv:1209.0735v2 [cs.MS],
// which is available at https://arxiv.org/abs/1209.0735v2.
// The go code is a simplified version of the original C++.

// W0 returns the principal value of the Lambert W function of x.
//
// Special cases are:
//  W0(0) = 0
//  W0(-1/e) = -1
//  W0(+Inf) = +Inf
//  W0(x < -1/e) = NaN
//
func W0(x float64) float64 {
	const xmin = -1 / math.E

	// special cases
	switch {
	case math.IsInf(x, 1):
		return x
	case x == 0:
		return 0
	case x <= xmin:
		if x < xmin {
			return nan
		}
		return -1
	}

	const (
		x0 = -0.32358170806015724
		x1 = 0.14546954290661823
		x2 = 8.706658967856612
	)

	var w float64
	switch {
	case x < x0: // branchpoint polynomial
		w = wbranch(x, 0)
	case x < x1: // rational polynomial w01A(x)/w01B(x)
		w = x * poleval(x, _w01A...) / poleval(x, _w01B...)
	case x < x2: // rational polynomial w02A(x)/w02B(x)
		w = x * poleval(x, _w02A...) / poleval(x, _w02B...)
	default: // asymptotic series approximation
		a := math.Log(x)
		b := math.Log(a)

		ba := b / a
		b2 := b * b
		b3 := b2 * b
		b4 := b2 * b2

		q0 := b - 2
		q1 := 2*b2 - 9*b + 6
		q2 := 3*b3 - 22*b2 + 36*b - 12
		q3 := 12*b4 - 125*b3 + 350*b2 - 300*b + 60
		w = a - b + ba*poleval(1/a, 1, q0/2, q1/6, q2/12, q3/60)
	}
	return wfritsch(w, x)
}

// W1 returns the secondary value of Lambert W function of x.
//
// Special cases are:
//  W1(0) = -Inf
//  W1(-1/e) = -1
//  W1(x > 0) = NaN
//  W1(x < -1/e) = NaN
//
func W1(x float64) float64 {
	const xmin = -1 / math.E

	// special cases
	if x >= 0 || x <= xmin || math.IsNaN(x) {
		if x == 0 {
			return -inf
		}
		if x == xmin {
			return -1
		}
		return nan
	}

	const (
		x0 = -0.30298541769
		x1 = -0.051012
	)

	var w float64
	switch { // initial estimate
	case x < x0: // branchpoint polynomial
		w = wbranch(x, -1)
	case x < x1: // rational polynomial w11A(x)/w11B(x)
		w = poleval(x, _w11A...) / poleval(x, _w11B...)
	default: // W(x) = ln(x / W(x)) = ln(x / ln(x / ln(x / ln(x / ...))))
		w = math.Log(-x)
		ww := 0.
		for math.Abs(ww/w-1) > 1e-3 {
			ww = w
			w = math.Log(x / w)
		}
	}
	return wfritsch(w, x)
}

// wbranch returns an estimate of W0(x) or W1(x)
// near the branch point at -1/e
func wbranch(x float64, k int) float64 {
	p2 := 1 + math.E*x
	p2 *= 2
	p := math.Sqrt(p2)
	if k == -1 {
		p = -p
	}
	return poleval(p, _wbranch...)
}

// wfritsch returns an improved approximation for W0(x)
// or W1(x), given an initial estimate, using Fritsch iteration
func wfritsch(w, x float64) float64 {
	z := math.Log(x/w) - w
	w1 := w + 1
	p := w / w1
	q := 2*w1*w1 + 4*w1*z/3
	return poleval(z, w, p, p/(q-2*z))
}

// wbranch coefficients
var _wbranch = []float64{
	-1.,
	1.,
	-0.3333333333333333,
	0.1527777777777778,
	-0.07962962962962963,
	0.04450231481481481,
	-0.02598471487360376,
	0.01563563253233392,
	-0.009616892024299432,
	0.006014543252956118,
	-0.0038112980348919993,
}

// w01A coefficients
var _w01A = []float64{
	1.,
	5.931375839364438,
	11.39220550532913,
	7.33888339911111,
	0.653449016991959,
}

// w01B coefficients
var _w01B = []float64{
	1,
	6.931373689597704,
	16.82349461388016,
	16.43072324143226,
	5.115235195211697,
}

// w02A coefficients
var _w02A = []float64{
	1.,
	2.445053070726557,
	1.343664225958226,
	0.148440055397592,
	0.0008047501729130,
}

// w02B coefficients
var _w02B = []float64{
	1.,
	3.444708986486002,
	3.292489857371952,
	0.916460018803122,
	0.0530686404483322,
}

// w11A coefficients
var _w11A = []float64{
	-7.81417672390744,
	253.88810188892484,
	657.9493176902304,
}

// w11B coefficients
var _w11B = []float64{
	1.,
	-60.43958713690808,
	99.9856708310761,
	682.6073999909428,
	962.1784396969866,
	1477.9341280760887,
}
