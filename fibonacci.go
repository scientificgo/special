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
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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
//
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
