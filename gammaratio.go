package special

import "math"

// GammaRatio returns the ratio of products of Gamma functions, i.e.
//
//	                  n-1             m-1
//	GammaRatio(x, y) = ∏ Gamma(x[i]) / ∏ Gamma(y[j])
//	                  i=0             j=0
//
// where x = {x[0], ..., x[n-1]} and y = {y[0], ..., y[m-1]} are slices of length n
// and m respectively. The result is NaN if x and y contain a different number of
// infinite or NaN values.
//
// See http://mathworld.wolfram.com/GammaFunction.html for more information.
func GammaRatio(x, y []float64) float64 {
	lg, sg := LgammaRatio(x, y)
	return float64(sg) * math.Exp(lg)
}
