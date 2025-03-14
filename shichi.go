package special

import "math"

// Chi returns the hyperbolic cosine integral evaluated at x.
func Chi(x float64) float64 {
	return Ei(x) - Shi(x)
}

// Shi returns the hyperbolic sine integral evaluated at x.
func Shi(x float64) float64 {
	if x == 0 {
		return 0
	}

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	if x < 35 {
		z := x * x
		a := 1.0
		s := 1.0
		k := 2.0
		for math.Abs(a/s) > 1e-16 {
			a *= z / k
			k++
			a /= k
			s += a / k
			k++
		}

		return float64(sign) * s * x
	}

	return float64(sign) * Ei(x) / 2
}
