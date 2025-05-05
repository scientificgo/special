package special

import "math"

// LegendreAP returns the nth associated Legendre polynomial of the first kind
// with parameter m at x.
//
// See http://mathworld.wolfram.com/AssociatedLegendrePolynomial.html for more information.
func LegendreAP(n, m int, x float64) float64 {
	sign := float64(powN1(m)) // -1 for odd m

	if sign < 0 && math.Abs(x) > 1 {
		return math.NaN()
	}

	// P(-n, m, x) = P(n-1, m, x)
	if n < 0 {
		n = -(n + 1)
	}

	// P(n, -m, x) = r(n, m) * P(n, m, x) using the reflection formula
	reflect := m < 0
	if reflect {
		m = -m
	}

	// P(n<m, m, x) = 0
	if n < m {
		return 0
	}

	if x < 0 {
		x = -x
		sign *= float64(powN1(n - m))
	}

	// Base case with n=m
	// P(m, m, x) = (-1)^m * (2*m-1)!! * (1-x^2)^(m/2)
	res := sign *
		math.Pow(2, float64(m)) * math.Gamma(float64(m)+0.5) / math.SqrtPi * // (2*m-1)!!, see https://functions.wolfram.com/06.02.03.0005
		math.Pow(1-x*x, float64(m)/2)

	// Recurrence formula
	// P(k+1, m, x) = [(2*k+1)*x*P(k, m, x) - (k+m)*P(k-1, m, x)] / (k+1-m)
	if n > m {
		prev := 0. // P(m-1, m, x)
		for k := m; k < n; k++ {
			next := (float64(2*k+1)*x*res - float64(k+m)*prev) / float64(k+1-m)
			res, prev = next, res
		}
	}

	// Reflection formula
	// P(n, -m, x) = (-1)^m * (n-m)!/(n+m)! * P(n, m, x)
	if reflect {
		res *= sign * math.Gamma(float64(n-m+1)) / math.Gamma(float64(n+m+1))
	}

	return res
}
