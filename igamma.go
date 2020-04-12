package special

import "math"

// Igammal returns the "lower" incomplete Gamma function
// of a and x.
//
// Special cases are:
//  Incgammal(a < 0, x < 0) = NaN for non-integer a
//  Incgammal(-Inf, x) = NaN
//  Incgammal(+Inf, x) = 0
//  Incgammal(a, +Inf) = Gamma(a) for finite a
//  Incgammal(a, 0) = 0
//  Incgammal(±0, x) = ±Inf
//  Incgammal(a < 0, x) = NaN for integer a
//
func Igammal(a, x float64) float64 {
	switch {
	case math.IsNaN(x) || math.IsNaN(a) || math.IsInf(a, -1) || (x < 0 && a < 0 && !isInt(a)):
		return nan
	case math.IsInf(x, 1):
		return math.Gamma(a)
	case math.IsInf(a, 1):
		return 0
	case x == 0:
		return 0
	case isNonPosInt(a):
		if a == 0 {
			return math.Copysign(inf, a)
		}
		return nan
	case x > a && x > 1:
		return math.Gamma(a) - Igammau(a, x)
	}

	ax := math.Pow(x, a) * math.Exp(-x)
	if x < 1 {
		return ax * hyp1F1(1, 1+a, x) / a
	}
	return ax * igammalcf(a, x)
}

// Igammau returns the "upper" incomplete Gamma function
// of a and x.
//
// Special cases are:
//  Incgammau(a < 0, x < 0) = NaN for non-integer a
//  Incgammau(-Inf, x) = NaN
//  Incgammau(+Inf, x) = +Inf
//  Incgammau(a, +Inf) = 0 for finite a
//  Incgammau(a, 0) = Gamma(a)
//  Incgammau(a < 0, x) = x**a * En(1-a, x) for integer a
//
func Igammau(a, x float64) float64 {
	switch {
	case math.IsNaN(x) || math.IsNaN(a) || math.IsInf(a, -1) || (x < 0 && a < 0 && !isInt(a)):
		return nan
	case math.IsInf(x, 1):
		if math.IsInf(a, 1) {
			return +inf
		}
		return 0
	case math.IsInf(a, 1):
		return +inf
	case x == 0:
		return math.Gamma(a)
	case isNonPosInt(a):
		n := int(1.5 - a)
		return math.Pow(x, a) * En(n, x)
	case x < a || x < 1:
		return math.Gamma(a) - Igammal(a, x)
	}

	ax := math.Pow(x, a) * math.Exp(-x)
	return ax * igammaucf(a, x)
}

// gammacfu evaluates igammau(a, x) * e**x * x**(-a)
// using a generalized continued fraction.
func igammaucf(a, x float64) float64 {
	ai := func(i int) float64 {
		j := float64(i - 1)
		return j * (a - j)
	}
	bi := func(i int) float64 {
		j := float64(2*i - 1)
		return j + x - a
	}
	return gcfsteed(ai, bi)
}

// igammacfl evaluates igammal(a,x) * x**(-a) * e**(x)
// using a generalized continued fraction.
func igammalcf(a, x float64) float64 {
	ai := func(i int) float64 {
		j := float64((i - 1) / 2)
		if i%2 == 0 {
			j = -j - a
		}
		return j * x
	}
	bi := func(i int) float64 {
		return a + float64(i-1)
	}
	return gcfsteed(ai, bi)
}

// gcfsteed evaluates the generalized continued fraction
// defined by the coefficient functions a and b using
// Steed's adapted forward recurrence algorithm with
// Kahan-Babusak-Neumaier summation to reduce roundoff
// error accumulation and numerical instability.
//
// Note that it is assumed, wlog, that b(0)=0 and a(1)=1;
// for alternative values, say b(0)=b0 and a(1)=a1, the
// result is given in terms of this function's result y
// as y' = b0 + a1*y.
//
// See http://dlmf.nist.gov/3.10.E16
// and http://dlmf.nist.gov/3.10.E17.
func gcfsteed(a, b func(int) float64) float64 {
	D := 1 / b(1)
	dC := D // * a(1)
	C := dC // + b(0)

	r := 0. // accumulator for roundoff errors
	for i := 2; math.Abs(dC/C) > macheps; i++ {
		ai := a(i)
		bi := b(i)
		D = 1 / (ai*D + bi)
		dC *= bi*D - 1

		// KBN summation
		tmp := C + dC
		if math.Abs(C) >= math.Abs(dC) {
			r += (C - tmp) + dC
		} else {
			r += (dC - tmp) + C
		}
		C = tmp
	}
	return C + r
}
