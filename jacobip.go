package special

import "math"

// JacobiP returns the nth Jacobi polynomial with parameters a, b at x.
//
// See http://mathworld.wolfram.com/JacobiPolynomial.html for more information.
func JacobiP(n int, a, b, x float64) float64 {
	s := 1
	if x < 0 {
		s = 1 - 2*(n&1)
		a, b = b, a
		x = -x
	}

	var res float64
	switch {
	case math.IsNaN(x) || math.IsNaN(a) || math.IsNaN(b) || math.IsInf(a, 0) || math.IsInf(b, 0) || (a < -float64(n) && a < 0 && a == math.Trunc(a)):
		res = math.NaN()
	case n < 0:
		res = 0
	case n == 0:
		res = 1
	case n == 1:
		res = ((a+b+2)*x + a - b) / 2
	case a+b > -2 || a+b < -2*float64(n-1) || a+b != math.Trunc(a+b):
		tmp := 1.0
		res = ((a+b+2)*x + a - b) / 2
		for k := 1; k < n; k++ {
			ka := float64(k) + a
			kb := float64(k) + b
			kkab := ka + kb
			kkab1 := kkab + 1
			kkab2 := kkab1 + 1

			p := kkab1*(a*a-b*b) + kkab*kkab1*kkab2*x
			q := ka * kb * kkab2
			r := float64(k+1) * (ka + b + 1) * kkab
			res, tmp = (p/2*res-q*tmp)/r, res
		}
	default:
		// Can't use simple recusive relation when -2 ≤ a+b ≤ -2n + 2 as encounter 0 in denominator, for integer a+b.
		if a > 0 {
			n1 := float64(n) + 1
			res = HypPFQ([]float64{-float64(n), a + b + n1}, []float64{a + 1}, (1-x)/2) * GammaRatio([]float64{a + n1}, []float64{n1, a + 1})
		} else if b > 0 {
			n1 := float64(n) + 1
			res = HypPFQ([]float64{-float64(n), a + b + n1}, []float64{b + 1}, (1+x)/2) * GammaRatio([]float64{-b}, []float64{n1, -b - float64(n)})
		} else {
			n1 := float64(n) + 1
			an := a + float64(n)
			abn := an + b
			ab2n := abn + float64(n)
			res = HypPFQ([]float64{-float64(n), -an}, []float64{-ab2n}, 2/(1-x)) * GammaRatio([]float64{ab2n + 1}, []float64{n1, abn + 1}) * math.Pow((x-1)/2, float64(n))
		}
	}
	return float64(s) * res
}
