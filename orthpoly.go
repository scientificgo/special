// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// ChebyshevT returns the nth Chebyshev polynomial of the first kind at x.
//
// See http://mathworld.wolfram.com/ChebyshevPolynomialoftheFirstKind.html for more information.
func ChebyshevT(n int, x float64) float64 {
	s := 1
	if n < 0 {
		n = -n
	}
	if x < 0 {
		x = -x
		s = 1 - 2*(n&1)
	}

	switch {
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, 1):
		return float64(s) * x
	case x == 0:
		if n%2 == 1 {
			return 0
		}
		return float64(s) * math.Cos(math.Pi*float64(n)/2)
	case x == 1:
		return float64(s)
	case n == 0:
		return float64(s)
	case n == 1:
		return float64(s) * x
	}

	const nlarge = 45

	var res float64
	if n <= nlarge {
		tmp := 1.0
		res = x
		x *= 2
		for k := 2; k <= n; k++ {
			res, tmp = x*res-tmp, res
		}
	} else {
		// For large n, use the trigonometric definitions.
		if math.Abs(x) < 1 {
			res = math.Cos(float64(n) * math.Acos(x))
		} else {
			res = math.Cosh(float64(n) * math.Acosh(x))
		}
	}
	return float64(s) * res
}

// ChebyshevU returns the nth Chebyshev polynomial of the second kind at x.
//
// See http://mathworld.wolfram.com/ChebyshevPolynomialoftheSecondKind.html for more information.
func ChebyshevU(n int, x float64) float64 {
	s := 1
	if n <= -2 {
		s = -1
		n = -n - 2
	}
	if x < 0 {
		x = -x
		s *= 1 - 2*(n&1)
	}

	switch {
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, 1):
		return float64(s) * x
	case x == 0:
		if n%2 == 1 {
			return 0
		}
		return float64(s) * math.Cos(math.Pi*float64(n)/2)
	case x == 1:
		return float64(s * (n + 1))
	case n == -1:
		return 0
	case n == 0:
		return float64(s)
	case n == 1:
		return float64(s) * 2 * x
	}

	const nlarge = 55

	var res float64
	if n <= nlarge {
		tmp := 1.0
		x *= 2
		res = x
		for k := 2; k <= n; k++ {
			res, tmp = x*res-tmp, res
		}
	} else {
		// For large n, use the trigonometric definitions.
		if x < 1 {
			t := math.Acos(x)
			res = math.Sin(float64(n+1)*t) / math.Sin(t)
		} else {
			t := math.Acosh(x)
			res = math.Sinh(float64(n+1)*t) / math.Sinh(t)
		}
	}
	return float64(s) * res
}

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

// ZernikeR returns the nth Zernike polynomial with parameter m at x.
//
// See http://mathworld.wolfram.com/ZernikePolynomial.html for more information.
func ZernikeR(n, m int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0 || m < 0:
		return math.NaN()
	case (n-m)&1 == 1 || n < m:
		return 0
	default:
		s := (n - m) / 2
		s = 1 - 2*(s&1)
		return float64(s) * math.Pow(x, float64(m)) * JacobiP((n-m)/2, float64(m), 0, 1-2*x*x)
	}
}

