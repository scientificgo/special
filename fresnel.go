// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// The original C code, the long comment, and the constants
// below are from http://netlib.sandia.gov/cephes/misc.tgz.
// The go code is a simplified version of the original C.

/*							fresnl.c
 *
 *	Fresnel integral
 *
 *
 *
 * SYNOPSIS:
 *
 * double x, S, C;
 * void fresnl();
 *
 * fresnl( x, _&S, _&C );
 *
 *
 * DESCRIPTION:
 *
 * Evaluates the Fresnel integrals
 *
 *           x
 *           -
 *          | |
 * C(x) =   |   cos(pi/2 t**2) dt,
 *        | |
 *         -
 *          0
 *
 *           x
 *           -
 *          | |
 * S(x) =   |   sin(pi/2 t**2) dt.
 *        | |
 *         -
 *          0
 *
 *
 * The integrals are evaluated by a power series for x < 1.
 * For x >= 1 auxiliary functions f(x) and g(x) are employed
 * such that
 *
 * C(x) = 0.5 + f(x) sin( pi/2 x**2 ) - g(x) cos( pi/2 x**2 )
 * S(x) = 0.5 - f(x) cos( pi/2 x**2 ) - g(x) sin( pi/2 x**2 )
 *
 *
 *
 * ACCURACY:
 *
 *  Relative error.
 *
 * Arithmetic  function   domain     # trials      peak         rms
 *   IEEE       S(x)      0, 10       10000       2.0e-15     3.2e-16
 *   IEEE       C(x)      0, 10       10000       1.8e-15     3.3e-16
 *   DEC        S(x)      0, 10        6000       2.2e-16     3.9e-17
 *   DEC        C(x)      0, 10        5000       2.3e-16     3.9e-17
 *
 * Cephes Math Library Release 2.8:  June, 2000
 * Copyright 1984, 1987, 1989, 2000 by Stephen L. Moshier
 */

// S(x) for small x
var _sn = []float64{
	3.18016297876567817986E11,
	-4.42979518059697779103E10,
	2.54890880573376359104E9,
	-6.29741486205862506537E7,
	7.08840045257738576863E5,
	-2.99181919401019853726E3,
}
var _sd = []float64{
	6.07366389490084639049E11,
	2.24411795645340920940E10,
	4.19320245898111231129E8,
	5.17343888770096400730E6,
	4.55847810806532581675E4,
	2.81376268889994315696E2,
	1.00000000000000000000E0,
}

// C(x) for small x
var _cn = []float64{
	9.99999999999999998822E-1,
	-2.05525900955013891793E-1,
	1.88843319396703850064E-2,
	-6.45191435683965050962E-4,
	9.50428062829859605134E-6,
	-4.98843114573573548651E-8,
}
var _cd = []float64{
	1.00000000000000000118E0,
	4.12142090722199792936E-2,
	8.68029542941784300606E-4,
	1.22262789024179030997E-5,
	1.25001862479598821474E-7,
	9.15439215774657478799E-10,
	3.99982968972495980367E-12,
}

// Auxiliary function f(x)
var _fn = []float64{
	3.76329711269987889006E-20,
	1.34283276233062758925E-16,
	1.72010743268161828879E-13,
	1.02304514164907233465E-10,
	3.05568983790257605827E-8,
	4.63613749287867322088E-6,
	3.45017939782574027900E-4,
	1.15220955073585758835E-2,
	1.43407919780758885261E-1,
	4.21543555043677546506E-1,
}
var _fd = []float64{
	1.25443237090011264384E-20,
	4.52001434074129701496E-17,
	5.88754533621578410010E-14,
	3.60140029589371370404E-11,
	1.12699224763999035261E-8,
	1.84627567348930545870E-6,
	1.55934409164153020873E-4,
	6.44051526508858611005E-3,
	1.16888925859191382142E-1,
	7.51586398353378947175E-1,
	1.00000000000000000000E0,
}

// Auxiliary function g(x)
var _gn = []float64{
	1.86958710162783235106E-22,
	8.36354435630677421531E-19,
	1.37555460633261799868E-15,
	1.08268041139020870318E-12,
	4.45344415861750144738E-10,
	9.82852443688422223854E-8,
	1.15138826111884280931E-5,
	6.84079380915393090172E-4,
	1.87648584092575249293E-2,
	1.97102833525523411709E-1,
	5.04442073643383265887E-1,
}
var _gd = []float64{
	1.86958710162783236342E-22,
	8.39158816283118707363E-19,
	1.38796531259578871258E-15,
	1.10273215066240270757E-12,
	4.60680728146520428211E-10,
	1.04314589657571990585E-7,
	1.27545075667729118702E-5,
	8.14679107184306179049E-4,
	2.53603741420338795122E-2,
	3.37748989120019970451E-1,
	1.47495759925128324529E0,
	1.00000000000000000000E0,
}

// Fresnel returns the Fresnel integrals S(x) and C(x).
//
// Special cases are:
//  Fresnel(±0) = ±0
//  Fresnel(±Inf) = 0.5, 0.5
//
func Fresnel(x float64) (s, c float64) {
	switch {
	case math.IsNaN(x):
		s = x
		c = x
		return
	case math.Abs(x) > 1e30:
		s = math.Copysign(0.5, x)
		c = s
		return
	}

	x2 := x * x
	if x2 < 2.5625 {
		t := x2 * x2
		s = x * x2 * poleval(t, _sn...) / poleval(t, _sd...)
		c = x * poleval(t, _cn...) / poleval(t, _cd...)
		return
	}

	// Asymptotic power series auxiliary functions
	// for large argument

	y := 1 / (math.Pi * x2)
	y2 := y * y
	f := 1 - y2*poleval(y2, _fn...)/poleval(y2, _fd...) // f(x)
	g := y * poleval(y2, _gn...) / poleval(y2, _gd...)  // g(x)

	sin, cos := SincosPi(0.5 * x2)
	pix := math.Pi * x
	a := math.Copysign(0.5, x)
	c = a + (f*sin-g*cos)/pix
	s = a - (f*cos+g*sin)/pix
	return
}
