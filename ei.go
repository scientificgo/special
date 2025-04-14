package special

import "math"

// Ei returns the exponential integral of x, defined by
//
//	        x
//	Ei(x) = ∫ dt Exp(t) / t
//	       t=-∞
//
// See http://mathworld.wolfram.com/ExponentialIntegral.html for more information.
func Ei(x float64) float64 {
	const (
		xsmall = 3
		xlarge = 50
		xover  = 716
		xunder = -705
	)
	switch xabs := math.Abs(x); {
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, -1) || x < xunder:
		return 0
	case math.IsInf(x, 1) || x > xover:
		return math.Inf(1)
	case x == 0:
		return math.Inf(-1)
	case xabs <= xsmall:
		return ei_small(x, xabs)
	case xabs >= xlarge:
		return ei_large(x)
	case x < -xsmall:
		return ei_cf(x)
	default:
		return ei_series(x, xabs)
	}
}

// ei_small returns the exponential integral Ei(x) for small |x|.
func ei_small(x, xabs float64) float64 {
	const (
		c0  = 1.
		c1  = 1. / 4
		c2  = 1. / 18
		c3  = 1. / 96
		c4  = 1. / 600
		c5  = 1. / 4320
		c6  = 1. / 35280
		c7  = 1. / 322560
		c8  = 1. / 3265920
		c9  = 1. / 36288000
		c10 = 1. / 439084800
		c11 = 1. / 5748019200
		c12 = 1. / 80951270400
		c13 = 1. / 1220496076800
		c14 = 1. / 19615115520000
		c15 = 1. / 334764638208000
		c16 = 1. / 6046686277632000
		c17 = 1. / 115242726703104000
		c18 = 1. / 2311256907767808000
	)
	return math.Log(xabs) + EulerGamma +
		x*poly(x, c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16, c17, c18)
}

// ei_large returns the exponential integral Ei(x) for large |x|.
func ei_large(x float64) float64 {
	const (
		c0  = 1
		c1  = 1
		c2  = 2
		c3  = 6
		c4  = 24
		c5  = 120
		c6  = 720
		c7  = 5040
		c8  = 40320
		c9  = 362880
		c10 = 3628800
		c11 = 39916800
		c12 = 479001600
	)
	y := 1 / x
	sum := y * poly(y, c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12)

	s := math.Copysign(1, sum)
	sum = math.Abs(sum)
	return s * math.Exp(x+math.Log(sum))
}

// ei_cf returns the exponential integral Ei(x) using a continued fraction.
func ei_cf(x float64) float64 {
	// cf = a1 + b1/(a2 + b2/(a3 + b3/(...)))
	depth := 20
	an := float64(2*depth-1) - x
	bn := -float64(depth * depth)
	res := an
	for depth > 1 {
		depth--
		an -= 2
		bn += float64(depth<<1 + 1)
		res = an + bn/res
	}
	return -math.Exp(x) / res
}

// ei_series returns the exponential integral Ei(x) using the infinite series definition.
func ei_series(x, xabs float64) float64 {
	const (
		tol     = 1e-16
		maxiter = 1e3
	)

	res := math.Log(xabs) + EulerGamma + x
	for i, tmp := 2, x; i < maxiter && math.Abs(tmp/res) > tol; i++ {
		tmp *= x * float64(i-1) / float64(i*i)
		res += tmp
	}
	return res
}
