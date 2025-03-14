package special

import "math"

// LaguerreL returns the nth Laguerre polynomial at x.
//
// See http://mathworld.wolfram.com/LaguerrePolynomial.html for more information.
func LaguerreL(n int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0:
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return 1 - x
	}

	tmp := 1.0
	res := 1 - x
	for k := 1; k < n; k++ {
		p := float64(2*k+1) - x
		q := float64(k)
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}
