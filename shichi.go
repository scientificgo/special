// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// The original C code, the long comment, and the constants
// below are from http://netlib.sandia.gov/cephes/misc.tgz.
// The go code is a modified version of the original C.

/*							shichi.c
 *
 *	Hyperbolic sine and cosine integrals
 *
 *
 *
 * SYNOPSIS:
 *
 * double x, Chi, Shi, shichi();
 *
 * shichi( x, &Chi, &Shi );
 *
 *
 * DESCRIPTION:
 *
 * Approximates the integrals
 *
 *                            x
 *                            -
 *                           | |   cosh t - 1
 *   Chi(x) = eul + ln x +   |    -----------  dt,
 *                         | |          t
 *                          -
 *                          0
 *
 *               x
 *               -
 *              | |  sinh t
 *   Shi(x) =   |    ------  dt
 *            | |       t
 *             -
 *             0
 *
 * where eul = 0.57721566490153286061 is Euler's constant.
 * The integrals are evaluated by power series for x < 8
 * and by Chebyshev expansions for x between 8 and 88.
 * For large x, both functions approach exp(x)/2x.
 * Arguments greater than 88 in magnitude return MAXNUM.
 *
 *
 * ACCURACY:
 *
 * Test interval 0 to 88.
 *                      Relative error:
 * arithmetic   function  # trials      peak         rms
 *    DEC          Shi       3000       9.1e-17
 *    IEEE         Shi      30000       6.9e-16     1.6e-16
 *        Absolute error, except relative when |Chi| > 1:
 *    DEC          Chi       2500       9.3e-17
 *    IEEE         Chi      30000       8.4e-16     1.4e-16
 *
 *
 * Cephes Math Library Release 2.8:  June, 2000
 * Copyright 1984, 1987, 2000 by Stephen L. Moshier
 */

// Shichi returns Shi(x), Chi(x), the hyperbolic sine and cosine integrals
// of x respectively, or Shi(x), Chi(x)-iπ for negative x.
//
// Special cases are:
//  Shichi(0) = 0, -Inf
//  Shichi(±Inf) = ±Inf, +Inf
//
func Shichi(x float64) (shi, chi float64) {
	// special cases
	switch {
	case x == 0:
		shi = 0
		chi = -inf
		return
	case math.IsInf(x, 0):
		shi = +inf
		chi = +inf
		if x < 0 {
			shi = -shi
		}
		return
	case math.IsNaN(x):
		shi = nan
		chi = nan
		return
	}

	s := math.Copysign(1, x)
	if x < 0 {
		x = -x
	}

	switch {
	case x < 8: // direct power series expansion
		shi = 0
		chi = 0
		x2 := x * x

		y := 1.
		i := 2
		for math.Abs(y/shi) > macheps {
			y *= x2 / float64(i)
			chi += y / float64(i)
			i++

			y /= float64(i)
			shi += y / float64(i)
			i++
		}
		shi = shi*x + x

	case x < 18:
		y := (576./x - 52.) / 20.
		k := math.Exp(x) / x
		shi = k * chebeval(y, _shi1...)
		chi = k * chebeval(y, _chi1...)

	case x < 88:
		y := (6336./x - 212.) / 140.
		k := math.Exp(x) / x
		shi = k * chebeval(y, _shi2...)
		chi = k * chebeval(y, _chi2...)

	default: // asymptotic expansion for x ≥ 88

		//                                 1!     2!     3!
		// 2 * x * e**(-x) * Shi(x) ~ 1 + ---- + ---- + ---- + ...
		//                                 x     x**2   x**3

		y := 1 / x
		z := math.Exp(x / 2) // use e**x = (e**(x/2))**2 to avoid overflows
		shi = factorialseries(y) / 2 * y * z * z
		chi = shi // equal to machine precision
		goto done
	}

	chi += math.Log(x) + Euler

done:
	if s == -1 {
		shi = -shi
	}
	return
}

// x exp(-x) shi(x), inverted interval 8 to 18
var _shi1 = []float64{
	1.1184775104725704,
	0.029606444085563326,
	0.0020255847474384687,
	8.901367419507276e-05,
	-3.1245820216895986e-05,
	-5.3991911840380505e-06,
	7.820182151840513e-07,
	1.4481887738426735e-07,
	-3.5669961111498254e-08,
	-1.032571217928195e-09,
	1.3545546976724696e-09,
	-1.7789978443643032e-10,
	-1.6159618114543546e-11,
	9.490446262242236e-12,
	-1.4005976461311712e-12,
	-3.4719701049774914e-14,
	5.939762262643143e-14,
	-1.313135343440926e-14,
	1.0989694907490535e-15,
	2.0432610598087988e-16,
	-9.554855322796556e-17,
	1.8388923017339947e-17,
}

// x exp(-x) shi(x), inverted interval 18 to 88
var _shi2 = []float64{
	1.0366572258879834,
	0.012847806525964761,
	0.000349810375601054,
	1.610382601173763e-05,
	1.1622994706867733e-06,
	1.2539177122848704e-07,
	1.690502288794213e-08,
	1.7628162914426453e-09,
	-1.7928943718335563e-10,
	-1.585806616664827e-10,
	-3.492781410247309e-11,
	1.6689495475283908e-12,
	2.7260035212915307e-12,
	3.347149541759945e-13,
	-1.6081820451980247e-13,
	-4.211281703076408e-14,
	1.0176556596972905e-14,
	3.9339787543705e-15,
	-8.306080263669358e-16,
	-3.3845981187810305e-16,
	8.820901356253682e-17,
	2.624460955963552e-17,
	-1.0531157415485094e-17,
}

// x exp(-x) chin(x), inverted interval 8 to 18
var _chi1 = []float64{
	1.1144615087669922,
	0.026434749603137454,
	0.0004971647898231161,
	-0.00031308547749299745,
	-6.133870010764944e-05,
	4.335058892573164e-06,
	1.946355313732725e-06,
	-2.2177501880184887e-07,
	-4.7551393092476546e-08,
	1.3145815098947459e-08,
	-1.879920756405693e-10,
	-4.4009247621328235e-10,
	8.763022886090549e-11,
	-1.4035943813649127e-12,
	-2.9349607260759987e-12,
	7.163396491560285e-13,
	-6.008651785534474e-14,
	-1.2100997011373292e-14,
	5.355463116474652e-15,
	-9.488121105916906e-16,
	5.226243949240722e-17,
	2.175864132903392e-17,
	-8.12435385225864e-18,
}

// x exp(-x) chin(x), inverted interval 18 to 88
var _chi2 = []float64{
	1.0366569391793428,
	0.012847538753006526,
	0.000349592575153778,
	1.595031648023132e-05,
	1.069427655664015e-06,
	7.793874743909149e-08,
	-3.0009517802868167e-09,
	-4.725430648762717e-09,
	-1.6102137516380345e-09,
	-2.56600180000356e-10,
	2.7143600637761244e-11,
	2.3378184398545344e-11,
	2.7485114193531538e-12,
	-1.271354181323383e-12,
	-3.7190244809311924e-13,
	6.634067317189116e-14,
	3.4963969541080694e-14,
	-4.4282320733253195e-15,
	-3.1073491733529946e-15,
	4.523139416989047e-16,
	2.6853395108594575e-16,
	-5.981113296582723e-17,
	-2.0807416818014817e-17,
	8.069134082551556e-18,
}
