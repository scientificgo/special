package special

import "math"

// Beta returns the complete beta function, defined by
//
//	Beta(x, y) = Gamma(x) Gamma(y) / Gamma(x+y)
//
// where Gamma is the gamma function.
//
// See http://mathworld.wolfram.com/BetaFunction.html for more information.
func Beta(x, y float64) float64 {
	switch {
	case math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, -1) || math.IsInf(y, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		if isNonPosInt(y) {
			return float64(GammaSign(y)) * x
		}
		return 0
	case math.IsInf(y, 1):
		if isNonPosInt(x) {
			return float64(GammaSign(x)) * y
		}
		return 0
	}
	return GammaRatio([]float64{x, y}, []float64{x + y})
}
