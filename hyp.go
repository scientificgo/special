// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// HypPFQ returns the type-(p,q) generalised hypergeometric function, defined by
//
//	               ∞            p-1         q-1
//	pFq(a; b; x) = ∑ x**k / k! * ∏ (a[i])_k / ∏ (b[j])_k
//	              k=0           i=0         j=0
//
// where a and b are length p and q respectively and (x)_k is the k-th Pochhammer
// symbol of x. The function pFq is commonly denoted
//
//	   | a[0], a[1], ..., a[p-1]    |
//	pFq|                        ; x |
//	   | b[0], b[1], ..., b[q-1]    |
//
// for p, q ≥ 2.
//
// See http://mathworld.wolfram.com/HypergeometricFunction.html for more information.
func HypPFQ(a, b []float64, x float64) float64 {
	if x == 0 {
		return 1
	}

	// Remove values common to a and b.
	a, b, na, nb := removeCommonElements(a, b)

	// If a or b have any infinite parameters, note the index of those parameters.
	// Any infinite parameters remaining in a and b be oppositely signed due to the removal
	// of common values in the previous step.

	ainf := []int{}
	binf := []int{}
	for i, ai := range a {
		if math.IsInf(ai, 0) {
			ainf = append(ainf, i)
		}
	}
	for i, bi := range b {
		if math.IsInf(bi, 0) {
			binf = append(binf, i)
		}
	}
	nainf := len(ainf)
	nbinf := len(binf)

	// Deal with the infinite parameters. There are 3 cases:
	// 1) a and b contain a ±∞ pair. Then reduce pFq to p-1Fq-1 using:
	//    lim {t -> ∞} pFq(a0,...,±t;b0,...,∓t;x) = p-1Fq-1(a0,...;b0,...;-x)
	// 2) b contains ±∞ and a does not. Reduce to pFq-1 using
	//    lim {|t| -> ∞} pFq(a0,...;b0,...,t;x) = lim {|t| -> ∞} pFq-1(a0,...;b0,...;x/t)
	//    - If x is infinite then x/t -> ±1. Then the sum has a non-trivial value.
	//    - If x is finite then x/t = 0 and the result is trivially 1.
	// 3) a contains ±∞ and b does not. The function (obviously) diverges and the result is NaN.

	// Case 1. Reduce pFq to p-1Fq-1, remove infinite parameters and flip the sign of x.
	for nainf > 0 && nbinf > 0 {
		ainfi := ainf[nainf-1]
		nainf--
		a = append(a[:ainfi], a[ainfi+1:]...)
		na--

		binfi := binf[nbinf-1]
		nbinf--
		b = append(b[:binfi], b[binfi+1:]...)
		nb--

		x = -x
	}

	// Case 2. Reduce pFq to pFq-1 and examine whether x is infinite.
	for nbinf > 0 {
		if nbinf > 1 || !math.IsInf(x, 0) {
			return 1
		}
		binfi := binf[0]
		nbinf--
		x = math.Copysign(1, x*b[binfi])
		b = append(b[:binfi], b[binfi+1:]...)
		nb--
	}

	// Case 3. The sum is divergent.
	if nainf > 0 {
		return math.NaN()
	}

	// Sum reduces to ∑ x**k / k! = Exp(x) when na = nb = 0.
	if na == 0 && nb == 0 {
		return math.Exp(x)
	}

	// Sum reduces to (1 - x)**(-a[0]) when na = 0 & nb = 1.
	if na == 1 && nb == 0 {
		return math.Pow(1-x, -a[0])
	}

	// Get the greatest non-positive element from from a and b, or
	// the first element if all elements are positive.
	// If a or b are empty slices, use NaN.

	amin := math.NaN()
	if na > 0 {
		amin = a[0]
	}
	bmin := math.NaN()
	if nb > 0 {
		bmin = b[0]
	}
	for _, ai := range a {
		if ai <= 0 && (amin > 0 || ai > amin) {
			amin = ai
		}
	}
	for _, bi := range b {
		if bi <= 0 && (bmin > 0 || bi > bmin) {
			bmin = bi
		}
	}

	// Define parameters for infinite sum. istrunc is used to flag whether the sum is truncated.
	const tol = 1e-16
	numt := math.MaxInt32
	istrunc := false

	// Need to check for negative integers to address any that would
	// cause the sum to diverge. In particular, the sum will diverge when
	// amin = -i, bmin = -j if i > j, or amin ≠ -i and bmin = -j,
	// for integers i and j. In these cases the sum is undefined, i.e. NaN.
	//
	// If i ≤ j or bmin > 0, on the other hand, then the sum is truncated at term i as
	// all subsequent terms contain a factor of 0.

	if bmin <= 0 {
		if bmin == math.Trunc(bmin) && (amin != math.Trunc(amin) || bmin > amin || amin > 0) {
			return math.NaN()
		}
	}

	if amin <= 0 && math.Trunc(amin) == amin {
		numt = -int(amin)
		istrunc = true
	}

	// Sum diverges when na > nb + 1 unless it has been truncated.
	if na > nb+1 && !istrunc {
		return math.NaN()
	}

	// Sum diverges when na = nb + 1  and |x| > 1 unless it has been truncated.
	if na == nb+1 && math.Abs(x) > 1 && !istrunc {
		return math.NaN()
	}

	switch {
	case na == 1 && nb == 1:
		return hyp1f1(a[0], b[0], x, numt, tol, istrunc)
	case na == 2 && nb == 1:
		return hyp2f1(a[0], a[1], b[0], x, numt, tol)
	default:
		t := 1.0
		res := t
		for k := 1; k <= numt && math.Abs(t/res) > tol; k++ {
			kk := float64(k)
			t *= x / kk
			kk--
			for i := 0; i < na; i++ {
				t *= kk + a[i]
			}
			for i := 0; i < nb; i++ {
				t /= kk + b[i]
			}
			res += t
		}
		return res
	}
}

