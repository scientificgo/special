package special

import "math"

// GegenbauerC returns the nth Gegenbauer polynomial with paramater a at x.
//
// See http://mathworld.wolfram.com/GegenbauerPolynomial.html for more information.
func GegenbauerC(n int, a, x float64) float64 {
	switch {
	case math.IsNaN(a) || math.IsNaN(x) || n < 0:
		return math.NaN()
	case a <= 0 && a == math.Trunc(a):
		return 0
	case n == 0:
		return 1
	case n == 1:
		return 2 * a * x
	}

	tmp := 1.0
	res := 2 * a * x
	for k := 1; k < n; k++ {
		p := 2 * (float64(k) + a) * x
		q := float64(k-1) + 2*a
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}
