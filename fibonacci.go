package special

import "math"

// Fibonacci returns the nth Fibonacci. The Fibonacci numbers
// are defined, for integer n, by
//
//	F(n) = F(n-1) + F(n-2)
//	F(0) = F(1) = 1
//
// and can extended to non-integer n by
//
//	Fibonacci(x) = (φ**x - Cos(πx)/φ**x) / √5
//
// where φ = (1 + √5) / 2 is the golden ratio.
//
// See http://mathworld.wolfram.com/FibonacciNumber.html.
func Fibonacci(x float64) float64 {
	if math.IsNaN(x) || math.IsInf(x, -1) {
		return math.NaN()
	}
	if math.IsInf(x, +1) {
		return math.Inf(+1)
	}

	const sqrt5 = 2*math.Phi - 1

	phin := math.Pow(math.Phi, x)
	res := phin - math.Cos(math.Pi*x)/phin
	return res / sqrt5
}
