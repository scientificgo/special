package special

import (
	"math"
)

// Harmonic returns the harmonic numbers, defined for integer n by
//
//	              n
//	Harmonic(n) = ∑ 1/k
//	             k=1
//
// and extended to non-integer x by
//
//	Harmonic(x) = EulerGamma + Digamma(x+1)
//
// where Digamma is the logarithmic derivative of the Gamma function.
//
// See http://mathworld.wolfram.com/HarmonicNumber.html for more information.
func Harmonic(x float64) float64 {
	switch {
	case math.IsInf(x, 1) || x == 0 || x == 1:
		return x
	case x >= 1 && x <= 25 && x == math.Trunc(x):
		res := 1.0
		for ; x > 1; x-- {
			res += 1 / x
		}
		return res
	default:
		return EulerGamma + Digamma(x+1)
	}
}
