// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

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
//	Fibonacci(n) = (φ**n - Cos(nπ)/φ**n) / √5
//
// where φ = (1 + √5) / 2 is the golden ratio.
//
// See http://mathworld.wolfram.com/FibonacciNumber.html for more
// information.
func Fibonacci(n float64) float64 {
	if math.IsNaN(n) || math.IsInf(n, -1) {
		return math.NaN()
	}
	if math.IsInf(n, +1) {
		return math.Inf(+1)
	}
	res := math.Pow(math.Phi, n)
	res = res - math.Cos(math.Pi*n)/res
	res = res / (2*math.Phi - 1)
	return res
}
