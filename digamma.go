package special

import (
	"math"
)

// Digamma returns the first logarithmic derivative of the Gamma function, defined by
//
//	Digamma(x) = d/dx Lgamma(x)
//
// See http://mathworld.wolfram.com/DigammaFunction.html for more information.
func Digamma(x float64) float64 {
	// Special cases.
	switch {
	case math.IsNaN(x) || math.IsInf(x, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		return x
	case math.Trunc(x) == x && x <= 0:
		return math.NaN()
	}

	const xmin = 5

	// If |x| < xmin, use the recurrence relation Digamma(x+1) = Digamma(x) + 1/x
	// to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		res -= 1 / x
		x++
	}

	// For |x| > xmin, use the Taylor series expansion about x = ±∞.

	const (
		c0 = -1. / 12
		c1 = 1. / 120
		c2 = -1. / 252
		c3 = 1. / 240
		c4 = -1. / 132
		c5 = 691. / 32760
		c6 = -1. / 12
	)
	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	res += math.Log(x) - (s/2)/x + y*poly(y, c0, c1, c2, c3, c4, c5, c6)

	if s < 0 {
		res += math.Pi / math.Tan(math.Pi*x)
	}

	return res
}
