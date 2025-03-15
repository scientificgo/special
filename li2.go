package special

import (
	"math"
)

// Li2 returns the secondary logarithmic integral of x, defined for x ≥ 0 by
//
//	         x
//	Li2(x) = ∫ dt / Log(t) = Li(x) - Li(2)
//	        t=2
//
// such that Li2(2) = 0, where Li(x) is the primary logarithmic integral.
//
// See http://mathworld.wolfram.com/LogarithmicIntegral.html for more information.
func Li2(x float64) float64 {
	const li2 = 1.045163780117492784844588889194613136522615578151201575832
	switch {
	case math.IsInf(x, 1):
		return x
	case x == 1:
		return math.Inf(-1)
	default:
		return Li(x) - li2
	}
}