// GegenbauerC returns the nth Gegenbauer polynomial with paramater a at x.
//
// See http://mathworld.wolfram.com/GegenbauerPolynomial.html for more information.
func GegenbauerC(n int, a, x float64) float64 {
	switch {
	case math.IsNaN(a) || math.IsNaN(x) || n < 0:
		return math.NaN()
	case a <= 0 && a == math.Trunc(a):
		return 0
	case n == 0:
		return 1
	case n == 1:
		return 2 * a * x
	}

	tmp := 1.0
	res := 2 * a * x
	for k := 1; k < n; k++ {
		p := 2 * (float64(k) + a) * x
		q := float64(k-1) + 2*a
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}

// LegendreP returns the nth Legendre polynomial of the first kind at x.
//
// See http://mathworld.wolfram.com/LegendrePolynomial.html for more information.
func LegendreP(n int, x float64) float64 {
	if n < 0 {
		n = -n - 1
	}
	switch {
	case math.IsNaN(x):
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return x
	}

	tmp := 1.0
	res := x
	for k := 1; k < n; k++ {
		p := float64(2*k+1) * x
		q := float64(k)
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}

// LegendreAP returns the nth associated Legendre polynomial of the first kind
// with parameter m at x.
//
// See http://mathworld.wolfram.com/AssociatedLegendrePolynomial.html for more information.
func LegendreAP(n, m int, x float64) float64 {
	// Calculate P(|m|, |m|, x) and then use the recurrence
	// relation to get P(n, |m|, x). Finally, if m < 0, use the
	// reflection formula between m and -m to get P(n, m, x).

	if n < 0 {
		n = -n - 1
	}

	sign := 1 - 2*(m&1)
	switch {
	case sign < 0 && math.Abs(x) > 1:
		return math.NaN()
	case m > n || m < -n:
		return 0
	}

	// Calculate the prefactor for cases with m < 0 using the reflection formula.
	negmfac := 1.0
	if m < 0 {
		m = -m
		negmfac = float64(sign) * math.Gamma(float64(n-m+1)) / math.Gamma(float64(n+m+1))
	}

	// P(|m|, |m|, x)
	res := float64(sign) * math.Pow(1-x*x, float64(m)/2) * math.Gamma(float64(m)+0.5) / math.SqrtPi * float64(int(1<<uint(m)))
	if n > m {
		// Use recurrance formula to go from P(|m|, |m|, x) to P(n, |m|, x),
		// using a special case to get P(|m|+1, |m|, x).
		tmp := res
		res = x * float64(2*m+1) * res
		for k := m + 1; k < n; k++ {
			res, tmp = (x*float64(2*k+1)*res-float64(k+m)*tmp)/float64(k-m+1), res
		}
	}
	return negmfac * res
}

// SphericalHarmonicY returns the angular portion of the solutions to Laplace's equation
// in spherical coordinates, where theta is in [0, π], phi is in [0, 2π] and |l| ≤ m.
//
// See http://mathworld.wolfram.com/SphericalHarmonic.html for more information.
func SphericalHarmonicY(l, m int, theta, phi float64) (float64, float64) {
	st, ct := math.Sincos(theta)
	switch {
	case m > l || m < -l || l < 0 || math.IsNaN(theta) || math.IsNaN(phi):
		return math.NaN(), math.NaN()
	case m == 0:
		re := math.Sqrt(float64(2*l+1)) * LegendreP(l, ct) / (2 * math.SqrtPi)
		return re, 0
	case m == -l:
		s := math.Sqrt(math.Gamma(2*float64(l+1))) / math.Gamma(float64(l+1))
		s *= math.Pow(st/2, float64(l)) / (2 * math.SqrtPi)
		im, re := math.Sincos(phi * float64(l))
		return re * s, -im * s
	}

	reflect := false
	if m < 0 {
		m = -m
		reflect = true
	}

	im, re := math.Sincos(phi * float64(m))
	s := math.Sqrt(float64(2*l+1) * math.Gamma(float64(l-m+1)) / math.Gamma(float64(l+m+1)))
	s *= LegendreAP(l, m, ct) / (2 * math.SqrtPi)
	if reflect {
		im = -im
		s *= float64(1 - 2*(m&1))
	}
	return re * s, im * s
}

// LegendreQ returns the nth Legendre polynomial of the second kind at x.
//
// See http://mathworld.wolfram.com/LegendreFunctionoftheSecondKind.html for more information.
func LegendreQ(n int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0 || x < -1 || x > 1:
		return math.NaN()
	case n == 0:
		return math.Log((1+x)/(1-x)) / 2
	case n == 1:
		return x*math.Log((1+x)/(1-x))/2 - 1
	}

	tmp := math.Log((1+x)/(1-x)) / 2
	res := tmp*x - 1
	for k := 1; k < n; k++ {
		p := float64(2*k+1) * x
		q := float64(k)
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}

// TODO LegendreAssociatedQ

// LaguerreL returns the nth Laguerre polynomial at x.
//
// See http://mathworld.wolfram.com/LaguerrePolynomial.html for more information.
func LaguerreL(n int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0:
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return 1 - x
	}

	tmp := 1.0
	res := 1 - x
	for k := 1; k < n; k++ {
		p := float64(2*k+1) - x
		q := float64(k)
		res, tmp = (p*res-q*tmp)/float64(k+1), res
	}
	return res
}

// LaguerreAL returns the nth associated Laguerre polynomial with parameter a at x.
//
// See http://mathworld.wolfram.com/AssociatedLaguerrePolynomial.html for more information.
func LaguerreAL(n int, a, x float64) float64 {
	switch {
	case math.IsNaN(a) || math.IsNaN(x) || n < 0:
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return 1 + a - x
	}

	scale := 1.0
	if a < 0 && a == math.Trunc(a) {
		// scale = Gamma(-n+|a|) / Gamma(-n)
		if -float64(n)-a <= 0 {
			// scale = (-1)**|a| n! / (n-|a|)!
			scale *= float64(1 - 2*(int(a)&1))
			lnfac, snfac := math.Lgamma(float64(n + 1))
			lanfac, sanfac := math.Lgamma(float64(n) + a + 1)
			scale *= float64(snfac*sanfac) * math.Exp(lanfac-lnfac)
		} else {
			return math.NaN()
		}
		n, a = n+int(a), -a
		scale *= math.Pow(x, a)
	}

	tmp := 1.0
	res := 1 + a - x
	for k := 1; k < n; k++ {
		p := a + float64(2*k) + 1 - x
		q := a + float64(k)
		r := float64(k + 1)
		res, tmp = (p*res-q*tmp)/r, res
	}
	return scale * res
}

// HermiteH returns the nth unnormalised, or physics, Hermite polynomial, which is
// related to the normalised Hermite polynomial by
//
//  H(n, x) = √2**n He(n, x√2)
//
// where He is the normalised Hermite polynomial.
//
// See http://mathworld.wolfram.com/HermitePolynomial.html for more information.
func HermiteH(n int, x float64) float64 {
	return math.Exp(float64(n)*math.Ln2/2) * HermiteHe(n, math.Sqrt2*x)
}

// HermiteHe returns the nth normalised Hermite polynomial, which is
// related to the "physics" Hermite polynomial by
//
//  H(n, x) = √2**n He(n, x√2)
//
// where H is the unnormalised, or physics, Hermite polynomial.
//
// See http://mathworld.wolfram.com/HermitePolynomial.html for more information.
func HermiteHe(n int, x float64) float64 {
	switch {
	case math.IsNaN(x) || n < 0:
		return math.NaN()
	case n == 0:
		return 1
	case n == 1:
		return x
	}

	tmp := 1.0
	res := x
	for k := 1; k < n; k++ {
		res, tmp = x*res-float64(k)*tmp, res
	}
	return res
}
