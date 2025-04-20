package special

import (
	"math"
)

// GammaIncL returns the lower incomplete gamma function,
// defined by
//
//	                  x
//	GammaIncL(a, x) = ∫ dt Exp(-t) * t**(a-1)
//	                 t=0
//
// GammaIncL also satisfies the identity
//
//	GammaIncL(a, x) + GammaIncU(a, x) = Gamma(a)
//
// where GammaIncU is the upper incomplete gamma function.
//
// See http://mathworld.wolfram.com/IncompleteGammaFunction.html
// for more information.
func GammaIncL(a, x float64) float64 {
	lga, sga := math.Lgamma(a)
	gp := GammaRegP(a, x)
	if gp < 0 {
		gp = -gp
		sga = -sga
	}
	return float64(sga) * math.Exp(lga+math.Log(gp))
}

// GammaIncU returns the lower incomplete gamma function,
// defined by
//
//	                  ∞
//	GammaIncU(a, x) = ∫ dt Exp(-t) * t**(a-1)
//	                 t=x
//
// GammaIncU also satisfies the identity
//
//	GammaIncL(a, x) + GammaIncU(a, x) = Gamma(a)
//
// where GammaIncL is the lower incomplete gamma function.
//
// See http://mathworld.wolfram.com/IncompleteGammaFunction.html
// for more information.
func GammaIncU(a, x float64) float64 {
	switch {
	case isNonPosInt(a):
		return math.Pow(x, a) * En(int(1-a), x)
	default:
		lga, sga := math.Lgamma(a)
		gq := GammaRegQ(a, x)
		if gq < 0 {
			gq = -gq
			sga = -sga
		}
		return float64(sga) * math.Exp(lga+math.Log(gq))
	}
}
