package special

import "math"

// LegendreP returns the nth Legendre polynomial of the first kind at x.
//
// See http://mathworld.wolfram.com/LegendrePolynomial.html for more information.
func LegendreP(n int, x float64) float64 {
	if n < 0 {
		n = -n - 1
	}
	switch {
	case math.IsNaN(x):
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return x
	}

	tmp := 1.0
	res := x
	for k := 1; k < n; k++ {
		p := float64(2*k+1) * x
		q := float64(k)
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}
