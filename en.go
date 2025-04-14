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
	case x > 5 || n >= 100:
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

// en_rec returns the exponential integral En(x) using the recurrence relation
//
//	En(n+1, x) = (Exp(-x) - x*En(n, x))/n
func en_rec(n int, x float64) float64 {
	k := math.Exp(-x)
	y := Ei(-x)
	switch res := 0.0; {
	case n == 2:
		return k + x*y
	case n == 3:
		return (k - x*(k+x*y)) / 2
	case n == 4:
		return (k - x/2*(k-x*(k+x*y))) / 3
	case n == 5:
		return (k - x/3*(k-x/2*(k-x*(k+x*y)))) / 4
	case n == 6:
		return (k - x/4*(k-x/3*(k-x/2*(k-x*(k+x*y))))) / 5
	case n == 7:
		return (k - x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y)))))) / 6
	case n == 8:
		return (k - x/6*(k-x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y))))))) / 7
	case n == 9:
		return (k - x/7*(k-x/6*(k-x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y)))))))) / 8
	case n >= 10:
		res = k - x/8*(k-x/7*(k-x/6*(k-x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y))))))))
		fallthrough
	default:
		if n == 10 {
			return res / 9
		}
		for i := 9; i < n-1; i++ {
			res *= -x / float64(i)
			res += k
		}
		return res / float64(n-1)
	}
}
