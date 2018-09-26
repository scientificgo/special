// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Digamma returns the first logarithmic derivative of the Gamma function, defined by
//
//  Digamma(x) = d/dx Lgamma(x)
//
// See http://mathworld.wolfram.com/DigammaFunction.html for more information.
func Digamma(x float64) float64 {
	// Special cases.
	switch {
	case math.IsNaN(x) || math.IsInf(x, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		return x
	case math.Trunc(x) == x && x <= 0:
		return math.NaN()
	}

	const xmin = 5

	// If |x| < xmin, use the recurrence relation Digamma(x+1) = Digamma(x) + 1/x
	// to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		res -= 1 / x
		x++
	}

	// For |x| > xmin, use the Taylor series expansion about x = ±∞.

	const (
		c0 = -1. / 12
		c1 = 1. / 120
		c2 = -1. / 252
		c3 = 1. / 240
		c4 = -1. / 132
		c5 = 691. / 32760
		c6 = -1. / 12
	)
	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	res += math.Log(x) - (s/2)/x + y*(c0+y*(c1+y*(c2+y*(c3+y*(c4+y*(c5+y*(c6)))))))

	if s < 0 {
		res += math.Pi / math.Tan(math.Pi*x)
	}
	return res
}

// Trigamma returns the logarithmic second derivative of Gamma(x), or, equivalently,
// the first derivative of the Digamma function.
//
//  Trigamma(x) = d/dx Digamma(x)
//
// See http://mathworld.wolfram.com/TrigammaFunction.html for more information.
func Trigamma(x float64) float64 {
	// Special cases.
	switch {
	case math.IsNaN(x) || math.IsInf(x, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		return 0
	case x <= 0 && math.Trunc(x) == x:
		return math.NaN()
	}

	const xmin = 8

	// If |x| < xmin, use the recurrence relation to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		res += 1 / (x * x)
		x++
	}

	// For |x| > min, use an asymptotic (divergent) series expansion about x = ±∞.

	const (
		c0 = 1. / 6
		c1 = -1. / 30
		c2 = 1. / 42
		c3 = -1. / 30
		c4 = 5. / 66
		c5 = -691. / 2730
		c6 = 7. / 6
		c7 = -3617. / 510
	)
	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	xinv := 1 / x
	res += s * xinv * (1 + s*xinv/2 + y*(c0+y*(c1+y*(c2+y*(c3+y*(c4+y*(c5+y*(c6+y*(c7)))))))))
	if s < 0 {
		cot := 1 / math.Tan(math.Pi*x)
		res += math.Pi * math.Pi * (1 + cot*cot)
	}
	return res
}

// Polygamma returns the nth derivative of the Digamma function.
//
//  Polygamma(n, x) = (d/dx)**n Digamma(x)
//
// See http://mathworld.wolfram.com/PolygammaFunction.html for more information.
func Polygamma(n int, x float64) float64 {
	switch {
	case n < 0 || math.IsNaN(x) || math.IsInf(x, -1) || (x <= 0 && math.Trunc(x) == x):
		return math.NaN()
	case n == 0:
		return Digamma(x)
	case n == 1:
		return Trigamma(x)
	case math.IsInf(x, 1):
		return 0
	case n == 2:
		return polygamma2(x)
	case n == 3:
		return polygamma3(x)
	case n == 4:
		return polygamma4(x)
	case n == 5:
		return polygamma5(x)
	default:
		// s = (-1)**n
		s := 1 - 2*(n&1)
		lnfac, _ := math.Lgamma(float64(n + 1))

		// Reflection formula, where pg(n, x) = polygamma(n, x):
		// pg(n, x) = (-1)**n pg(n, 1-x) - π (d/dx)**n cot(πx)
		c := 0.0
		reflect := false
		if x < 0 {
			c = scaledcotderiv(n, x)
			// If c!= 0 then multiply by n! to get the unscaled derivative.
			// For large n, use sign(c)*Exp(Log(n!) + Log|c|) to avoid overflow.
			if c != 0 {
				c *= math.Exp(lnfac)
			}
			if math.IsInf(c, 0) {
				return -c
			}
			x = 1 - x
			reflect = true
		}

		pg := 0.0
		xsmall := 2 * math.Sqrt(float64(n))
		if n >= 10 && x < xsmall {
			pg = polygammanseries(n, lnfac, x)
		} else {
			// Recurrence formula, where pg(n, x) = polygamma(n, x):
			// pg(n, x) = pg(n, x+1) - (-1)**n n! / x**(n+1)
			shift := 0.0
			xlarge := math.Min(float64(n), 10)
			for x < math.Max(2*xsmall, xlarge) {
				shift -= math.Exp(lnfac - float64(n+1)*math.Log(x))
				x++
			}
			pg = polygamman(n, lnfac, x) + float64(s)*shift
		}

		if reflect {
			return float64(s)*pg - c
		}
		return pg
	}
}