// hyp1f1 returns the 1F1(a; b; x) series, and assumes that b is not a non-positive integer.
func hyp1f1(a, b, x float64, numt int, tol float64, istrunc bool) float64 {
	t := 1.0
	n := int(math.Abs(x/700) + 1)
	count := n
	scalef := 1.0
	if x < 0 && !istrunc {
		a = b - a
		x = -x
		t *= math.Exp(-x)
		count = n
		if t < 1e-300 {
			t = math.Exp(-x / float64(n))
			scalef = t
			count = 1
		}
	}

	res := t
	for k := 1; k <= numt && math.Abs(t/res) > tol; k++ {
		kk := float64(k)
		t *= x / kk
		kk--
		t *= (kk + a) / (kk + b)
		if res > 1e200 && count < n {
			t *= scalef
			res *= scalef
			count++
		}
		res += t
	}

	for count < n {
		res *= scalef
		count++
	}
	return res
}

// hyp2f1 returns the 2F1(a, b; c; x) series, and assumes that c is not a non-positive integer.
func hyp2f1(a, b, c, x float64, numt int, tol float64) float64 {
	// For x < 0, transform to x -> x/(x-1) > 0. See 15.3.4, p559, Ambramowitz & Stegun.
	scale := 1.0
	if x < 0 {
		if b > a {
			b = c - b
		} else {
			a, b = b, c-a
		}
		scale *= math.Pow(1-x, -a)
		x /= (x - 1)
	}

	// Gauss's formula for x = 1.
	if x == 1 && c-a-b > 0 {
		lgc, sgc := math.Lgamma(c)
		lgcab, sgcab := math.Lgamma(c - a - b)
		lgca, sgca := math.Lgamma(c - a)
		lgcb, sgcb := math.Lgamma(c - b)
		return scale * float64(sgc*sgcab*sgca*sgcb) * math.Exp(lgc+lgcab-lgca-lgcb)
	}

	t := 1.0
	res := t
	for k := 1; k <= numt && math.Abs(t/res) > tol; k++ {
		kk := float64(k)
		t *= x / kk
		kk--
		t *= (kk + a) * (kk + b) / (kk + c)
		res += t
	}
	return scale * res
}
