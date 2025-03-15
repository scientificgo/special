package special

import (
	"math"
)

// Eta returns the Dirichlet eta function, defined by
//
//	         ∞
//	Eta(x) = ∑ (-1)**(n+1) / n**x = (1 - 2**(1-x)) Zeta(x)
//	        n=1
//
// where Zeta is the Riemann zeta function.
//
// See http://mathworld.wolfram.com/DirichletEtaFunction.html for more information.
func Eta(x float64) float64 {
	switch {
	case math.IsNaN(x) || math.IsInf(x, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		return 1
	case x == 0:
		return 1. / 2
	case x == 1:
		return math.Ln2
	case x == -1:
		return 1. / 4
	case x <= -2 && math.Trunc(x) == x && int(x)&1 == 0:
		return 0
	default:
		return (1 - math.Exp((1-x)*math.Ln2)) * Zeta(x)
	}
}
