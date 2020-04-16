// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// The original C code, the long comment, and the constants
// below are from http://netlib.sandia.gov/cephes/misc.tgz.
// The go code is a simplified version of the original C.

/*							dawsn.c
 *
 *	Dawson's Integral
 *
 *
 *
 * SYNOPSIS:
 *
 * double x, y, dawsn();
 *
 * y = dawsn( x );
 *
 *
 *
 * DESCRIPTION:
 *
 * Approximates the integral
 *
 *                             x
 *                             -
 *                      2     | |        2
 *  dawsn(x)  =  exp( -x  )   |    exp( t  ) dt
 *                          | |
 *                           -
 *                           0
 *
 * Three different rational approximations are employed, for
 * the intervals 0 to 3.25; 3.25 to 6.25; and 6.25 up.
 *
 *
 * ACCURACY:
 *
 *                      Relative error:
 * arithmetic   domain     # trials      peak         rms
 *    IEEE      0,10        10000       6.9e-16     1.0e-16
 *    DEC       0,10         6000       7.4e-17     1.4e-17
 *
 *
 *
 * Cephes Math Library Release 2.8:  June, 2000
 * Copyright 1984, 1987, 1989, 2000 by Stephen L. Moshier
 */

// Dawson's integral, interval 0 to 3.25
var _AN = []float64{
	9.99999999999999994612E-1,
	-9.17480371773452345351E-2,
	4.22618223005546594270E-2,
	-8.50149846724410912031E-4,
	3.52513368520288738649E-4,
	3.07828309874913200438E-6,
	9.53151741254484363489E-7,
	1.94434204175553054283E-8,
	8.49262267667473811108E-10,
	1.13681498971755972054E-11,
}
var _AD = []float64{
	1.00000000000000000539E0,
	5.74918629489320327824E-1,
	1.58874241960120565368E-1,
	2.79448531198828973716E-2,
	3.48805814657162590916E-3,
	3.25524741826057911661E-4,
	2.32490249820789513991E-5,
	1.27258478273186970203E-6,
	5.21265281010541664570E-8,
	1.48864681368493396752E-9,
	2.40372073066762605484E-11,
}

// interval 3.25 to 6.25
var _BN = []float64{
	3.59233385440928410398E-11,
	-2.40274520828250956942E-9,
	9.10010780076391431042E-8,
	-2.14640351719968974225E-6,
	3.59641304793896631888E-5,
	-4.23209114460388756528E-4,
	3.66207612329569181322E-3,
	-2.18711255142039025206E-2,
	9.41512335303534411857E-2,
	-2.44754418142697847934E-1,
	5.08955156417900903354E-1,
}
var _BD = []float64{
	7.18466403235734541950E-11,
	-4.91324691331920606875E-9,
	1.89100358111421846170E-7,
	-4.55875153252442634831E-6,
	7.81025592944552338085E-5,
	-9.47996768486665330168E-4,
	8.48041718586295374409E-3,
	-5.31806367003223277662E-2,
	2.36706788228248691528E-1,
	-6.31839869873368190192E-1,
	1.00000000000000000000E0,
}

// 6.25 to infinity
var _CN = []float64{
	-4.86827613020462700845E-4,
	1.64837047825189632310E-2,
	-1.72858975380388136411E-1,
	6.29235242724368800674E-1,
	-5.90592860534773254987E-1,
}
var _CD = []float64{
	-9.73655226040941223894E-4,
	3.44278924041233391079E-2,
	-3.93708582281939493482E-1,
	1.73270799045947845857E0,
	-2.69820057197544900361E0,
	1.00000000000000000000E0,
}

// DawsonF returns the Dawson's integral F(x).
//
// Special cases are:
//  DawsonF(±0) = ±0
//  DawsonF(±Inf) = ±0
//
func DawsonF(x float64) float64 {
	if math.IsNaN(x) {
		return x
	}

	sign := math.Copysign(1, x)
	x = math.Abs(x)
	t := 1 / (x * x)

	switch {
	case x < 3.25:
		t = 1 / t
		y := x * poleval(t, _AN...) / poleval(t, _AD...)
		return sign * y

	case x < 6.25:
		y := 1.0/x + t*poleval(t, _BN...)/(poleval(t, _BD...)*x)
		return sign * 0.5 * y

	case x > 1e9:
		return sign * 0.5 / x

	default: // 6.25 to infinity
		y := 1/x + t*poleval(t, _CN...)/(poleval(t, _CD...)*x)
		return sign * 0.5 * y
	}
}
