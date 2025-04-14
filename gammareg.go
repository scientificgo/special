package special

import "math"

// GammaRegP returns the regularised lower incomplete gamma function,
// defined by
//
//	                                 x
//	GammaRegP(a, x) = [1 / Gamma(a)] ∫ dt Exp(-t) * t**(a-1)
//	                                t=0
//
// The regularised lower incomplete gamma function has a series representation
//
//	                               ∞
//	GammaRegP(a, x) = Exp(-x) a**x ∑ x**k / Gamma(a+k+1)
//	                              k=0
//
// and also satisfies the identity
//
//	GammaRegP(a, x) + GammaRegQ(a, x) = 1
//
// where GammaRegQ is the regularised upper incomplete gamma function.
//
// See http://mathworld.wolfram.com/RegularizedGammaFunction.html
// for more information.
func GammaRegP(a, x float64) float64 {
	// Special cases.
	switch {
	case x < 0 || math.IsNaN(x) || math.IsNaN(a) || math.IsInf(a, -1):
		return math.NaN()
	case x == 0:
		return 0
	case math.IsInf(a, 1):
		if math.IsInf(x, 1) {
			return 0.5
		}
		return 0
	case math.IsInf(x, 1) || (a <= 0 && a == math.Trunc(a)):
		return 1
	case a == 1:
		return 1 - math.Exp(-x)
	}

	// Use gammaQ as primary function and calculate using
	// the continued fraction representation.
	if x > a && !(x < 2 && a > -10) {
		return 1 - gammaQ_cf(a, x)
	}

	return gammaP_series(a, x)
}

// GammaRegQ returns the regularised upper incomplete gamma function,
// defined by
//
//	                                 ∞
//	GammaRegQ(a, x) = [1 / Gamma(a)] ∫ dt Exp(-t) * t**(a-1)
//	                                t=x
//
// GammaRegQ also satisfies the identity
//
//	GammaRegP(a, x) + GammaRegQ(a, x) = 1
//
// where GammaRegP is the regularised lower incomplete gamma function.
//
// See http://mathworld.wolfram.com/RegularizedGammaFunction.html
// for more information.
func GammaRegQ(a, x float64) float64 {
	// Special cases.
	switch {
	case x < 0 || math.IsNaN(x) || math.IsNaN(a) || math.IsInf(a, -1):
		return math.NaN()
	case x == 0:
		return 1
	case math.IsInf(a, 1):
		if math.IsInf(x, 1) {
			return 0.5
		}
		return 1
	case math.IsInf(x, 1) || (a <= 0 && a == math.Trunc(a)):
		return 0
	case a == 1:
		return math.Exp(-x)
	}

	// Calculate with the explicit power series when the continued
	// fraction for gammaQ can't be used.
	if x < a || (x < 2 && a > -10) {
		return 1 - gammaP_series(a, x)
	}

	return gammaQ_cf(a, x)
}

// gammaP_series returns GammaRegP using the hypergeometric series definition
func gammaP_series(a, x float64) float64 {
	const (
		maxiter = 2000
		rtol    = 1e-16
	)

	// Note that each term is proportional to the last by a factor of x/(a+k)
	// and hence it is unnecessary to calculate Γ(a+k+1) at each iteration.
	res := 1.0
	for k, tmp := 1, 1.0; k < maxiter && math.Abs(tmp/res) > rtol; k++ {
		tmp *= x / (a + float64(k))
		res += tmp
	}

	lga1, sga1 := math.Lgamma(a + 1)
	res *= float64(sga1) * math.Exp(a*math.Log(x)-x-lga1)
	if a > 0 {
		res = math.Min(res, 1)
	}
	return res
}

// gammaQ_cf returns GammaRegQ using a continued fraction.
func gammaQ_cf(a, x float64) float64 {
	lga, sga := math.Lgamma(a)
	s := math.Copysign(1, x)
	lx := math.Log(math.Abs(x))
	xma := x - a

	d := gammaQ_cfdepth(a, x)
	cf := xma + float64(d<<1+1)
	for i := d; i > 0; i-- {
		j := (i-1)<<1 + 1
		ai := float64(i) * (a - float64(i))
		bj := xma + float64(j)
		cf = bj + ai/cf
	}
	return s * float64(sga) * math.Exp(a*lx-x-lga) / cf
}

// gammaQ_cfdepth returns the depth required for convergence for the continued fraction for GammaRegQ.
func gammaQ_cfdepth(a, x float64) int {
	switch y := x / a; {
	case y > 1.5:
		return 10
	case y > 1.3:
		return 20
	case y > 1.1:
		return 40
	case y > 1.05:
		return 50
	default:
		return 100
	}
}
