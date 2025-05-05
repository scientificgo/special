package special

import "math"

// LegendreP returns the nth Legendre polynomial of the first kind at x.
//
// See http://mathworld.wolfram.com/LegendrePolynomial.html for more information.
func LegendreP(n int, x float64) float64 {
	if math.IsNaN(x) {
		return math.NaN()
	}

	if n < 0 {
		n = -(n + 1)
	}

	res, tmp := 1., 0.
	for k := 0; k < n; k++ {
		p := float64(2*k + 1)
		q := float64(k)
		res, tmp = (p*x*res-q*tmp)/float64(k+1), res
	}

	return res
}
