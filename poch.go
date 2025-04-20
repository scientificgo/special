package special

import (
	"math"
)

// Poch returns the kth Pochhammer symbol of x, defined by
//
//	Poch(x, k) = Gamma(x+k) / Gamma(x)
//
// See http://mathworld.wolfram.com/PochhammerSymbol.html for more information.
func Poch(x, k float64) float64 {
	if math.IsInf(x, 0) {
		if x > 0 && k > 0 {
			return x
		}
		return 0
	}

	if math.IsInf(k, 0) {
		if isNonPosInt(x) {
			return 0
		}
		if k < 0 {
			return math.NaN()
		}
		return k * float64(GammaSign(x))
	}

	return GammaRatio([]float64{x + k}, []float64{x})
}
