package special

import "math"

// Si, Ci, f and g implementations from:
// Rowe, B. T. P. et al. GALSIM: The modular galaxy image simulation toolkit.
// Astronomy and Computing 10, 121–150 (2015). arXiv:1407.7676 [astro-ph.IM]

// Si returns the sine integral, defined by
//
//	         x                       ∞
//	 Si(x) = ∫ dt Sin(t) / t = π/2 - ∫ dt Sin(t) / t
//		       t=0                     t=x
//
// See http://mathworld.wolfram.com/SineIntegral.html for more information.
func Si(x float64) float64 {
	if math.IsNaN(x) {
		return math.NaN()
	}

	if math.IsInf(x, 0) {
		return math.Copysign(math.Pi/2, x)
	}

	switch xabs := math.Abs(x); {
	case xabs <= 4:
		const (
			c0 = 1
			c1 = -4.54393409816329991e-2
			c2 = 1.15457225751016682e-3
			c3 = -1.41018536821330254e-5
			c4 = 9.43280809438713025e-8
			c5 = -3.53201978997168357e-10
			c6 = 7.08240282274875911e-13
			c7 = -6.05338212010422477e-16

			d0 = 1
			d1 = 1.01162145739225565e-2
			d2 = 4.99175116169755106e-5
			d3 = 1.55654986308745614e-7
			d4 = 3.28067571055789734e-10
			d5 = 4.5049097575386581e-13
			d6 = 3.21107051193712168e-16
		)
		x2 := x * x

		return x * poly(x2, c0, c1, c2, c3, c4, c5, c6, c7) /
			poly(x2, d0, d1, d2, d3, d4, d5, d6)

	default:
		s := math.Copysign(1, x)
		sin, cos := math.Sincos(xabs)
		return s * (math.Pi/2 - cos*fsici(xabs) - sin*gsici(xabs))
	}
}

// Ci returns the cosine integral, defined by
//
//	          ∞                              x
//	Ci(x) = - ∫ dt Cos(t) / t = γ + Log(x) + ∫ dt (Cos(t) - 1) / t
//	         t=x                            t=0
//
// where γ is the Euler-Mascheroni constant.
//
// See http://mathworld.wolfram.com/CosineIntegral.html for more information.
func Ci(x float64) float64 {
	switch {
	case x < 0:
		return math.NaN()
	case math.IsInf(x, 1):
		return 0
	case x <= 4:
		return EulerGamma + math.Log(x) - cinsmall(x)
	default:
		return cilarge(x)
	}
}

// Cin returns the secondary cosine integral, defined by
//
//	         x
//	Cin(x) = ∫ dt (1 - Cos(t)) / t = γ + Log(x) - Ci(x)
//	        t=0
//
// where γ is the Euler-Mascheroni constant and Ci is the primary
// cosine integral.
//
// See http://mathworld.wolfram.com/CosineIntegral.html for more information.
func Cin(x float64) float64 {
	switch xabs := math.Abs(x); {
	case math.IsNaN(xabs) || math.IsInf(xabs, 0) || xabs == 0:
		return xabs
	case xabs <= 4:
		return cinsmall(xabs)
	default:
		return EulerGamma + math.Log(xabs) - cilarge(xabs)
	}
}

// cilarge returns an approximation to Ci for x >= 4
func cilarge(x float64) float64 {
	sin, cos := math.Sincos(x)
	return sin*fsici(x) - cos*gsici(x)
}

// cin returns a rational approximation for Cin for x <= 4
func cinsmall(x float64) float64 {
	const (
		c0 = -1. / 4
		c1 = 7.51851524438898291e-3
		c2 = -1.27528342240267686e-4
		c3 = 1.05297363846239184e-6
		c4 = -4.68889508144848019e-9
		c5 = 1.06480802891189243e-11
		c6 = -9.93728488857585407e-15

		d0 = 1
		d1 = 1.1592605689110735e-2
		d2 = 6.72126800814254432e-5
		d3 = 2.55533277086129636e-7
		d4 = 6.97071295760958946e-10
		d5 = 1.38536352772778619e-12
		d6 = 1.89106054713059759e-15
		d7 = 1.39759616731376855e-18
	)
	x2 := x * x
	return -x2 * poly(x2, c0, c1, c2, c3, c4, c5, c6) /
		poly(x2, d0, d1, d2, d3, d4, d5, d6, d7)
}

// fsici returns Ci(x)*Sin(x) + (π/2 - Si(x))*Cos(x)
func fsici(x float64) float64 {
	const (
		c0  = 1
		c1  = 7.44437068161936700618e2
		c2  = 1.96396372895146869801e5
		c3  = 2.37750310125431834034e7
		c4  = 1.43073403821274636888e9
		c5  = 4.33736238870432522765e10
		c6  = 6.40533830574022022911e11
		c7  = 4.20968180571076940208e12
		c8  = 1.00795182980368574617e13
		c9  = 4.94816688199951963482e12
		c10 = -4.94701168645415959931e11

		d0 = 1
		d1 = 7.46437068161927678031e2
		d2 = 1.97865247031583951450e5
		d3 = 2.41535670165126845144e7
		d4 = 1.47478952192985464958e9
		d5 = 4.58595115847765779830e10
		d6 = 7.08501308149515401563e11
		d7 = 5.06084464593475076774e12
		d8 = 1.43468549171581016479e13
		d9 = 1.11535493509914254097e13
	)
	y2 := 1 / (x * x)
	return 1 / x * poly(y2, c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) /
		poly(y2, d0, d1, d2, d3, d4, d5, d6, d7, d8, d9)
}

// gsici returns -Ci(x)*Cos(x) + (π/2 - Si(x))*Sin(x)
func gsici(x float64) float64 {

	const (
		c0  = 1
		c1  = 8.1359520115168615e2
		c2  = 2.35239181626478200e5
		c3  = 3.12557570795778731e7
		c4  = 2.06297595146763354e9
		c5  = 6.83052205423625007e10
		c6  = 1.09049528450362786e12
		c7  = 7.57664583257834349e12
		c8  = 1.81004487464664575e13
		c9  = 6.43291613143049485e12
		c10 = -1.36517137670871689e12

		d0 = 1
		d1 = 8.19595201151451564e2
		d2 = 2.40036752835578777e5
		d3 = 3.26026661647090822e7
		d4 = 2.23355543278099360e9
		d5 = 7.87465017341829930e10
		d6 = 1.39866710696414565e12
		d7 = 1.17164723371736605e13
		d8 = 4.01839087307656620e13
		d9 = 3.99653257887490811e13
	)
	y2 := 1 / (x * x)
	return y2 * poly(y2, c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) /
		poly(y2, d0, d1, d2, d3, d4, d5, d6, d7, d8, d9)
}
