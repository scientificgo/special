/*
   SciGo is a scientific library for the Go language.
   Copyright (C) 2018, Jack Parkinson

   This program is free software: you can redistribute it and/or modify it
   under the terms of the GNU Lesser General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package special

import "math"

// Ei returns the exponential integral of x, defined by
//
//		x
//	Ei(x) = ∫ dt Exp(t) / t
//	       t=-∞
//
// See http://mathworld.wolfram.com/ExponentialIntegral.html for more information.
func Ei(x float64) float64 {
	const (
		xsmall = 3
		xlarge = 50
		xover  = 716
		xunder = -705
	)
	switch xabs := math.Abs(x); {
	case math.IsNaN(x):
		return math.NaN()
	case math.IsInf(x, -1) || x < xunder:
		return 0
	case math.IsInf(x, 1) || x > xover:
		return math.Inf(1)
	case x == 0:
		return math.Inf(-1)
	case xabs <= xsmall:
		return eismall(x, xabs)
	case xabs >= xlarge:
		return eilarge(x)
	case x < -xsmall:
		return eicf(x)
	default:
		return eiseries(x, xabs)
	}
}

// eismall returns the exponential integral Ei(x) for small |x|.
func eismall(x, xabs float64) float64 {
	const (
		c0  = 1.
		c1  = 1. / 4
		c2  = 1. / 18
		c3  = 1. / 96
		c4  = 1. / 600
		c5  = 1. / 4320
		c6  = 1. / 35280
		c7  = 1. / 322560
		c8  = 1. / 3265920
		c9  = 1. / 36288000
		c10 = 1. / 439084800
		c11 = 1. / 5748019200
		c12 = 1. / 80951270400
		c13 = 1. / 1220496076800
		c14 = 1. / 19615115520000
		c15 = 1. / 334764638208000
		c16 = 1. / 6046686277632000
		c17 = 1. / 115242726703104000
		c18 = 1. / 2311256907767808000
	)
	return math.Log(xabs) + EulerGamma +
		x*(c0+x*(c1+x*(c2+x*(c3+x*(c4+x*(c5+x*(c6+x*(c7+x*(c8+x*(c9+x*(c10+x*(c11+
			x*(c12+x*(c13+x*(c14+x*(c15+x*(c16+x*(c17+x*c18))))))))))))))))))
}

// eilarge returns the exponential integral Ei(x) for large |x|.
func eilarge(x float64) float64 {
	const (
		c0  = 1
		c1  = 1
		c2  = 2
		c3  = 6
		c4  = 24
		c5  = 120
		c6  = 720
		c7  = 5040
		c8  = 40320
		c9  = 362880
		c10 = 3628800
		c11 = 39916800
		c12 = 479001600
	)
	y := 1 / x
	sum := y * (c0 + y*(c1+y*(c2+y*(c3+y*(c4+y*(c5+y*(c6+y*(c7+y*(c8+y*(c9+y*(c10+y*(c11+y*(c12)))))))))))))

	s := math.Copysign(1, sum)
	sum = math.Abs(sum)
	return s * math.Exp(x+math.Log(sum))
}

// eicf returns the exponential integral Ei(x) using a continued fraction.
func eicf(x float64) float64 {
	// cf = a1 + b1/(a2 + b2/(a3 + b3/(...)))
	depth := 20
	an := float64(2*depth-1) - x
	bn := -float64(depth * depth)
	res := an
	for depth > 1 {
		depth--
		an -= 2
		bn += float64(depth<<1 + 1)
		res = an + bn/res
	}
	return -math.Exp(x) / res
}

// eiseries returns the exponential integral Ei(x) using the infinite series definition.
func eiseries(x, xabs float64) float64 {
	const (
		tol     = 1e-20
		maxiter = 1e3
	)

	// Do 5 steps per iteration; it is substantially faster based on empirical testing.
	res := math.Log(xabs) + EulerGamma + x
	for i, tmp := 2, x; i < maxiter && math.Abs(tmp/res) > tol; {
		tmp *= x * float64(i-1) / float64(i*i)
		res += tmp
		i++
		tmp *= x * float64(i-1) / float64(i*i)
		res += tmp
		i++
		tmp *= x * float64(i-1) / float64(i*i)
		res += tmp
		i++
		tmp *= x * float64(i-1) / float64(i*i)
		res += tmp
		i++
		tmp *= x * float64(i-1) / float64(i*i)
		res += tmp
		i++
	}
	return res
}

// Li returns the logarithmic integral of x, defined for x ≥ 0 by
//		x
//	Li(x) = ∫ dt / Log(t) = Ei(Log(x))
//	       t=0
//
// where Ei(x) is the exponential integral.
//
// See http://mathworld.wolfram.com/LogarithmicIntegral.html for more information.
func Li(x float64) float64 { return Ei(math.Log(x)) }

// Li2 returns the secondary logarithmic integral of x, defined for x ≥ 0 by
//
//		 x
//	Li2(x) = ∫ dt / Log(t) = Li(x) - Li(2)
//		t=2
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

// En returns the En function, defined by
//
//		   ∞
//	En(n, x) = ∫ dt Exp(-x*t) / t**n
//		  t=1
//
// See http://mathworld.wolfram.com/En-Function.html for more information.
func En(n int, x float64) float64 {
	// Special cases.
	switch {
	case math.IsNaN(x) || n < 0 || (n == 0 && x == 0) || (x < 0 && n > 1):
		return math.NaN()
	case math.IsInf(x, 1):
		return 0
	}

	switch {
	case n == 0:
		return math.Exp(-x) / x
	case x == 0:
		return 1 / float64(n-1)
	case n == 1:
		return -Ei(-x)
	case x > 5 || n >= 100:
		return encf(n, x)
	default:
		return enrec(n, x)
	}
}

// encf returns the exponential integral En(x) using a continued fraction.
func encf(n int, x float64) float64 {
	depth := 15
	res := 1.0
	for depth > 0 {
		b1 := float64(n + depth - 1)
		b2 := float64(depth)
		res = x + b1/(1+b2/res)
		depth--
	}
	return math.Exp(-x) / res
}

// enrec returns the exponential integral En(x) using the recurrence relation
// En(n+1, x) = (Exp(-x) - x*En(n, x))/n
func enrec(n int, x float64) float64 {
	k := math.Exp(-x)
	y := Ei(-x)
	switch res := 0.0; {
	case n == 2:
		return k + x*y
	case n == 3:
		return (k - x*(k+x*y)) / 2
	case n == 4:
		return (k - x/2*(k-x*(k+x*y))) / 3
	case n == 5:
		return (k - x/3*(k-x/2*(k-x*(k+x*y)))) / 4
	case n == 6:
		return (k - x/4*(k-x/3*(k-x/2*(k-x*(k+x*y))))) / 5
	case n == 7:
		return (k - x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y)))))) / 6
	case n == 8:
		return (k - x/6*(k-x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y))))))) / 7
	case n == 9:
		return (k - x/7*(k-x/6*(k-x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y)))))))) / 8
	case n >= 10:
		res = k - x/8*(k-x/7*(k-x/6*(k-x/5*(k-x/4*(k-x/3*(k-x/2*(k-x*(k+x*y))))))))
		fallthrough
	default:
		if n == 10 {
			return res / 9
		}
		for i := 9; i < n-1; i++ {
			res *= -x / float64(i)
			res += k
		}
		return res / float64(n-1)
	}
}
