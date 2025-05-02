package special

import (
	"math"
)

// En returns the En function, defined by
//
//	           ∞
//	En(n, x) = ∫ dt Exp(-x*t) / t**n
//	          t=1
//
// See http://mathworld.wolfram.com/En-Function.html for more information.
func En(n int, x float64) float64 {
	// Special cases.
	switch {
	case math.IsNaN(x) || n < 0 || (n == 0 && x == 0) || (x < 0 && n > 1):
		return math.NaN()
	case math.IsInf(x, 1):
		return 0
	}

	switch {
	case n == 0:
		return math.Exp(-x) / x
	case x == 0:
		return 1 / float64(n-1)
	case n == 1:
		return -Ei(-x)
	case n >= 100 || x > 5:
		return en_cf(n, x)
	default:
		return en_rec(n, x)
	}
}

// en_cf returns the exponential integral En(x) using a continued fraction.
func en_cf(n int, x float64) float64 {
	depth := 15
	res := 1.0
	for depth > 0 {
		b1 := float64(n + depth - 1)
		b2 := float64(depth)
		res = x + b1/(1+b2/res)
		depth--
	}
	return math.Exp(-x) / res
}

// en_rec returns the exponential integral En(x) using the recurrence relation for n >= 2
//
//	En(n+1, x) = (Exp(-x) - x*En(n, x))/n
func en_rec(n int, x float64) float64 {
	u := math.Exp(-x)
	v := En(1, x)

	res := u - v*x

	for i := 1; i < n-1; i++ {
		res = math.FMA(res, -x/float64(i), u)
	}

	return res / float64(n-1)
}