// scaledcotderiv returns π (d/dx)**n cot(πx) / n! using the derivative of Euler's partial fraction expansion
//                                          ∞
//  π (d/dx)**n cot(πx) / n! = 1/x**(n+1) + ∑ 1/(x+k)**(n+1) + 1/(x-k)**(n+1)
//                                         k=1
func scaledcotderiv(n int, x float64) float64 {
	// s = (-1)**n
	s := 1 - 2*(n&1)

	// The function cot(z) and its derivatives have a domain of z in [-π, π),
	// hence cot(πx) has domain x in [-1, 1).
	x = x - math.Trunc(x)

	// The nth derivative is an odd function when n is even and vice versa. We can therefore
	// always choose to make x > 0, if we add a factor of -1 to the result for even n.
	if x < 0 {
		x = -x
		if n&1 == 0 {
			s = -s
		}
	}

	// The nth derivative equals 0 at x=1/2 for even n since there is a leading factor of cot(πx),
	// which vanishes at x=1/2.
	if x == 1./2 && n&1 == 0 {
		return 0
	}

	const (
		maxiter = 200
		tol     = 1e-14
	)

	// Calculate the sum.
	res := math.Pow(x, -float64(n+1))
	tmp := x
	for k := 1; k <= maxiter && math.Abs(tmp/res) > tol; k++ {
		tmp = math.Pow(x+float64(k), -float64(n+1)) + math.Pow(x-float64(k), -float64(n+1))
		res += tmp
	}
	return float64(s) * res
}

// polygamman returns formula 6.4.11, p.260 from Ambramowitz & Stegun
func polygamman(n int, lnfac, x float64) float64 {
	// Coefficients for asymptotic expansion:
	// b_{2k} = BernoulliB_{2k}/(2k)!
	const (
		b2  = 1. / 12
		b4  = -1. / 720
		b6  = 1. / 30240
		b8  = -1. / 1209600
		b10 = 1. / 47900160
		b12 = -691. / 1307674368000
		b14 = 1. / 74724249600
		b16 = -3617. / 10670622842880000
		b18 = 43867. / 5109094217170944000
		b20 = -174611. / 802857662698291200000
		b22 = 77683. / 14101100039391805440000
		b24 = -236364091. / 1693824136731743669452800000
		b26 = 657931. / 186134520519971831808000000
		b28 = -3392780147. / 37893265687455865519472640000000
		/*
			b30 = 1723168255201. / 759790291646040068357842010112000000
			b32 = -7709321041217. / 134196726836183700385281186201600000000
			b34 = 151628697551. / 104199811425742637946218332815360000000
			b36 = -26315271553053477373. / 713925872841910517552409860896601407488000000000
			b38 = 154210205991661. / 165165037094716140555791754978970828800000000
			b40 = -261082718496449122051. / 11039333782344056345696120477635448049500160000000000
		*/
	)

	// s = (-1)**(n+1)
	s := -1 + 2*(n&1)

	y := 1 / (x * x)
	lx := math.Log(x)

	return float64(s) * math.Exp(lnfac-float64(n)*lx) * (1/float64(n) + 1/x*(0.5+
		float64(n+1)/x*(b2+
			float64((n+2)*(n+3))*y*(b4+
				float64((n+4)*(n+5))*y*(b6+
					float64((n+6)*(n+7))*y*(b8+
						float64((n+8)*(n+9))*y*(b10+
							float64((n+10)*(n+11))*y*(b12+
								float64((n+12)*(n+13))*y*(b14+
									float64((n+14)*(n+15))*y*(b16+
										float64((n+16)*(n+17))*y*(b18+
											float64((n+18)*(n+19))*y*(b20+
												float64((n+20)*(n+21))*y*(b22+
													float64((n+22)*(n+23))*y*(b24+
														float64((n+24)*(n+25))*y*(b26+
															float64((n+26)*(n+27))*y*b28)))))))))))))))
}

// polygammanseries returns polygamma(n, x) using the series definition
//                       ∞
//  polygamma(n, x) = n! ∑ (k+x)**(-n-1)
//                      k=0
func polygammanseries(n int, lnfac, x float64) float64 {
	const (
		maxiter = 200
		tol     = 1e-12
	)

	// s = (-1)**(n+1)
	s := -1 + 2*(n&1)

	res := 0.0
	tmp := 1.0
	for k := 0; k <= maxiter && math.Abs(tmp/res) > tol; k++ {
		kx := float64(k) + x
		tmp = math.Exp(lnfac - float64(n+1)*math.Log(kx))
		res += tmp
	}
	return float64(s) * res
}

