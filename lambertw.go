package special

import "math"

// The following implementation is based on:
// Darko Veberic, "Lambert W Function for Applications in Physics",
// Computer Physics Communications 183 (2012) 2622-2628, arXiv:1209.0735v2

const lambertw_branchpoint = -1 / math.E

// LambertW returns the real branches of the Lambert W function, implicitly
// defined by
//
//	W(x) Exp(W(x)) = x
//
// or, equivalently, as the inverse of the function
//
//	f(x) = x Exp(x)
//
// where Exp is the exponential function. The principal branch (k=0) is defined
// on x ≥ -1/e and the secondary branch (k=-1) is defined over -1/e ≤ x < 0.
//
// See http://mathworld.wolfram.com/LambertW-Function.html for more information.
func LambertW(k int, x float64) float64 {
	// Special cases.
	switch {
	case k < -1 || k > 0 || x < lambertw_branchpoint || (k == -1 && x > 0) || math.IsNaN(x):
		return math.NaN()
	case x == 0:
		if k == 0 {
			return 0
		}
		return math.Inf(-1)
	case x == lambertw_branchpoint:
		return -1
	case math.IsInf(x, 1):
		return x
	}

	w := lambertw_estimate(k, x)
	return lambertw_fritsch(w, x)
}

// lambertw_estimate returns an lambertw_estimate estimate of W(x) on branch k
func lambertw_estimate(k int, x float64) float64 {
	switch k {
	case 0:
		const (
			xmax_branchpoint = -0.32358170806015724
			xmax_p0          = 0.14546954290661823
			xmax_p1          = 8.706658967856612
		)
		switch {
		case x < xmax_branchpoint:
			return lambertw_nearbranchpoint(k, x)
		case x < xmax_p0:
			return lambwertw_primary_rp0(x)
		case x < xmax_p1:
			return lambertw_primary_rp1(x)
		default:
			return lambertw_primary_asymptotic(x)
		}
	default: // k=-1
		const xmax_branchpoint = -0.30298541769
		switch {
		case x < xmax_branchpoint:
			return lambertw_nearbranchpoint(k, x)
		default:
			return lambertw_secondary_rp0(x)
		}
	}
}

// lambertw_nearbranchpoint returns an estimate of W(k, x)
// using an expansion around the branch point x=-1/e
func lambertw_nearbranchpoint(k int, x float64) float64 {
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
	return poly(p, b0, b1, b2, b3, b4, b5, b6, b7, b8, b9)
}

// lambwertw_primary_rp0 returns an approximation of W(x) (k=0)
// using a rational polynomial approximation
func lambwertw_primary_rp0(x float64) float64 {
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
	return x * poly(x, a0, a1, a2, a3, a4) / poly(x, b0, b1, b2, b3, b4)
}

// lambertw_primary_rp1 returns an approximation of W(x) (k=0) using a rational
// polynomial approximation
func lambertw_primary_rp1(x float64) float64 {
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
	return x * poly(x, a0, a1, a2, a3, a4) / poly(x, b0, b1, b2, b3, b4)
}

// lambertw_primary_asymptotic returns an asymptotic estimate of W(x) (k=0) for large x
func lambertw_primary_asymptotic(x float64) float64 {
	logx := math.Log(x)
	loglogx := math.Log(logx)
	r := loglogx / logx

	q1 := poly(loglogx, -2, 1)
	q2 := poly(loglogx, 6, -9, 2)
	q3 := poly(loglogx, -12, 36, -22, 3)
	q4 := poly(loglogx, 60, -300, 350, -125, 12)

	t := q1 + (1/(3*logx))*(q2+(1/(2*logx))*(q3+(1/(5*logx))*q4))
	return logx - loglogx + r*(1+t/(2*logx))
}

// lambertw_secondary_rp0 returns a rational polynomial approximation of
// W(x) (k=-1)
func lambertw_secondary_rp0(x float64) float64 {
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

	return poly(x, a0, a1, a2) / poly(x, b0, b1, b2, b3, b4, b5)
}

// lambertw_fritsch returns an improved approximation of W(x) given
// an initial guess w, on either branch
func lambertw_fritsch(w, x float64) float64 {
	z := math.Log(x/w) - w
	v := w + 1
	r := z / v * (1 + z/(2*(v*(v+2*z/3)-z)))
	return w * (1 + r)
}
