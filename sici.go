// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// The constants below are from the publication:
// Rowe, B. T. P. et al. "GALSIM: The modular galaxy image simulation toolkit"
// Astronomy and Computing 10, 121–150 (2015). arXiv:1407.7676v3 [astro-ph.IM],
// which is available at https://arxiv.org/abs/1407.7676v3.

// Si returns the sine integral of x.
//
// Special cases are:
//  Si(0) = 0
//  Si(±Inf) = ±π/2
//
func Si(x float64) float64 {
	s := math.Copysign(1, x)
	x = math.Abs(x)
	switch {
	case math.IsInf(x, 0):
		return s * math.Pi / 2
	case x <= 4:
		return s * sismall(x)
	default:
		f := sicif(x)
		g := sicig(x)
		sin, cos := math.Sincos(x)
		return s * si(sin, cos, f, g)
	}
}

// Ci returns the cosine integral of x.
//
// Special cases are:
//  Ci(0) = -Inf
//  Ci(+Inf) = 0
//  Ci(x) = NaN for x < 0
//
func Ci(x float64) float64 {
	switch {
	case x < 0:
		return nan
	case math.IsInf(x, 1):
		return 0
	case x < 4:
		return Euler + math.Log(x) - cinsmall(x)
	default:
		f := sicif(x)
		g := sicig(x)
		sin, cos := math.Sincos(x)
		return ci(sin, cos, f, g)
	}
}

// Cin returns the complementary cosine integral of x.
//
// Special cases are:
//  Cin(0) = 0
//  Cin(±Inf) = +Inf
//
func Cin(x float64) float64 {
	x = math.Abs(x)
	switch {
	case math.IsInf(x, 0) || x == 0:
		return x
	case x < 4:
		return cinsmall(x)
	default:
		f := sicif(x)
		g := sicig(x)
		sin, cos := math.Sincos(x)
		return Euler + math.Log(x) - ci(sin, cos, f, g)
	}
}

// rational approximations for Si, Cin and Ci for |x| <= 4
func sismall(x float64) float64 {
	x2 := x * x
	return x * poleval(x2, _sismallA...) / poleval(x2, _sismallB...)
}
func cinsmall(x float64) float64 {
	x2 := x * x
	return -x2 * poleval(x2, _cinsmallA...) / poleval(x2, _cinsmallB...)
}

// sicif approximates Ci(x)*sin(x) + (π/2 - Si(x))*cos(x) for x > 4
func sicif(x float64) float64 {
	y := 1 / x
	y2 := y * y
	return y * poleval(y2, _sicifA...) / poleval(y2, _sicifB...)
}

// sicig approximates -Ci(x)*cos(x) + (π/2 - Si(x))*sin(x) for x > 4
func sicig(x float64) float64 {
	y2 := 1 / (x * x)
	return y2 * poleval(y2, _sicigA...) / poleval(y2, _sicigB...)
}

// si and ci are extract Si and Ci from sicif and sicig given sin and cos
func si(sin, cos, f, g float64) float64 { return math.Pi/2 - cos*f - sin*g }
func ci(sin, cos, f, g float64) float64 { return sin*f - cos*g }

// coefficients for sismall
var _sismallA = []float64{
	1.,
	-4.54393409816329991e-2,
	1.15457225751016682e-3,
	-1.41018536821330254e-5,
	9.43280809438713025e-8,
	-3.53201978997168357e-10,
	7.08240282274875911e-13,
	-6.05338212010422477e-16,
}
var _sismallB = []float64{
	1.,
	1.01162145739225565e-2,
	4.99175116169755106e-5,
	1.55654986308745614e-7,
	3.28067571055789734e-10,
	4.5049097575386581e-13,
	3.21107051193712168e-16,
}

// coefficients for cinsmall
var _cinsmallA = []float64{
	-0.25,
	7.51851524438898291e-3,
	-1.27528342240267686e-4,
	1.05297363846239184e-6,
	-4.68889508144848019e-9,
	1.06480802891189243e-11,
	-9.93728488857585407e-15,
}
var _cinsmallB = []float64{
	1.,
	1.1592605689110735e-2,
	6.72126800814254432e-5,
	2.55533277086129636e-7,
	6.97071295760958946e-10,
	1.38536352772778619e-12,
	1.89106054713059759e-15,
	1.39759616731376855e-18,
}

// coefficients for sicif
var _sicifA = []float64{
	1.,
	7.44437068161936700618e+2,
	1.96396372895146869801e+5,
	2.37750310125431834034e+7,
	1.43073403821274636888e+9,
	4.33736238870432522765e+10,
	6.40533830574022022911e+11,
	4.20968180571076940208e+12,
	1.00795182980368574617e+13,
	4.94816688199951963482e+12,
	-4.94701168645415959931e+11,
}
var _sicifB = []float64{
	1.,
	7.46437068161927678031e+2,
	1.97865247031583951450e+5,
	2.41535670165126845144e+7,
	1.47478952192985464958e+9,
	4.58595115847765779830e+10,
	7.08501308149515401563e+11,
	5.06084464593475076774e+12,
	1.43468549171581016479e+13,
	1.11535493509914254097e+13,
}

// coefficients for sicig
var _sicigA = []float64{
	1.,
	8.1359520115168615e2,
	2.35239181626478200e5,
	3.12557570795778731e7,
	2.06297595146763354e9,
	6.83052205423625007e10,
	1.09049528450362786e12,
	7.57664583257834349e12,
	1.81004487464664575e13,
	6.43291613143049485e12,
	-1.36517137670871689e12,
}
var _sicigB = []float64{
	1.,
	8.19595201151451564e2,
	2.40036752835578777e5,
	3.26026661647090822e7,
	2.23355543278099360e9,
	7.87465017341829930e10,
	1.39866710696414565e12,
	1.17164723371736605e13,
	4.01839087307656620e13,
	3.99653257887490811e13,
}
