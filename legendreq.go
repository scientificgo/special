package special

import "math"

// LegendreQ returns the nth Legendre polynomial of the second kind at x.
//
// See https://mathworld.wolfram.com/LegendreFunctionoftheSecondKind.html for more information.
func LegendreQ(n int, x float64) float64 {
	if math.IsNaN(x) {
		return math.NaN()
	}

	if n < 0 || x < -1 || x > 1 {
		return math.NaN()
	}

	if math.Abs(x) == 1 {
		return math.Inf(1)
	}

	res := (math.Log1p(x) - math.Log1p(-x)) / 2
	if n == 0 {
		return res
	}

	if n == 1 {
		return x*res - 1
	}

	res, tmp := x*res-1, res
	for k := 2; k <= n; k++ {
		p := float64(2*k-1) / float64(k)
		q := float64(k-1) / float64(k)
		res, tmp = (p*x*res - q*tmp), res
	}

	return res

}
