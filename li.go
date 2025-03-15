package special

import (
	"math"
)

// Li returns the logarithmic integral of x, defined for x ≥ 0 by
//
//	        x
//	Li(x) = ∫ dt / Log(t) = Ei(Log(x))
//	       t=0
//
// where Ei(x) is the exponential integral.
//
// See http://mathworld.wolfram.com/LogarithmicIntegral.html for more information.
func Li(x float64) float64 { return Ei(math.Log(x)) }