func polygamma2(x float64) float64 {
	const xmin = 7

	// If |x| < xmin, use the recurrence relation to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		res -= 2 / (x * x * x)
		x++
	}

	// For |x| > min, use an asymptotic (divergent) series expansion about x = ±∞.

	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	const (
		c0 = -1
		c1 = -1
		c2 = -1. / 2
		c3 = 1. / 6
		c4 = -1. / 6
		c5 = 3. / 10
		c6 = -5. / 6
		c7 = 691. / 210
	)

	res += y * (c0 + s*c1/x + y*(c2+y*(c3+y*(c4+y*(c5+y*(c6+y*c7))))))
	if s < 0 {
		cot := 1 / math.Tan(math.Pi*x)
		csc := 1 / math.Sin(math.Pi*x)
		res += 2 * math.Pi * math.Pi * math.Pi * cot * csc * csc
	}
	return res
}

func polygamma3(x float64) float64 {
	const xmin = 16

	// If |x| < xmin, use the recurrence relation to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		x2 := x * x
		res += 6 / (x2 * x2)
		x++
	}

	// For |x| > min, use an asymptotic (divergent) series expansion about x = ±∞.

	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	const (
		c0 = 2
		c1 = 3
		c2 = 2
		c3 = -1
		c4 = 4. / 3
		c5 = -3. / 11
		c6 = 10
		c7 = -691. / 15
	)
	res += s * y / x * (c0 + s*c1/x + y*(c2+y*(c3+y*(c4+y*(c5+y*(c6+y*c7))))))
	if s < 0 {
		cot := 1 / math.Tan(math.Pi*x)
		cot2 := cot * cot
		res += 2 * math.Pi * math.Pi * math.Pi * math.Pi * (1 + cot2*(4+3*cot2))
	}
	return res
}

func polygamma4(x float64) float64 {
	const xmin = 13

	// If |x| < xmin, use the recurrence relation to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		x2 := x * x
		res -= 24 / (x2 * x2 * x)
		x++
	}

	// For |x| > min, use an asymptotic (divergent) series expansion about x = ±∞.

	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	const (
		c0 = -6
		c1 = -12
		c2 = -10
		c3 = 7
		c4 = -12
		c5 = 33
		c6 = -130
		c7 = 691
	)
	res += y * y * (c0 + s*c1/x + y*(c2+y*(c3+y*(c4+y*(c5+y*(c6+y*c7))))))

	if s < 0 {
		cot := 1 / math.Tan(math.Pi*x)
		csc := 1 / math.Sin(math.Pi*x)
		cot2 := cot * cot
		csc2 := csc * csc
		pi5 := math.Pi * math.Pi * math.Pi * math.Pi * math.Pi
		res += 8 * pi5 * csc2 * cot * (cot2 + 2*csc2)
	}
	return res
}

func polygamma5(x float64) float64 {
	const xmin = 10

	// If |x| < xmin, use the recurrence relation to increment x until x >= xmin.

	res := 0.0
	for math.Abs(x) < xmin {
		x2 := x * x
		res += 120 / (x2 * x2 * x2)
		x++
	}

	// For |x| > min, use an asymptotic (divergent) series expansion about x = ±∞.

	s := math.Copysign(1, x)
	x = math.Abs(x)
	y := 1 / (x * x)
	const (
		c0 = 24
		c1 = 60
		c2 = 60
		c3 = -56
		c4 = 120
		c5 = -396
		c6 = 1820
		c7 = -11056
	)
	res += s * y * y / x * (c0 + s*c1/x + y*(c2+y*(c3+y*(c4+y*(c5+y*(c6+y*c7))))))
	if s < 0 {
		cot := 1 / math.Tan(math.Pi*x)
		csc := 1 / math.Sin(math.Pi*x)
		cot2 := cot * cot
		csc2 := csc * csc

		pi6 := math.Pi * math.Pi * math.Pi * math.Pi * math.Pi * math.Pi
		res += 8 * pi6 * csc2 * (2*csc2*csc2 + 2*cot2*cot2 + 11*cot2*csc2)
	}
	return res
}

// Harmonic returns the harmonic numbers, defined for integer n by
//
//                n
//  Harmonic(n) = ∑ 1/k
//               k=1
//
// and extended to non-integer x by
//
//  Harmonic(x) = EulerGamma + Digamma(x+1)
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
