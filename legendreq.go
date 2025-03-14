package special

import "math"

// LegendreQ returns the nth Legendre polynomial of the second kind at x.
//
// See http://mathworld.wolfram.com/LegendreFunctionoftheSecondKind.html for more information.
func LegendreQ(n int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0 || x < -1 || x > 1:
		return math.NaN()
	case n == 0:
		return math.Log((1+x)/(1-x)) / 2
	case n == 1:
		return x*math.Log((1+x)/(1-x))/2 - 1
	}

	tmp := math.Log((1+x)/(1-x)) / 2
	res := tmp*x - 1
	for k := 1; k < n; k++ {
		p := float64(2*k+1) * x
		q := float64(k)
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}
