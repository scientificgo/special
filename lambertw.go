package special

import "math"

// The following implementation is based on:
// Darko Veberic, "Lambert W Function for Applications in Physics",
// Computer Physics Communications 183 (2012) 2622-2628, arXiv:1209.0735v2

// LambertW returns the real branches of the
// Lambert W function, implicitly defined by
//
//	W(x) Exp(W(x)) = x
//
// or, equivalently, as the inverse of the function
//
//	f(x) = x Exp(x)
//
// where Exp is the exponential function. The principal branch (k=0) is defined on
// x ≥ -1/e and the secondary branch (k=-1) is defined over -1/e ≤ x < 0.
//
// See http://mathworld.wolfram.com/LambertW-Function.html for more information.
func LambertW(k int, x float64) float64 {
	// Special cases.
	switch {
	case k < -1 || k > 0 || x < -1/math.E || (k == -1 && x > 0) || math.IsNaN(x):
		return math.NaN()
	case x == 0:
		if k == 0 {
			return 0
		}
		return math.Inf(-1)
	case x == -1/math.E:
		return -1
	case math.IsInf(x, 1):
		return x
	}

	// Estimate an initial value using approximations and then use
	// Fritsch iteration (once) to get an improved estimate with O(1e-15) error
	w := initial(k, x)
	return fritsch(w, x)
}

// initial returns an initial estimate of W(x) on branch k
func initial(k int, x float64) float64 {
	switch k {
	case 0:
		const (
			xbranch = -0.32358170806015724
			xratp0  = 0.14546954290661823
			xratp1  = 8.706658967856612
		)
		switch {
		case x < xbranch:
			return branchpoint(k, x)
		case x < xratp0:
			return rationalp0(x)
		case x < xratp1:
			return rationalp1(x)
		default:
			return asymptotic(k, x)
		}
	default: // k=-1
		const xbranch = -0.30298541769
		switch {
		case x < xbranch:
			return branchpoint(k, x)
		default:
			return rationalm(x)
		}
	}
}

// branchpoint returns an estimate of W(k, x)
// using an expansion around the branch point x=-1/e
func branchpoint(k int, x float64) float64 {
	s := 1 + 2*k
	p := float64(s) * math.Sqrt2 * math.Sqrt(1+math.E*x)
	const (
		b0 = -1
		b1 = 1
		b2 = -0.3333333333333333
		b3 = 0.1527777777777778
		b4 = -0.07962962962962963
		b5 = 0.04450231481481481
		b6 = -0.02598471487360376
		b7 = 0.01563563253233392
		b8 = -0.009616892024299432
		b9 = 0.006014543252956118
	)
	return b0 + p*(b1+p*(b2+p*(b3+p*(b4+p*(b5+p*(b6+p*(b7+p*(b8+p*b9))))))))
}

// rational returns an approximation of W(x) (k=0)
// using a rational polynomial approximation
func rationalp0(x float64) float64 {
	const (
		a0 = 1
		a1 = 5.931375839364438
		a2 = 11.39220550532913
		a3 = 7.33888339911111
		a4 = 0.653449016991959

		b0 = 1
		b1 = 6.931373689597704
		b2 = 16.82349461388016
		b3 = 16.43072324143226
		b4 = 5.115235195211697
	)
	num := a0 + x*(a1+x*(a2+x*(a3+x*a4)))
	den := b0 + x*(b1+x*(b2+x*(b3+x*b4)))
	return x * num / den
}

func rationalp1(x float64) float64 {
	const (
		a0 = 1
		a1 = 2.445053070726557
		a2 = 1.343664225958226
		a3 = 0.148440055397592
		a4 = 0.0008047501729130

		b0 = 1
		b1 = 3.444708986486002
		b2 = 3.292489857371952
		b3 = 0.916460018803122
		b4 = 0.0530686404483322
	)
	num := a0 + x*(a1+x*(a2+x*(a3+x*a4)))
	den := b0 + x*(b1+x*(b2+x*(b3+x*b4)))
	return x * num / den
}

// rationalm returns a rational polynomial approximation of W(x) (k=-1)
func rationalm(x float64) float64 {
	const (
		a0 = -7.81417672390744
		a1 = 253.88810188892484
		a2 = 657.9493176902304

		b0 = 1
		b1 = -60.43958713690808
		b2 = 99.9856708310761
		b3 = 682.6073999909428
		b4 = 962.1784396969866
		b5 = 1477.9341280760887
	)

	return (a0 + x*(a1+x*a2)) / (b0 + x*(b1+x*(b2+x*(b3+x*(b4+x*b5)))))
}

// asymptotic returns an asymptotic estimate of W(x, k)
func asymptotic(k int, x float64) float64 {
	s := 1 + 2*k
	a := math.Log(float64(s) * x)
	b := math.Log(float64(s) * a)

	ba := b / a
	b2 := b * b
	b3 := b2 * b
	b4 := b2 * b2

	q0 := b - 2
	q1 := 2*b2 - 9*b + 6
	q2 := 3*b3 - 22*b2 + 36*b - 12
	q3 := 12*b4 - 125*b3 + 350*b2 - 300*b + 60
	return a - b + ba*(1+1/(2*a)*(q0+1/(3*a)*(q1+1/(2*a)*(q2+1/(5*a)*q3))))
}

// fritsch returns an improved approximation of W(x) given
// an initial guess w, on either branch
func fritsch(w, x float64) float64 {
	z := math.Log(x/w) - w
	w1 := w + 1
	q := 2 * w1 * (w1 + 2*z/3)
	eps := z / w1 * (q - z) / (q - 2*z)
	return w * (1 + eps)
}
