package special

import (
	"math"
)

// Trigamma returns the logarithmic second derivative of Gamma(x), or, equivalently,
// the first derivative of the Digamma function.
//
//	Trigamma(x) = d/dx Digamma(x)
//
// See http://mathworld.wolfram.com/TrigammaFunction.html for more information.
func Trigamma(x float64) float64 {
	// Special cases.
	switch {
	case math.IsNaN(x) || math.IsInf(x, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		return 0
	case isNonPosInt(x):
		return math.NaN()
	}

	const xmin = 8

	// If |x| < xmin, use the recurrence relation to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		res += 1 / (x * x)
		x++
	}

	// For |x| > min, use an asymptotic (divergent) series expansion about x = ±∞.

	const (
		c0 = 1. / 6
		c1 = -1. / 30
		c2 = 1. / 42
		c3 = -1. / 30
		c4 = 5. / 66
		c5 = -691. / 2730
		c6 = 7. / 6
		c7 = -3617. / 510
	)
	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	xinv := 1 / x
	res += s * xinv * (1 + s*xinv/2 + y*poly(y, c0, c1, c2, c3, c4, c5, c6, c7))
	if s < 0 {
		cot := 1 / math.Tan(math.Pi*x)
		res += math.Pi * math.Pi * (1 + cot*cot)
	}
	return res
}
